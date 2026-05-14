import {
  createConnection,
  TextDocuments,
  CompletionItem,
  ProposedFeatures,
  InitializeParams,
  CompletionParams,
  TextDocumentSyncKind,
} from "vscode-languageserver/node";
import { TextDocument } from "vscode-languageserver-textdocument";
import { provideCompletions } from "./completion/provider";
import { getParser } from "./parser/tree-sitter-wrapper";
import path from "path";
import * as fs from "fs";

const logFile = path.join(__dirname, "TEST_LOG.txt");
fs.writeFileSync(logFile, "created\n");

const connection = createConnection(ProposedFeatures.all);
const documents = new TextDocuments(TextDocument);

connection.onInitialize((params: InitializeParams) => {
  return {
    capabilities: {
      textDocumentSync: TextDocumentSyncKind.Incremental,
      completionProvider: {
        triggerCharacters: [" "],
      },
    },
  };
});

connection.onCompletion((params: CompletionParams): CompletionItem[] => {
  const doc = documents.get(params.textDocument.uri);
  if (!doc) {
    return [];
  }

  try {
    return provideCompletions(doc, params.position);
  } catch (error) {
    fs.appendFileSync(logFile, `error: ${error}\n`);
    return [];
  }
});

// validate: parse text, find ERROR nodes
connection.onRequest("custom/validate", (params: any) => {
  const text: string = params.text || "";
  const tree = getParser().parse(text);
  const errors: any[] = [];
  function walk(node: any) {
    if (node.type === "ERROR" || node.isMissing) {
      errors.push({
        line: node.startPosition.row,
        col: node.startPosition.column,
        endLine: node.endPosition.row,
        endCol: node.endPosition.column,
        msg: node.isMissing ? `missing '${node.type}'` : "syntax error",
      });
    }
    for (let i = 0; i < node.childCount; i++) walk(node.child(i));
  }
  walk(tree.rootNode);
  return { valid: errors.length === 0, errors };
});

documents.listen(connection);
connection.listen();
