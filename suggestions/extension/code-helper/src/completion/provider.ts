import * as vscode from "vscode";
import { getParser } from "../parser/tree-sitter-wrapper";
import { createKeyword, createSnippet } from "../utils/snippet-factory";
import { suggestOperatorsAfterValue } from "../rules/expression";
import { handleNext } from "./context-handler";
import { handleMissingNode } from "./missing-handler";
import { findRelevantNodes } from "./node-finder";

/**
 * Node types that represent a "value" — i.e., a complete expression
 * after which operators like ?, ==, &&, etc. could follow.
 */
const VALUE_TYPES = new Set([
  "identifier", "number", "string", "constant",
  "primary_expression", "simple_expression", "expression",
  "binary_expression", "unary_expression", "function_call",
  "conditional_expression", "iterate_expression", ")"
]);

export class MyCompletionProvider implements vscode.CompletionItemProvider {
  provideCompletionItems(
    document: vscode.TextDocument,
    position: vscode.Position
  ) {
    try {
      const text = document.getText();
      const offset = document.offsetAt(position);
      const parser = getParser();
      const tree = parser.parse(text);

      const { node, candidates } = findRelevantNodes(tree, offset);

      console.log("Node:", node?.type);
      console.log(tree.rootNode.toString());

      // --- Step 1.5: Robust text-based heuristic for IN [...] lists ---
      const textBeforeCursor = text.substring(0, offset);
      const lastOpenBracket = textBeforeCursor.lastIndexOf("[");
      const lastCloseBracket = textBeforeCursor.lastIndexOf("]");
      
      if (lastOpenBracket > lastCloseBracket) {
        const textInside = textBeforeCursor.substring(lastOpenBracket + 1);
        const trimmedInside = textInside.trim();
        const lastChar = trimmedInside.slice(-1);
        
        const { suggestComparisonRHS } = require("../rules/expression");
        let listSuggestions: vscode.CompletionItem[] = [];

        if (trimmedInside === "" || lastChar === ",") {
           // Right after '[' or ',' -> suggest ONLY list values (variables, constants, numbers)
           listSuggestions = suggestComparisonRHS();
        } else {
           // After a value -> suggest ONLY ',' or ']'
           listSuggestions = [
             createKeyword(",", ", "),
             createKeyword("]", "] ")
           ];
        }
        
        return finalize(listSuggestions, position);
      }

      // --- Drill down to true leaf before evaluating contexts ---
      const findLastLeaf = (n: any) => {
        let curr = n;
        while (curr && curr.childCount > 0) {
          let found = false;
          for (let i = curr.childCount - 1; i >= 0; i--) {
            const child = curr.child(i);
            if (!child.isMissing) {
              curr = child;
              found = true;
              break;
            }
          }
          if (!found) break;
        }
        return curr;
      };
      const leafNode = findLastLeaf(node);

      // --- Step 2: ERROR parent → context-based suggestions ---
      // We look for an ERROR node either at the leaf position or as the direct parent.
      const errorNode = (node?.type === "ERROR") ? node : (node?.parent?.type === "ERROR" ? node.parent : null);
      if (errorNode) {
        console.log("=== STEP 2 TRIGGERED ===");
        const ctxSuggestions = handleNext(errorNode.children, offset);
        console.log("Step 2 ctxSuggestions length:", ctxSuggestions.length);
        
        if (ctxSuggestions.length > 0) {
          console.log("Step 2 ctxSuggestions items:", ctxSuggestions.map((s: any) => typeof s.label === "string" ? s.label : s.label.label));
          return finalize(ctxSuggestions, position);
        }
        
        // Failsafe: if we are at an identifier inside an error, it's likely an assignment start
        if (leafNode && leafNode.type === "identifier") {
          console.log("Step 2 Failsafe triggered: forcing '=' operator for identifier.");
          const { createOperator } = require("../utils/snippet-factory");
          return finalize([createOperator("=", "=")], position);
        }
        
        return finalize([], position);
      }
      console.log("=== STEP 2 SKIPPED ===");

      // --- Step 3: After a complete value/expression → suggest operators + structural tokens ---
      if (leafNode && VALUE_TYPES.has(leafNode.type)) {
        console.log("=== STEP 3 TRIGGERED ===");
        let suggestions = suggestAfterCompleteExpression(leafNode, offset, position);
        console.log("Step 3 initial suggestions length:", suggestions.length);
        
        let w: any = leafNode;
        let inLambda = false;
        let normalOpNode = null;
        let applyNode = null;

        while (w) {
          if (w.type === "lambda") inLambda = true;
          if (w.type === "normal_operation") normalOpNode = w;
          if (w.type === "apply") applyNode = w;
          w = w.parent;
        }

        // Handle LHS and RHS logic for any normal_operation (basic expressions, including EXPLAIN and EXISTS)
        if (normalOpNode) {
             console.log("Found normal_operation context");
             let hasEquals = false;
             let searchQueue = [normalOpNode];
             while (searchQueue.length > 0) {
               const n = searchQueue.pop();
               if (n.type === "=" && n.endIndex <= offset) {
                 hasEquals = true;
                 break;
               }
               for (let i = 0; i < n.childCount; i++) searchQueue.push(n.child(i));
             }

             if (!hasEquals) {
               // LHS logic
               // If we don't have an equal sign yet, and we are on an identifier or expression that wraps one,
               // the ONLY syntactic validity in `assignment_expression` is an '=' operator.
               if (leafNode && (leafNode.type === "identifier" || leafNode.type === "expression" || leafNode.type === "primary_expression")) {
                   console.log("LHS detected: forcing '=' suggestion");
                   const { createOperator } = require("../utils/snippet-factory");
                   return finalize([createOperator("=", "=")], position);
               }
               
               const ctxSuggestions = handleNext(normalOpNode.children, offset);
               if (ctxSuggestions.length > 0) {
                 return finalize(ctxSuggestions, position);
               }
             } else {
               // RHS logic
               console.log("RHS detected: filtering for punctuation");
               const lineSuffix = text.substring(offset).split('\n')[0];
               const insideParentheses = lineSuffix.includes(')');
               
               const firstNode = normalOpNode.childCount > 0 ? normalOpNode.child(0) : null;
               const hasExplainOrExists = firstNode && ["EXPLAIN", "explain", "EXISTS", "exists"].includes(firstNode.type);
               
               suggestions = suggestions.filter(s => {
                 const lbl = typeof s.label === "string" ? s.label : s.label.label;
                 const isPunctuation = lbl.trim() === "," || lbl.trim() === ";";
                 
                 if (insideParentheses) {
                   // Inside (): We only block structural punctuation like ; and ,
                   return !isPunctuation;
                 } else {
                   // Outside (): Standard clean-up, only show structural terminals
                   if (lbl.trim() === ";") return true;
                   if (lbl.trim() === "," && hasExplainOrExists) return true;
                   return false;
                 }
               });

               // Balanced suggestions inside parentheses:
               // 1. High Priority for Values (Variable, Constant, Number)
               // 2. Low Priority for Math/Logic Operators
               if (insideParentheses && leafNode && VALUE_TYPES.has(leafNode.type)) {
                 const { suggestExpressionValue } = require("../rules/expression");
                 const valSuggestions = suggestExpressionValue();
                 
                 // Move current suggestions (operators) to the bottom
                 for (const s of suggestions) {
                    s.sortText = "999";
                 }
                 
                 // Move value suggestions (IDs, numbers, constants) to the top
                 for (const s of valSuggestions) {
                    s.sortText = "050";
                 }

                 suggestions = [...valSuggestions, ...suggestions];
               }
               
               return finalize(suggestions, position);
             }
        }

        // Handle apply statement context (e.g. `apply ... to a |`)
        if (applyNode) {
          console.log("Found apply context: filtering for punctuation in target list");
          suggestions = suggestions.filter(s => {
            const lbl = typeof s.label === "string" ? s.label : s.label.label;
            return lbl.trim() === "," || lbl.trim() === ";";
          });
          return finalize(suggestions, position);
        }

        if (inLambda) {
          suggestions.push(createKeyword("TO", "TO "));
        }

        // Apply user preference: if the current leaf is a basic value and we are in a lambda,
        // ONLY suggest TO (remove operators to reduce clutter)
        const { DSL_CONSTANTS } = require("../rules/expression");
        if (inLambda && (
            ["constant", "number", "string"].includes(leafNode.type) ||
            DSL_CONSTANTS.includes(leafNode.type)
        )) {
           suggestions = [createKeyword("TO", "TO ")];
           return finalize(suggestions, position);
        }

        if (suggestions.length > 0) {
          return finalize(suggestions, position);
        }
      }

      // --- Step 3.5: After an operator, suggest values ---
      const OPERATOR_TYPES = new Set(["+", "-", "*", "/", "&&", "||", "==", "!=", ">", "<", ">=", "<=", "=>", "?", ":", "IN", "="]);
      if (leafNode && (OPERATOR_TYPES.has(leafNode.type.toUpperCase()) || OPERATOR_TYPES.has(leafNode.type))) {
         if (["+", "-", "*", "/", ">", "<", ">=", "<="].includes(leafNode.type)) {
            const { suggestArithmeticRHS } = require("../rules/expression");
            return finalize(suggestArithmeticRHS(), position);
         }
         const { suggestComparisonRHS } = require("../rules/expression");
         return finalize(suggestComparisonRHS(), position);
      }

      // --- Step 4: Default top-level suggestions ---
      return [
        createKeyword("APPLY", "APPLY "),
        createKeyword("EXPLAIN", "EXPLAIN "),
        createKeyword("WHEN", "WHEN "),
        createSnippet("variable", "${1:x}", "simple variable"),
      ];
    } catch (e) {
      console.error(e);
      return [];
    }
  }
}

/**
 * When the cursor is right after a complete expression/value node,
 * suggest both expression operators (?, ==, &&...) and structural
 * tokens from MISSING siblings up the tree (TO, ;, :).
 */
function suggestAfterCompleteExpression(
  node: any,
  offset: number,
  position: vscode.Position
): vscode.CompletionItem[] {
  const suggestions: vscode.CompletionItem[] = [];

  // Add expression continuation operators (?, ==, &&, ||, +, -, etc.)
  suggestions.push(...suggestOperatorsAfterValue());

  // Walk up the tree and find the first MISSING sibling after our position
  // This gives us structural tokens the grammar expects (TO, ;, :)
  let walker = node;
  while (walker && walker.parent) {
    const parent = walker.parent;

    let foundWalker = false;
    for (let i = 0; i < parent.childCount; i++) {
      const sibling = parent.child(i);
      if (!sibling) { continue; }

      // Find our walker node in the parent's children
      if (sibling.id === walker.id) {
        foundWalker = true;
        continue;
      }

      // After finding our position, look for the first MISSING sibling
      if (foundWalker && sibling.isMissing) {
        const missingSuggestions = handleMissingNode(sibling, position);
        suggestions.push(...missingSuggestions);
        break; // Only the first MISSING sibling matters
      }
    }

    walker = walker.parent;
    if (walker.type === "source_file") { break; }
  }

  return suggestions;
}

/**
 * Assigns zero-width Replacement ranges to operators and punctuation completions
 * so that VS Code does not filter them out when the cursor touches a word boundary.
 */
function finalize(
  suggestions: vscode.CompletionItem[],
  position: vscode.Position
): vscode.CompletionItem[] {
  for (const s of suggestions) {
    const labelResult = typeof s.label === "string" ? s.label : s.label?.label || "";
    const label = labelResult.trim();
    if (label === "=") {
      if (!s.range) {
        s.range = new vscode.Range(position, position);
      }
    }
  }
  return suggestions;
}
