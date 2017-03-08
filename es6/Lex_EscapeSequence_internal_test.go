package es6

import "testing"

func TestLex_lexEscapeSequence01(t *testing.T) {
	expected := []Token{
		Token{StringLiteralToken, "\"\\u0074\\x61z\nzz\""},
	}
	js := "\"\\u0074\\x61z\nzz\""
	expectedTokens(t, expected, Lex("", js, true))
}
