import {
  CompletionItem,
} from "vscode-languageserver/node";
import { createSnippet } from "../utils/snippet-factory";
export function handleOperation(
  prevNode: any,
  leftNode: any
): CompletionItem[] {
  if (leftNode.type === "EXPLAIN" || leftNode.type === "EXISTS") {
    return [
      createSnippet("expression", "${1:x} = ${2:y + z}", "simple expression"),
    ];
  }

  if (leftNode.type === "," && prevNode.type === "expression") {
    return [
      createSnippet(
        "explanation",
        '"${1:explanation}"',
        "explain what operation does"
      ),
    ];
  }

  return [];
}
