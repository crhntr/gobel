package es6lexer_test

import (
	"fmt"
	"testing"

	"github.com/crhntr/gobel/es6/es6lexer"
)

func TestLexer_Next01(t *testing.T) {
	l := es6lexer.Lex("", "var foo = 123", true)

	{
		next := l.Next(es6lexer.InputElementDiv)
		expected := es6lexer.Token{es6lexer.ReservedWord, "var"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6lexer.InputElementDiv)
		expected := es6lexer.Token{es6lexer.WhiteSpace, " "}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6lexer.InputElementDiv)
		expected := es6lexer.Token{es6lexer.IdentifierName, "foo"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6lexer.InputElementDiv)
		expected := es6lexer.Token{es6lexer.WhiteSpace, " "}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6lexer.InputElementDiv)
		expected := es6lexer.Token{es6lexer.Punctuator, "="}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6lexer.InputElementDiv)
		expected := es6lexer.Token{es6lexer.WhiteSpace, " "}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6lexer.InputElementDiv)
		expected := es6lexer.Token{es6lexer.NumericLiteral, "123"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
}

func TestLexer_Next02(t *testing.T) {
	l := es6lexer.Lex("", "var foo = `\\u006d\\x70`", true)

	{
		next := l.Next(es6lexer.InputElementDiv)
		expected := es6lexer.Token{es6lexer.ReservedWord, "var"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6lexer.InputElementDiv)
		expected := es6lexer.Token{es6lexer.WhiteSpace, " "}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6lexer.InputElementDiv)
		expected := es6lexer.Token{es6lexer.IdentifierName, "foo"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6lexer.InputElementDiv)
		expected := es6lexer.Token{es6lexer.WhiteSpace, " "}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6lexer.InputElementDiv)
		expected := es6lexer.Token{es6lexer.Punctuator, "="}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6lexer.InputElementDiv)
		expected := es6lexer.Token{es6lexer.WhiteSpace, " "}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6lexer.InputElementDiv)
		expected := es6lexer.Token{es6lexer.NoSubstitutionTemplate, "`\\u006d\\x70`"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
}

func TestLexerGoal_String(t *testing.T) {
	if fmt.Sprintf("%s", es6lexer.InputElementDiv) == "" {
		t.Fail()
	}
	if fmt.Sprintf("%s", es6lexer.InputElementRegExp) == "" {
		t.Fail()
	}
	if fmt.Sprintf("%s", es6lexer.InputElementRegExpOrTemplateTail) == "" {
		t.Fail()
	}
	if fmt.Sprintf("%s", es6lexer.InputElementTemplateTail) == "" {
		t.Fail()
	}
	if fmt.Sprintf("%s", es6lexer.LexerGoal(-1)) == "" {
		t.Fail()
	}
}
