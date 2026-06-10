import * as fs from 'fs';
import * as path from 'path';
import {
  CompletionItem,
  Position,
  Range,
  TextEdit
} from "vscode-languageserver/node";
import {
  createKeyword,
  createOperator,
  createSnippet,
} from "../utils/snippet-factory";

const logPath = path.join(__dirname, 'PROVIDER_LOG.txt');

function logToFile(message: any) {
  const text = typeof message === 'string' ? message : JSON.stringify(message, null, 2);
  fs.appendFileSync(logPath, text + '\n');
}

export function handleMissingNode(
  node: any,
  position: Position
): CompletionItem[] {
  logToFile(`MISSING NODE: ${node.type}`);

  if (node.type === ";") {    
    const prevNode = node.previousSibling;
    logToFile(prevNode ? `Previous node: ${prevNode.type}` : "Previous node: null");
    if (
      prevNode &&
      (prevNode.type === "list" || prevNode.type === "expression")
    ) {
      const insertRange = Range.create(position, position);
      
      const comma = createOperator(",", "Separator");
      comma.textEdit = TextEdit.replace(insertRange, ",");
      
      const semicolon = createOperator(";", "End Statement");
      semicolon.textEdit = TextEdit.replace(insertRange, ";");
      
      return [comma, semicolon];
    }
  }

  if (node.type === "identifier") {
    return [
      createSnippet("missing variable", "${1:var_name}", "Required variable"),
    ];
  }

  if (node.type === "expression" || node.type === "primary_expression") {
    return [createSnippet("missing value", "${1:value}", "Required value")];
  }

  const item = createKeyword(node.type, node.type);
  item.detail = "end operation operator";
  item.sortText = "000";

  return [item];
}