import { handleApply } from "../rules/apply";
import { createKeyword } from "../utils/snippet-factory";
import {
  suggestComparisonRHS,
  suggestExpressionValue,
  suggestOperatorsAfterValue,
  suggestTernaryBranch,
} from "../rules/expression";
import { handleIdentifier } from "../rules/identifier";
import { handleOperation } from "../rules/normal-operation";
import { handleWhenOperation } from "../rules/when-operation";

const COMPARISON_OPS = new Set(["==", "!=", ">", ">=", "<", "<="]);
const VALUE_TYPES = new Set([
  "identifier", "number", "string", "constant",
  "primary_expression", "simple_expression", "expression",
  "binary_expression", "conditional_expression",
  "unary_expression", "function_call", "iterate_expression", ")"
]);

export function handleNext(children: any[], cursorOffset: number) {
  let leftNode = null;
  const types = Array.from(children, (child: any) => child.type);
  const firstType = types[0];

  for (const child of children) {
    if (child.endIndex <= cursorOffset) {
      if (!leftNode || child.endIndex > leftNode.endIndex) {
        leftNode = child;
      }
    }
  }

  if (!leftNode) {
    return [];
  }

  const leftNodeIndex = children.indexOf(leftNode);
  const prevNode = leftNodeIndex > 0 ? children[leftNodeIndex - 1] : null;

  console.log("Left:", leftNode.type, "| Prev:", prevNode?.type);

  // --- Expression-level operators (work in ANY context) ---

  // After ( → suggest values inside group
  if (leftNode.type === "(") {
    return suggestComparisonRHS();
  }

  // After regex match operator → suggest regex pattern snippet
  if (leftNode.type === "~" || leftNode.type === "!~") {
    const { createSnippet } = require("../utils/snippet-factory");
    return [createSnippet("regex", '"^${1:pattern}"', "Regex match string")];
  }

  // After ? → suggest ternary true-branch (values + nested ternary + constants)
  if (leftNode.type === "?") {
    return suggestTernaryBranch();
  }

  // After : → suggest ternary false-branch (values + nested ternary + constants)
  if (leftNode.type === ":") {
    return suggestTernaryBranch();
  }

  // After equality operators (==, !=) → suggest constants + values
  if (["==", "!="].includes(leftNode.type)) {
    return suggestComparisonRHS();
  }

  // After relational operators (>, <, >=, <=) → suggest only math concepts (variables, numbers)
  if ([">", "<", ">=", "<="].includes(leftNode.type)) {
    const { suggestArithmeticRHS } = require("../rules/expression");
    return suggestArithmeticRHS();
  }

  // After logical operators → suggest values and constants
  if (["&&", "||"].includes(leftNode.type)) {
    return suggestComparisonRHS();
  }

  // After arithmetic operators → suggest only math concepts (variables, numbers)
  if (["+", "-", "*", "/"].includes(leftNode.type)) {
    const { suggestArithmeticRHS } = require("../rules/expression");
    return suggestArithmeticRHS();
  }

  // Helps drill down into grouped nodes like iterate_expression
  const findLastLeaf = (n: any) => {
    let curr = n;
    while (curr && curr.childCount > 0) {
      curr = curr.child(curr.childCount - 1);
    }
    return curr;
  };

  const getOpSuggestions = (allowTO: boolean = true) => {
    const lastLeaf = findLastLeaf(leftNode);
    const { DSL_CONSTANTS } = require("../rules/expression");
    const isBasicValue = lastLeaf && (
      ["constant", "number", "string"].includes(lastLeaf.type) || 
      DSL_CONSTANTS.includes(lastLeaf.type)
    );

    if (allowTO && isBasicValue && (firstType === "APPLY" || firstType === "apply")) {
      return [createKeyword("TO", "TO ")];
    }

    const ops = suggestOperatorsAfterValue();
    if (allowTO && (firstType === "APPLY" || firstType === "apply")) {
      ops.push(createKeyword("TO", "TO "));
    }
    return ops;
  };

  // --- Statement-level dispatch (first priorities) ---
  // If the very first word in our evaluated statement is an identifier, and we are currently touching it
  // (or it's the only node), it shouldn't get generic math operators; it needs an Assignment '='!
  if (firstType === "identifier" && leftNode.type === "identifier") {
      const { handleIdentifier } = require("../rules/identifier");
      return handleIdentifier(prevNode, leftNode);
  }

  // After a value/expression → suggest operators (including ? for ternary)
  if (VALUE_TYPES.has(leftNode.type) && !prevNode) {
    return getOpSuggestions();
  }

  // "? value |" → suggest : (continue ternary) + next values + operators
  if (prevNode?.type === "?" && VALUE_TYPES.has(leftNode.type)) {
    const { suggestExpressionValue } = require("../rules/expression");
    const colon = createKeyword(":", ": ");
    colon.detail = "else branch";
    colon.sortText = "000"; // Priority!
    
    const nextValues = suggestExpressionValue();
    for (const v of nextValues) v.sortText = "050";

    const ops = getOpSuggestions(false);
    for (const op of ops) op.sortText = "999";

    return [colon, ...nextValues, ...ops];
  }

  // ": value |" → suggest operators (end of ternary, expression might continue)
  if (prevNode?.type === ":" && VALUE_TYPES.has(leftNode.type)) {
    return getOpSuggestions();
  }

  // After IN → suggest [ to start list
  if (leftNode.type === "IN") {
    return [createKeyword("[", "[ ")];
  }

  // After ] → suggest ? (optional ternary) + operators
  if (leftNode.type === "]") {
    return getOpSuggestions();
  }

  // --- Statement-level dispatch (based on first token) ---

  switch (firstType) {
    case "APPLY":
    case "apply":
      return handleApply(prevNode, leftNode);

    case "identifier":
      return handleIdentifier(prevNode, leftNode);

    case "EXPLAIN":
    case "explain":
      return handleOperation(prevNode, leftNode);

    case "WHEN":
    case "when":
      return handleWhenOperation(prevNode, leftNode);

    case "=":
    case "=>":
      return suggestExpressionValue();
  }

  return [];
}
