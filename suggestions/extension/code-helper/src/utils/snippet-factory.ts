import * as vscode from "vscode";

export function createKeyword(
  label: string,
  text: string
): vscode.CompletionItem {
  const item = new vscode.CompletionItem(
    label,
    vscode.CompletionItemKind.Keyword
  );
  item.insertText = text;
  return item;
}

export function createOperator(
  label: string,
  doc: string
): vscode.CompletionItem {
  const item = new vscode.CompletionItem(
    label,
    vscode.CompletionItemKind.Operator
  );
  item.insertText = label + " ";
  item.documentation = doc;
  return item;
}

export function createSnippet(
  label: string,
  snippetCode: string,
  description: string
): vscode.CompletionItem {
  const item = new vscode.CompletionItem(
    label,
    vscode.CompletionItemKind.Snippet
  );

  item.insertText = new vscode.SnippetString(snippetCode);
  item.detail = description;

  let cleanPreview = snippetCode
    .replace(/\$\{\d+:([^}]+)\}/g, "$1")
    .replace(/\$\d+/g, "");

  const docs = new vscode.MarkdownString();
  docs.appendCodeblock(cleanPreview, "plaintext");
  docs.appendMarkdown(`---`);
  docs.appendMarkdown(`\n${description}`);
  item.documentation = docs;
  return item;
}
