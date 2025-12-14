"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
Object.defineProperty(exports, "__esModule", { value: true });
exports.activate = activate;
exports.deactivate = deactivate;
const vscode = __importStar(require("vscode"));
let Parser;
let MyGrammar;
function activate(context) {
    console.log("Activare extensie...");
    try {
        Parser = require("tree-sitter");
        MyGrammar = require("tree-sitter-code-suggestion");
        const parser = new Parser();
        parser.setLanguage(MyGrammar);
        const provider = vscode.languages.registerCompletionItemProvider({ scheme: "file", language: "plaintext" }, {
            provideCompletionItems(document, position) {
                const text = document.getText();
                const tree = parser.parse(text);
                const offset = document.offsetAt(position);
                let node = tree.rootNode.descendantForIndex(offset, offset);
                console.log("🌲 Root:", tree.rootNode.toString());
                console.log("📍 Current node:", node.toString());
                console.log("🔍 Node type:", node.type);
                console.log("👨‍👩‍👧 Parent:", node.parent?.type);
                return getSuggestionsForContext(node, tree.rootNode, document, position, offset);
            },
        }, " ", "=", ">", ",", "[", "]", "?", ":", ";", "(", ")");
        context.subscriptions.push(provider);
        vscode.window.showInformationMessage("Autocompletare activată!");
    }
    catch (error) {
        console.error("CRASH:", error);
        vscode.window.showErrorMessage(`Error: ${error.message}`);
    }
}
function getSuggestionsForContext(node, rootNode, document, position, offset) {
    // Dacă suntem la source_file, caută ERROR în copii
    let errorNode = null;
    if (node.type === 'source_file') {
        console.log("🔍 Căutare ERROR în copiii source_file...");
        for (let child of node.children) {
            if (child.type === 'ERROR') {
                errorNode = child;
                break;
            }
        }
    }
    else {
        // Altfel caută ERROR în sus
        errorNode = findNearestError(node);
    }
    if (errorNode) {
        console.log("⚠️ Găsit ERROR:", errorNode.toString());
        return analyzeErrorNode(errorNode, document, position);
    }
    // Dacă nu e error, analizează contextul normal
    return getSuggestionsForNormalContext(node, document, position);
}
function findNearestError(node) {
    let current = node;
    // Urcă în arbore până găsești un ERROR
    while (current) {
        if (current.type === 'ERROR') {
            return current;
        }
        current = current.parent;
    }
    // Sau dacă nodul curent are un frate ERROR
    if (node.parent) {
        for (let sibling of node.parent.children) {
            if (sibling.type === 'ERROR') {
                return sibling;
            }
        }
    }
    return null;
}
function analyzeErrorNode(errorNode, document, position) {
    const suggestions = [];
    // Analizează copiii nodului ERROR pentru a deduce ce lipsește
    const children = errorNode.children || [];
    const childTypes = children.map((c) => c.type);
    console.log("👶 Copii ERROR:", childTypes);
    // Pattern: APPLY singur → lipsește lambda
    if (childTypes.includes('APPLY') && !childTypes.includes('lambda')) {
        console.log("💡 Sugestie: lambda (după APPLY)");
        suggestions.push(createSnippet("lambda", "${1:x} => ${2:x + 1}", "Funcție lambda: x => expresie"));
        suggestions.push(createSnippet("lambda-full", "${1:param} => ${2:param} ${3|==,!=,>,<,>=,<=,+,-,*,/|} ${4:value}", "Lambda completă cu operator"));
        return suggestions;
    }
    // Pattern: APPLY + lambda → lipsește TO
    if (childTypes.includes('APPLY') && childTypes.includes('lambda')) {
        const hasTO = childTypes.includes('TO');
        const hasList = childTypes.includes('list');
        if (!hasTO) {
            console.log("💡 Sugestie: TO (după APPLY + lambda)");
            suggestions.push(createKeywordItem("TO", "TO "));
            return suggestions;
        }
        if (hasTO && !hasList) {
            console.log("💡 Sugestie: listă (după TO)");
            suggestions.push(createSnippet("list", "${1:id1}, ${2:id2}", "Listă de identificatori"));
            suggestions.push(createSnippet("identifier", "${1:identifier}", "Un identificator"));
            return suggestions;
        }
        if (hasTO && hasList) {
            console.log("💡 Sugestie: ; (la final)");
            suggestions.push(createItem(";", "Termină statement-ul", vscode.CompletionItemKind.Text));
            return suggestions;
        }
    }
    // Pattern: WHEN + list → lipsește EXISTS
    if (childTypes.includes('WHEN') && childTypes.includes('list')) {
        const hasExists = childTypes.includes('EXISTS');
        if (!hasExists) {
            console.log("💡 Sugestie: EXISTS (după WHEN + list)");
            suggestions.push(createKeywordItem("EXISTS", "EXISTS "));
            return suggestions;
        }
    }
    // Pattern: WHEN → lipsește listă
    if (childTypes.includes('WHEN') && !childTypes.includes('list')) {
        console.log("💡 Sugestie: listă (după WHEN)");
        suggestions.push(createSnippet("list", "id1, id2", "Listă de identificatori"));
        return suggestions;
    }
    // Pattern: expresie incompletă
    if (childTypes.includes('expression') || childTypes.includes('simple_expression')) {
        console.log("💡 Sugestii: operatori pentru expresie");
        suggestions.push(...getExpressionContinuations());
        return suggestions;
    }
    // Fallback: sugestii generale
    return getGeneralSuggestions();
}
function getSuggestionsForNormalContext(node, document, position) {
    const suggestions = [];
    // Bazat pe tipul nodului curent
    switch (node.type) {
        case 'identifier':
            suggestions.push(...getIdentifierContinuations());
            break;
        case 'expression':
        case 'simple_expression':
        case 'binary_expression':
            suggestions.push(...getExpressionContinuations());
            break;
        case 'lambda':
            suggestions.push(createKeywordItem("TO", "TO "));
            break;
        case 'list':
            if (node.parent?.type === 'apply') {
                suggestions.push(createItem(";", "Termină statement-ul", vscode.CompletionItemKind.Text));
            }
            else if (node.parent?.type === 'when_operation') {
                suggestions.push(createKeywordItem("EXISTS", "EXISTS "));
            }
            break;
        case 'TO':
            suggestions.push(createSnippet("list", "id1, id2", "Listă de identificatori"));
            break;
        case 'WHEN':
            suggestions.push(createSnippet("list", "id1, id2", "Listă de identificatori"));
            break;
        case 'EXISTS':
            suggestions.push(...getExpressionSuggestions());
            break;
        default:
            suggestions.push(...getGeneralSuggestions());
    }
    return suggestions;
}
function getIdentifierContinuations() {
    return [
        createItem("=", "Atribuire", vscode.CompletionItemKind.Operator),
        createItem("==", "Egalitate", vscode.CompletionItemKind.Operator),
        createItem("!=", "Diferit", vscode.CompletionItemKind.Operator),
        createItem(">", "Mai mare", vscode.CompletionItemKind.Operator),
        createItem("<", "Mai mic", vscode.CompletionItemKind.Operator),
        createItem("~", "Regex match", vscode.CompletionItemKind.Operator),
        createItem("!~", "Regex not match", vscode.CompletionItemKind.Operator),
        createKeywordItem("IN", "IN [list]"),
        createItem(",", "Separator listă", vscode.CompletionItemKind.Text),
    ];
}
function getExpressionContinuations() {
    return [
        createItem("?", "Operator ternar (condition ? true : false)", vscode.CompletionItemKind.Operator),
        createItem("&&", "AND logic", vscode.CompletionItemKind.Operator),
        createItem("||", "OR logic", vscode.CompletionItemKind.Operator),
        createItem("==", "Egalitate", vscode.CompletionItemKind.Operator),
        createItem("!=", "Diferit", vscode.CompletionItemKind.Operator),
        createItem(">", "Mai mare", vscode.CompletionItemKind.Operator),
        createItem(">=", "Mai mare sau egal", vscode.CompletionItemKind.Operator),
        createItem("<", "Mai mic", vscode.CompletionItemKind.Operator),
        createItem("<=", "Mai mic sau egal", vscode.CompletionItemKind.Operator),
        createItem("+", "Adunare", vscode.CompletionItemKind.Operator),
        createItem("-", "Scădere", vscode.CompletionItemKind.Operator),
        createItem("*", "Înmulțire", vscode.CompletionItemKind.Operator),
        createItem("/", "Împărțire", vscode.CompletionItemKind.Operator),
        createItem(";", "Termină statement-ul", vscode.CompletionItemKind.Text),
    ];
}
function getExpressionSuggestions() {
    const suggestions = [];
    // Constante
    const constants = [
        'nothing', 'fraud', 'cancelled', 'invalid', 'alert',
        'conflict', 'ungraded', 'obscured', 'absent', 'present',
        'excused', 'toolow'
    ];
    constants.forEach(c => {
        suggestions.push(createItem(c, `Constantă: ${c}`, vscode.CompletionItemKind.Constant));
    });
    // Funcții
    suggestions.push(createFunctionItem("MAX", "MAX(expr)"), createFunctionItem("MIN", "MIN(expr)"), createFunctionItem("ROUND", "ROUND(expr)"));
    // Operatori unari
    suggestions.push(createItem("!", "NOT logic", vscode.CompletionItemKind.Operator), createItem("-", "Negare", vscode.CompletionItemKind.Operator));
    return suggestions;
}
function getGeneralSuggestions() {
    return [
        createSnippet("apply", "APPLY ${1:x} => ${2:x + 1} TO ${3:list1}, ${4:list2};", "Statement APPLY complet"),
        createSnippet("when", "WHEN ${1:list1}, ${2:list2} EXISTS ${3:expression};", "Statement WHEN complet"),
        createSnippet("explain", "EXPLAIN ${1:expression};", "Statement EXPLAIN"),
        createKeywordItem("APPLY", "APPLY "),
        createKeywordItem("WHEN", "WHEN "),
        createKeywordItem("EXPLAIN", "EXPLAIN "),
        createKeywordItem("TO", "TO"),
        createKeywordItem("EXISTS", "EXISTS"),
        createKeywordItem("IN", "IN"),
        ...getExpressionSuggestions(),
    ];
}
// Helper functions pentru crearea item-urilor
function createItem(label, detail, kind) {
    const item = new vscode.CompletionItem(label, kind);
    item.detail = detail;
    return item;
}
function createKeywordItem(keyword, insertText) {
    const item = new vscode.CompletionItem(keyword, vscode.CompletionItemKind.Keyword);
    item.insertText = insertText || keyword;
    item.detail = `Keyword: ${keyword}`;
    return item;
}
function createSnippet(label, snippet, detail) {
    const item = new vscode.CompletionItem(label, vscode.CompletionItemKind.Snippet);
    item.insertText = new vscode.SnippetString(snippet);
    item.detail = detail;
    return item;
}
function createFunctionItem(name, signature) {
    const item = new vscode.CompletionItem(name, vscode.CompletionItemKind.Function);
    item.insertText = new vscode.SnippetString(`${name}($1)$0`);
    item.detail = signature;
    return item;
}
function deactivate() { }
//# sourceMappingURL=extension.js.map