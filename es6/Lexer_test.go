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
		expected := es6.Token{es6.ReservedWord, "var"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{es6.WhiteSpace, " "}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{es6.IdentifierName, "foo"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{es6.WhiteSpace, " "}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{es6.Punctuator, "="}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{es6.WhiteSpace, " "}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{es6.NumericLiteral, "123"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
}

func TestLexer_Next02(t *testing.T) {
	l := es6.Lex("", "var foo = `\\u006d\\x70`", true)

	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{es6.ReservedWord, "var"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{es6.WhiteSpace, " "}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{es6.IdentifierName, "foo"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{es6.WhiteSpace, " "}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{es6.Punctuator, "="}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{es6.WhiteSpace, " "}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{es6.NoSubstitutionTemplate, "`\\u006d\\x70`"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
}

func TestLexer_Next03(t *testing.T) {
	l := es6.Lex("", "i", true)

	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{es6.IdentifierName, "i"}
		if !next.Equals(expected) {
			t.Errorf("expected token: %s, but got %s", expected, next)
		}
	}
	{
		next := l.Next(es6.InputElementDiv)
		expected := es6.Token{es6.EOF, ""}
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
