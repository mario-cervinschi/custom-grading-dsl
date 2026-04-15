import { CompletionItem, Position } from "vscode-languageserver/node";
import { TextDocument } from "vscode-languageserver-textdocument";
import { getParser } from "../parser/tree-sitter-wrapper";
import { createKeyword, createSnippet } from "../utils/snippet-factory";
import { handleNext } from "./context-handler";
import { handleMissingNode } from "./missing-handler";
import { findRelevantNodes } from "./node-finder";

import * as fs from "fs";
import * as path from "path";

const logPath = path.join(__dirname, "PROVIDER_LOG.txt");

function logToFile(message: any) {
  const text =
    typeof message === "string" ? message : JSON.stringify(message, null, 2);
  fs.appendFileSync(logPath, text + "\n");
}

export function provideCompletions(
  document: TextDocument,
  position: Position,
): CompletionItem[] {
  try {
    const text = document.getText();
    const offset = document.offsetAt(position);
    const parser = getParser();
    const tree = parser.parse(text);

    const { node, candidates } = findRelevantNodes(tree, offset);

    logToFile("------ ------ ------");
    logToFile(`Node type: ${node?.type}`);
    logToFile(`Tree structure: ${tree.rootNode.toString()}`);

    for (const candidate of candidates) {
      if (candidate.isMissing) {
        return handleMissingNode(candidate, position);
      }
    }

    if (node && node.parent && node.parent.type === "ERROR") {
      return handleNext(node.parent.children, offset);
    }

    return [
      createKeyword("APPLY", "APPLY "),
      createKeyword("EXPLAIN", "EXPLAIN "),
      createKeyword("WHEN", "WHEN "),
      createSnippet("variable", "${1:x}", "simple variable"),
    ];
  } catch (e) {
    console.error(e);
    return [];
  }
}
