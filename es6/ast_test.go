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
			t.Error("identifierNode.Name should be %q but got %q", justAnIdentifier, identifierNode.Name)
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
