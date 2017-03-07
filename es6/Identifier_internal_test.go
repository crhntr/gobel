package es6

import "testing"

func TestLex_Identifier1(t *testing.T) {
	expected := []Token{
		Token{IdentifierName, "$"},
		Token{Punctuator, "="},
		Token{IdentifierName, "_"},
		Token{Punctuator, "="},
		Token{IdentifierName, "foo"},
	}
	js := "$=_=foo"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_Identifier2(t *testing.T) {
	expected := []Token{
		Token{IdentifierName, "X"},
		Token{Punctuator, "&"},
		Token{IdentifierName, "ooooooooooooo___"},
		Token{LineTerminator, "\n"},
	}
	js := "X&ooooooooooooo___\n"
	expectedTokens(t, expected, Lex("", js, true))
}
