package lex

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

func TestLex_LineTerminator3(t *testing.T) {
	js := "\u000A\u000D\u2028\u2029"
	expected := []Token{
		Token{LineTerminator, "\u000A"},
		Token{LineTerminator, "\u000D"},
		Token{LineTerminator, "\u2028"},
		Token{LineTerminator, "\u2029"},
	}

	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}
