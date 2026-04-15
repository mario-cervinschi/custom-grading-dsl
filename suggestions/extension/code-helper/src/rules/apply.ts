import * as vscode from "vscode";
import { createKeyword, createSnippet } from "../utils/snippet-factory";
import {
  suggestComparisonRHS,
  suggestExpressionValue,
  suggestOperatorsAfterValue,
  suggestTernaryBranch,
} from "./expression";

const COMPARISON_OPS = new Set(["==", "!=", ">", ">=", "<", "<="]);

export function handleApply(prevNode: any, leftNode: any): vscode.CompletionItem[] {
  if (leftNode.type === "APPLY") {
    return [createSnippet("var", "${1:x}", "simple variable")];
  }

  if (prevNode?.type === "APPLY" && leftNode.type === "identifier") {
    return [createKeyword("=>", "=>")];
  }

  if (leftNode.type === "=>") {
    return suggestExpressionValue();
  }

  // After ? or : inside a ternary within APPLY lambda
  if (leftNode.type === "?" || leftNode.type === ":") {
    return suggestTernaryBranch();
  }

  // After comparison operator inside APPLY lambda
  if (COMPARISON_OPS.has(leftNode.type)) {
    return suggestComparisonRHS();
  }

  if (prevNode?.type === "APPLY" && leftNode.type === "lambda") {
    return [createKeyword("TO", "TO "), ...suggestOperatorsAfterValue()];
  }

  if (prevNode?.type === "lambda" && leftNode.type === "TO") {
    return [
      createSnippet("var", "${1:x}", "simple variable"),
      createSnippet(
        "list",
        "${1:x1}, ${2:x2}",
        "list of variables where the conditions apply"
      ),
    ];
  }

  if (prevNode?.type === "identifier" && leftNode.type === ",") {
    return [createSnippet("var", "${1:x}", "simple variable")];
  }

  return [];
}