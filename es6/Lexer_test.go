package es6_test

import (
	"fmt"
	"testing"

	"github.com/crhntr/gobel/es6"
)

func TestLexer_Next01(t *testing.T) {
	l := es6.Lex("", "var foo = 123", true)

	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{Type: es6.ReservedWordToken, Value: "var"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{Type: es6.IdentifierNameToken, Value: "foo"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{Type: es6.PunctuatorToken, Value: "="}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{Type: es6.NumericLiteralToken, Value: "123"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
}

func TestLexer_Next02(t *testing.T) {
	l := es6.Lex("", "var foo = `\\u006d\\x70`", true)

	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{Type: es6.ReservedWordToken, Value: "var"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{Type: es6.IdentifierNameToken, Value: "foo"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{Type: es6.PunctuatorToken, Value: "="}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{Type: es6.NoSubstitutionTemplateToken, Value: "`\\u006d\\x70`"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
}

func TestLexer_Next03(t *testing.T) {
	l := es6.Lex("", "i", true)

	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{Type: es6.IdentifierNameToken, Value: "i"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{Type: es6.EOFToken, Value: ""}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
}

func TestLexerGoal_String(t *testing.T) {
	if fmt.Sprintf("%s", es6.InputElementDiv) == "" {
		t.Fail()
	}
	if fmt.Sprintf("%s", es6.InputElementRegExp) == "" {
		t.Fail()
	}
	if fmt.Sprintf("%s", es6.InputElementRegExpOrTemplateTail) == "" {
		t.Fail()
	}
	if fmt.Sprintf("%s", es6.InputElementTemplateTail) == "" {
		t.Fail()
	}
	if fmt.Sprintf("%s", es6.LexerGoal(-1)) == "" {
		t.Fail()
	}
}

func TestLexer_Peek(t *testing.T) {
	l := es6.Lex("", "var foo = 123", true)

	{
		next := l.Peek(es6.InputElementDiv)
		expected := es6.Token{Type: es6.ReservedWordToken, Value: "var"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{Type: es6.ReservedWordToken, Value: "var"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Peek(es6.InputElementDiv)
		expected := es6.Token{Type: es6.IdentifierNameToken, Value: "foo"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{Type: es6.IdentifierNameToken, Value: "foo"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Peek(es6.InputElementDiv)
		expected := es6.Token{Type: es6.PunctuatorToken, Value: "="}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{Type: es6.PunctuatorToken, Value: "="}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Peek(es6.InputElementDiv)
		expected := es6.Token{Type: es6.NumericLiteralToken, Value: "123"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{Type: es6.NumericLiteralToken, Value: "123"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
}
