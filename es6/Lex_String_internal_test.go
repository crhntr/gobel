package es6

import "testing"

func TestLex_StringLiteralSingleQuote1(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{StringLiteralToken, "'foo'"},
	}
	js := "var foo = 'foo'"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_StringLiteralSingleQuote2(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{ErrorToken, "did not reach end of string literal reached eof"},
		Token{StringLiteralToken, "'foo"},
	}
	js := "var foo = 'foo"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_StringLiteralDoubleQuote1(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{StringLiteralToken, "\"foo\""},
	}
	js := "var foo = \"foo\""
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_StringLiteralDoubleQuote2(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{ErrorToken, "did not reach end of string literal reached eof"},
		Token{StringLiteralToken, "\"foo"},
	}
	js := "var foo = \"foo"
	expectedTokens(t, expected, Lex("", js, true))
}
