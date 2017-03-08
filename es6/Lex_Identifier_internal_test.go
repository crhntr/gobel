package es6

import "testing"

func TestLex_Identifier1(t *testing.T) {
	expected := []Token{
		Token{IdentifierNameToken, "$"},
		Token{PunctuatorToken, "="},
		Token{IdentifierNameToken, "_"},
		Token{PunctuatorToken, "="},
		Token{IdentifierNameToken, "foo"},
	}
	js := "$=_=foo"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_Identifier2(t *testing.T) {
	expected := []Token{
		Token{IdentifierNameToken, "X"},
		Token{PunctuatorToken, "&"},
		Token{IdentifierNameToken, "ooooooooooooo___"},
		Token{LineTerminatorToken, "\n"},
	}
	js := "X&ooooooooooooo___\n"
	expectedTokens(t, expected, Lex("", js, true))
}
