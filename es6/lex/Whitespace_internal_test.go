package lex

import "testing"

func TestLex_Whitespace1(t *testing.T) {
	js := " \t"
	expected := []Token{
		Token{WhiteSpace, js},
	}
	_, tokens := Lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_Whitespace2(t *testing.T) {
	js := " "
	expected := []Token{
		Token{WhiteSpace, js},
	}

	_, tokens := Lex("", js, true)
	expectedTokens(t, expected, tokens)
}
