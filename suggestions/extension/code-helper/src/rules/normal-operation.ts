import * as vscode from "vscode";
import { createKeyword, createSnippet } from "../utils/snippet-factory";
import { suggestExpressionValue } from "./expression";
export function handleOperation(
  prevNode: any,
  leftNode: any
): vscode.CompletionItem[] {
  if (leftNode.type === "EXPLAIN" || leftNode.type === "EXISTS") {
    return [
      createSnippet("expression", "${1:x} = ${2:y + z}", "simple expression"),
    ];
  }

  if (prevNode?.type === "EXPLAIN" || prevNode?.type === "EXISTS") {
    let hasEquals = false;
    let queue = [leftNode];
    while (queue.length > 0) {
      const node = queue.pop();
      if (node?.type === "=") {
        hasEquals = true;
        break;
      }
      for (let i = 0; i < (node?.childCount || 0); i++) {
        queue.push(node.child(i));
      }
    }

    if (!hasEquals) {
      let leaf = leftNode;
      while (leaf && leaf.childCount > 0) {
        leaf = leaf.child(leaf.childCount - 1);
      }
      if (leaf && leaf.type === "identifier") {
        return [createKeyword("=", "= ")];
      }
    }
  }

  let leafForEquals = leftNode;
  while (leafForEquals && leafForEquals.childCount > 0) {
    leafForEquals = leafForEquals.child(leafForEquals.childCount - 1);
  }

  if (leafForEquals && leafForEquals.type === "=") {
    return suggestExpressionValue();
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
