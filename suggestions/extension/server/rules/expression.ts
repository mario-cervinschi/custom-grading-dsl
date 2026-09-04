import {
  CompletionItem,
  CompletionItemKind,
  Position,
  Range,
  TextDocumentPositionParams,
} from "vscode-languageserver/node";
import { createKeyword, createSnippet } from "../utils/snippet-factory";

export const DSL_CONSTANTS = [
  "nothing",
  "absent",
  "present",
  "excused",
  "unexcused",
  "fraud",
  "suspended",
];

/**
 * Basic value suggestions (Variable, Number, Constants, and common snippets).
 * This is the generic "what can go here as a value" handler.
 */
export function suggestExpressionValue(): CompletionItem[] {
  const items: CompletionItem[] = [
    createSnippet("variable", "${1:x}", "simple variable"),
    createKeyword("number", "0"),
  ];

  for (const c of DSL_CONSTANTS) {
    const item = createKeyword(c, c);
    item.detail = "constant value";
    item.sortText = "1_" + c;
    items.push(item);
  }

  items.push(
    createSnippet(
      "ternary conditional",
      "${1:x} == ${2:val} ? ${3:trueVal} : ${4:falseVal}",
      "inline if-else"
    ),
    createSnippet(
      "nested ternary",
      "${1:x} == ${2:val1} ? ${3:result1} : ${4:x} == ${5:val2} ? ${6:result2} : ${7:default}",
      "chained if-else-if"
    ),
    createSnippet(
      "verify existence",
      "${1:x} in [${2:a}, ${3:b}] ? ${4:y} : ${5:n}",
      "if x in list else"
    ),
    createSnippet(
      "function call",
      "${1|MAX,MIN,ROUND|}(${2:expr})",
      "function call"
    )
  );

  return items;
}

/**
 * Suggestions after a comparison operator (==, !=, >, <, >=, <=).
 */
export function suggestComparisonRHS(): CompletionItem[] {
  const items: CompletionItem[] = [];

  for (const c of DSL_CONSTANTS) {
    const item = createKeyword(c, c);
    item.detail = "constant";
    item.sortText = "0_" + c; // sort constants
    items.push(item);
  }

  items.push(
    createSnippet("variable", "${1:var_name}", "compare with variable"),
    createKeyword("number", "0"),
  );

  return items;
}

/**
 * Suggestions after an arithmetic operator (+, -, *, /).
 */
export function suggestArithmeticRHS(): CompletionItem[] {
  const items: CompletionItem[] = [
    createSnippet("variable", "${1:var_name}", "use variable"),
    createKeyword("number", "0"),
  ];

  for (const c of DSL_CONSTANTS) {
    items.push(createKeyword(c, c));
  }

  return items;
}

/**
 * Simple list of all logical and arithmetic operators.
 */
export function suggestOperatorsAfterValue(): CompletionItem[] {
  const ops = ["==", "!=", ">", "<", ">=", "<=", "&&", "||", "+", "-", "*", "/", "~", "!~", "?", "IN"];
  return ops.map((op) => ({
    label: op,
    kind: CompletionItemKind.Operator,
    insertText: op + " ",
  }));
}

/**
 * Suggestions after '?' in a ternary, the "true" branch.
 */
export function suggestTernaryBranch(): CompletionItem[] {
  const items: CompletionItem[] = [];

  items.push(
    createSnippet("variable", "${1:x}", "simple variable"),
    createKeyword("number", "0"),
    createSnippet(
      "ternary conditional",
      "${1:x} == ${2:val} ? ${3:trueVal} : ${4:falseVal}",
      "nested ternary"
    )
  );

  return items;
}
