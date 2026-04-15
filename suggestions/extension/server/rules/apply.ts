import {
  CompletionItem,
} from "vscode-languageserver/node";
import { createKeyword, createSnippet } from "../utils/snippet-factory";

export function handleApply(prevNode: any, leftNode: any): CompletionItem[] {
  if (leftNode.type === "APPLY") {
    return [createSnippet("var", "${1:x}", "simple variable")];
  }

  if (prevNode.type === "APPLY" && leftNode.type === "identifier") {
    return [createKeyword("=>", "=>")];
  }

  if (leftNode.type === "=>") {
    return [
      createKeyword("simple value", "0"),
      createSnippet("var", "${1:x}", "simple variable"),

      createSnippet(
        "ternary conditional",
        "${1:x} == ${2:val} ? ${3:trueVal} : ${4:falseVal}",
        "inline if else"
      ),
      createSnippet(
        "verify existance",
        "${1:x} in [${2:a}, ${3:b}] ? ${4:y} : ${5:n}",
        "if x in list else"
      ),
    ];
  }

  if (prevNode.type === "APPLY" && leftNode.type === "lambda") {
    return [createKeyword("TO", "TO")];
  }

  if (prevNode.type === "lambda" && leftNode.type === "TO") {
    return [
      createSnippet("var", "${1:x}", "simple variable"),
      createSnippet(
        "list",
        "${1:x1}, ${2:x2}",
        "list of variables where the conditions apply"
      ),
    ];
  }

  if (prevNode.type === "identifier" && leftNode.type === ",") {
    return [createSnippet("var", "${1:x}", "simple variable")];
  }

  return [];
}