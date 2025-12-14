package tree_sitter_codesuggestion_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_codesuggestion "github.com/tree-sitter/tree-sitter-codesuggestion/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_codesuggestion.Language())
	if language == nil {
		t.Errorf("Error loading Codesuggestion grammar")
	}
}
