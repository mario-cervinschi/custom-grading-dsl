import {
  CompletionItem,
} from "vscode-languageserver/node";
import {
  createKeyword,
  createOperator,
  createSnippet,
} from "../utils/snippet-factory";
import { handleOperation } from "./normal-operation";

export function handleWhenOperation(
  prevNode: any,
  leftNode: any
): CompletionItem[] {
  if (leftNode.type === "WHEN") {
    return [
      createSnippet("var", "${1:x}", "simple variable"),
      createSnippet(
        "list",
        "${1:x1}, ${2:x2}",
        "list of variables where the conditions apply"
      ),
    ];
  }

  if (
    (leftNode.type === "identifier" || leftNode.type === "list") &&
    prevNode.type === "WHEN"
  ) {
    return [
      createOperator(",", "comma"),
      createSnippet("exists", "EXISTS", "terminator of when"),
    ];
  }

  if (leftNode.type === "," && prevNode.type === "identifier") {
    return [createSnippet("var", "${1:x}", "simple variable")];
  }

  let returnArray = handleOperation(prevNode, leftNode);
  if (leftNode.type === "EXISTS") {
    return [...returnArray, createKeyword("EXPLAIN", "EXPLAIN")];
  }
  return returnArray;
}
