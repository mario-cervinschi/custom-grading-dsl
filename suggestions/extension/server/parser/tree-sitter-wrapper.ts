import Parser from "tree-sitter";
import MyGrammar from "tree-sitter-code-suggestion";

let parserInstance: any = null;

export function getParser() {
  if (!parserInstance) {
    parserInstance = new Parser();
    parserInstance.setLanguage(MyGrammar);
  }
  return parserInstance;
}
