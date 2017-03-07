package es6lexer

import "testing"

func TestLex_Whitespace1(t *testing.T) {
	js := " \t"
	expected := []Token{
		Token{WhiteSpace, js},
	}
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_Whitespace2(t *testing.T) {
	js := " "
	expected := []Token{
		Token{WhiteSpace, js},
	}
	expectedTokens(t, expected, Lex("", js, true))
}
