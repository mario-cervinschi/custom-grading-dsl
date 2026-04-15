import {
  CompletionItem,
} from "vscode-languageserver/node";
import { createOperator, createSnippet } from "../utils/snippet-factory";

export function handleIdentifier(
  prevNode: any,
  leftNode: any
): CompletionItem[] {
  if (leftNode.type === "identifier") {
    return [
      createOperator("=", "="),
      createSnippet("expression", "= ${1:y + z}", "simple expression"),
    ];
  }

  if (leftNode.type === "=" && prevNode.type === "identifier") {
    return [createSnippet("expression", "${1:y + z}", "simple expression")];
  }

  return [];
}
