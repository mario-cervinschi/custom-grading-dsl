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

documents.listen(connection);
connection.listen();
