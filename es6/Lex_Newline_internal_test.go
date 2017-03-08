package es6

import "testing"

func TestLex_LineTerminator1(t *testing.T) {
	js := "\n"
	expected := []Token{
		Token{LineTerminatorToken, js},
	}
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_LineTerminator2(t *testing.T) {
	js := "\n\n"
	expected := []Token{
		Token{LineTerminatorToken, "\n"},
		Token{LineTerminatorToken, "\n"},
	}
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_LineTerminator3(t *testing.T) {
	js := "\u000A\u000D\u2028\u2029"
	expected := []Token{
		Token{LineTerminatorToken, "\u000A"},
		Token{LineTerminatorToken, "\u000D"},
		Token{LineTerminatorToken, "\u2028"},
		Token{LineTerminatorToken, "\u2029"},
	}
	expectedTokens(t, expected, Lex("", js, true))
}
