package es6lexer

import "testing"

func TestLex_StringLiteralSingleQuote1(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "var"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "foo"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "="},
		Token{WhiteSpace, " "},
		Token{StringLiteral, "'foo'"},
	}
	js := "var foo = 'foo'"
	_, tokens := Lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_StringLiteralSingleQuote2(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "var"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "foo"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "="},
		Token{WhiteSpace, " "},
		Token{Error, "did not reach end of string literal reached eof"},
		Token{StringLiteral, "'foo"},
	}
	js := "var foo = 'foo"
	_, tokens := Lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_StringLiteralDoubleQuote1(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "var"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "foo"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "="},
		Token{WhiteSpace, " "},
		Token{StringLiteral, "\"foo\""},
	}
	js := "var foo = \"foo\""
	_, tokens := Lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_StringLiteralDoubleQuote2(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "var"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "foo"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "="},
		Token{WhiteSpace, " "},
		Token{Error, "did not reach end of string literal reached eof"},
		Token{StringLiteral, "\"foo"},
	}
	js := "var foo = \"foo"
	_, tokens := Lex("", js, true)
	expectedTokens(t, expected, tokens)
}
