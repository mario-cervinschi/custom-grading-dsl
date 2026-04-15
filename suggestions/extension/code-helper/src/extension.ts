import * as vscode from "vscode";
import { MyCompletionProvider } from "./completion/provider";

export function activate(context: vscode.ExtensionContext) {
  console.log("Tree-Sitter Debugger");

  try {
    const provider = vscode.languages.registerCompletionItemProvider(
      { scheme: "file", language: "codesuggestion" },
      new MyCompletionProvider(),
      " ",
      "=",
      ">",
      ","
    );

    context.subscriptions.push(provider);
  } catch (e: any) {
    vscode.window.showErrorMessage(`Init failed: ${e.message}`);
  }
}

export function deactivate() { }
