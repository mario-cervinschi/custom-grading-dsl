import {
  CompletionItem,
  CompletionItemKind,
  InsertTextFormat,
  MarkupKind,
} from "vscode-languageserver/node";

export function createKeyword(
  label: string,
  text: string
): CompletionItem {
  return {
    label,
    kind: CompletionItemKind.Keyword,
    insertText: text,
  };
}

export function createOperator(
  label: string,
  doc: string
): CompletionItem {
  return {
    label,
    kind: CompletionItemKind.Operator,
    insertText: label + " ",
    documentation: doc,
  };
}

export function createSnippet(
  label: string,
  snippetCode: string,
  description: string
): CompletionItem {
  const cleanPreview = snippetCode
    .replace(/\$\{\d+:([^}]+)\}/g, "$1")
    .replace(/\$\d+/g, "");

  return {
    label,
    kind: CompletionItemKind.Snippet,
    insertText: snippetCode,
    insertTextFormat: InsertTextFormat.Snippet,
    detail: description,
    documentation: {
      kind: MarkupKind.Markdown,
      value: `\`\`\`plaintext\n${cleanPreview}\n\`\`\`\n---\n${description}`,
    },
  };
}