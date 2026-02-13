package ast

import (
	"new-er/keyboard-cheatsheet/data/ast/nodes"
	"testing"
)

var tests = []struct {
	input    string
	expected nodes.Node
}{
	{"/", nodes.NewKeyNode("/")},
	{"CTRL+T", nodes.NewComboNode(nodes.NewKeyNode("CTRL"), nodes.NewKeyNode("T"))},
	{"c+b c", nodes.NewSequenceNode(
		nodes.NewComboNode(nodes.NewKeyNode("c"), nodes.NewKeyNode("b")),
		nodes.NewKeyNode("c"),
	),
	},
	/*{"/", nodes.Key{Name: "/"}},
	{"CTRL+T", nodes.ComboNode{Keys: []nodes.Node{nodes.Key{Name: "CTRL"}, nodes.Key{Name: "T"}}}},
	{"ALT+(LEFT|RIGHT)", nodes.ComboNode{Keys: []nodes.Node{nodes.Key{Name: "ALT"}, nodes.AnyOfNode{Keys: []nodes.Node{nodes.Key{Name: "LEFT"}, nodes.Key{Name: "RIGHT"}}}}}},
	{"TAB+SHIFT?", nodes.ComboNode{Keys: []nodes.Node{nodes.Key{Name: "TAB"}, nodes.OptionalNode{nodes.Key{Name: "SHIFT"}}}}},
	{"SHIFT+5:%", nodes.ComboNode{Keys: []nodes.Node{nodes.Key{Name: "SHIFT"}, nodes.AliasNode{Key: nodes.Key{Name: "5"}, Alias: "%"}}}},*/
}

func TestParser(t *testing.T) {
	for _, tt := range tests {
		lexer := NewLexer(tt.input)
		tokens := lexer.AllTokens()

		parser := NewParser(tokens)
		node := parser.Parse()

		if !node.Equals(tt.expected) {
			t.Errorf("Expected node to be non-nil")
		}
	}
}
