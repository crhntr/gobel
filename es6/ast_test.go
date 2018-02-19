package es6_test

import (
	"testing"

	"github.com/crhntr/gobel/es6"
)

func TestParseIdentifierNode(t *testing.T) {

	t.Run("should return IdentifierNode token", func(t *testing.T) {
		justAnIdentifier := "foo"
		lex := es6.Lex("", justAnIdentifier, false)
		node, err := es6.ParseIdentifierNode(lex)
		if err != nil {
			t.Error(err)
		}
		identifierNode, ok := node.(es6.IdentifierNode)
		if !ok {
			t.Fail()
		}
		if identifierNode.Name != justAnIdentifier {
			t.Errorf("identifierNode.Name should be %q but got %q", justAnIdentifier, identifierNode.Name)
		}
	})

	t.Run("should not allow a ReservedWord", func(t *testing.T) {
		aReservedWord := "import"
		lex := es6.Lex("", aReservedWord, false)
		_, err := es6.ParseIdentifierNode(lex)
		if err == nil {
			t.Fail()
		}
	})
}

func TestParseExportsListNode(t *testing.T) {
	t.Run("should allow as Identifier", func(t *testing.T) {
		foo := "foo"
		lex := es6.Lex("", foo, false)
		node, err := es6.ParseExportsListNode(lex)
		if err != nil {
			t.Error(err)
		}

		n, ok := node.(es6.ExportsListNode)
		if !ok {
			t.Error(`!ok`)
		}
		if len(n.List) != 1 {
			t.Error(`len(n.List) != 3"`)
		}
		for i, str := range []string{"foo"} {
			if n.List[i].Name != str {
				t.Errorf("n.List[%d].Name != %q", i, str)
			}
		}
	})

	t.Run("should varify Exports name is acceptable", func(t *testing.T) {
		fooBarBaz := "foo, for"
		lex := es6.Lex("", fooBarBaz, false)
		_, err := es6.ParseExportsListNode(lex)
		if err == nil {
			t.Error("err == nil")
		}
	})

	t.Run("should allow as Identifier", func(t *testing.T) {
		fooBarBaz := "foo, bar, baz"
		lex := es6.Lex("", fooBarBaz, false)
		node, err := es6.ParseExportsListNode(lex)
		if err != nil {
			t.Error(err)
		}

		n, ok := node.(es6.ExportsListNode)
		if !ok {
			t.Error(`!ok`)
		}
		if len(n.List) != 3 {
			t.Error(`len(n.List) != 3"`)
		}
		for i, str := range []string{"foo", "bar", "baz"} {
			if n.List[i].Name != str {
				t.Errorf("n.List[%d].Name != %q", i, str)
			}
		}
	})
}

func TestParseExportSpecifierNode(t *testing.T) {
	t.Run("should allow as Identifier", func(t *testing.T) {
		fooAsBar := "foo as bar"
		lex := es6.Lex("", fooAsBar, false)

		node, err := es6.ParseExportSpecifierNode(lex)
		if err != nil {
			t.Error(err)
		}
		n, ok := node.(es6.ExportSpecifierNode)
		if !ok {
			t.Error(`!ok`)
		}
		if n.Name != "foo" {
			t.Error(`n.Name != "foo"`)
		}
		if n.As.Name != "bar" {
			t.Error(`n.As.Name != "bar"`)
		}
	})

	t.Run("should allow as Identifier", func(t *testing.T) {
		fooAsBar := " foo "
		lex := es6.Lex("", fooAsBar, false)

		node, err := es6.ParseExportSpecifierNode(lex)
		if err != nil {
			t.Error(err)
		}
		n, ok := node.(es6.ExportSpecifierNode)
		if !ok {
			t.Error(`!ok`)
		}
		if n.Name != "foo" {
			t.Error(`n.Name != "foo"`)
		}
	})

	t.Run("should allow as Identifier", func(t *testing.T) {
		fooAsBar := " for "
		lex := es6.Lex("", fooAsBar, false)

		_, err := es6.ParseExportSpecifierNode(lex)
		if err == nil {
			t.Error("should now allow reserved word as Name")
		}
	})
}

func TestParseLetOrConstNode(t *testing.T) {
	t.Run("should set node value to 'const'", func(t *testing.T) {
		cnst := "const"
		lex := es6.Lex("", cnst, false)

		node, err := es6.ParseLetOrConstNode(lex)
		if err != nil {
			t.Error(err)
		}
		n, ok := node.(es6.LetOrConstNode)
		if !ok {
			t.Error(`!ok`)
		}
		if n.Value != "const" {
			t.Error(`n.Value != "const"`)
		}
	})

	t.Run("should set node value to 'let'", func(t *testing.T) {
		lt := "let"
		lex := es6.Lex("", lt, false)

		node, err := es6.ParseLetOrConstNode(lex)
		if err != nil {
			t.Error(err)
		}
		n, ok := node.(es6.LetOrConstNode)
		if !ok {
			t.Error(`!ok`)
		}
		if n.Value != "let" {
			t.Error(`n.Value != "let"`)
		}
	})

	t.Run("should not recognize other than foo or bar", func(t *testing.T) {
		lt := "foo"
		lex := es6.Lex("", lt, false)

		_, err := es6.ParseLetOrConstNode(lex)
		if err == nil {
			t.Fail()
		}
	})
}
