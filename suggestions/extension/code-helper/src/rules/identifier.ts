import * as vscode from "vscode";
import { createOperator, createSnippet } from "../utils/snippet-factory";

export function handleIdentifier(
  prevNode: any,
  leftNode: any
): vscode.CompletionItem[] {
  if (leftNode.type === "identifier") {
    return [
      createOperator("=", "=")
    ];
  }

  if (leftNode.type === "=" && prevNode.type === "identifier") {
    return [createSnippet("expression", "${1:y + z}", "simple expression")];
  }

  return [];
}
