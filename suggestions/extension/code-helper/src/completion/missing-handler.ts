import * as vscode from "vscode";
import {
  createKeyword,
  createOperator,
  createSnippet,
} from "../utils/snippet-factory";
import { suggestTernaryBranch } from "../rules/expression";

export function handleMissingNode(
  node: any,
  position: vscode.Position
): vscode.CompletionItem[] {
  console.log("MISSING NODE:", node.type);

  if (node.type === ";") {    
    const prevNode = node.previousSibling;
    console.log(prevNode);
    if (
      prevNode &&
      (prevNode.type === "list" || prevNode.type === "expression")
    ) {
      const comma = createOperator(",", "Separator");
      const semicolon = createOperator(";", "End Statement");
      return [comma, semicolon];
    }
  }

  if (node.type === "identifier") {
    return [
      createSnippet("missing variable", "${1:var_name}", "Required variable"),
    ];
  }

  if (node.type === "expression" || node.type === "primary_expression") {
    return suggestTernaryBranch();
  }

  // If tree-sitter expects a constant from the constant list syntax
  // it might emit a MISSING "nothing" (the first literal in the grammar rule).
  const { suggestComparisonRHS, DSL_CONSTANTS } = require("../rules/expression");
  if (node.type === "nothing" || DSL_CONSTANTS.includes(node.type)) {
    return suggestComparisonRHS();
  }

  const item = createKeyword(node.type, node.type);
  item.detail = "end operation operator";
  item.sortText = "000";
  // item.kind = vscode.CompletionItemKind.Event;

  return [item];
}
