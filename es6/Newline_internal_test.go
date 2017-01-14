package es6

import "testing"

func TestLex_LineTerminator1(t *testing.T) {
	js := "\n"
	expected := []Token{
		Token{LineTerminator, js},
	}

	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_LineTerminator2(t *testing.T) {
	js := "\n\n"
	expected := []Token{
		Token{LineTerminator, "\n"},
		Token{LineTerminator, "\n"},
	}

	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}
