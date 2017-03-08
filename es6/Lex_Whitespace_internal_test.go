package es6

import "testing"

func TestLex_Whitespace1(t *testing.T) {
	js := " \t"
	expected := []Token{
		Token{WhiteSpaceToken, js},
	}
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_Whitespace2(t *testing.T) {
	js := " "
	expected := []Token{
		Token{WhiteSpaceToken, js},
	}
	expectedTokens(t, expected, Lex("", js, true))
}
