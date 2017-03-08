package es6

import "testing"

func TestLex_NumericLiteral0(t *testing.T) {
	expected := []Token{
		Token{NumericLiteralToken, "0"},
	}
	js := "0"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_NumericLiteral1(t *testing.T) {
	expected := []Token{
		Token{NumericLiteralToken, "1"},
	}
	js := "1"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_NumericLiteral2(t *testing.T) {
	expected := []Token{
		Token{NumericLiteralToken, "10"},
	}
	js := "10"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_NumericLiteral3(t *testing.T) {
	expected := []Token{
		Token{NumericLiteralToken, "0xAB10"},
	}
	js := "0xAB10"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_NumericLiteral4(t *testing.T) {
	expected := []Token{
		Token{NumericLiteralToken, "0b0100"},
	}
	js := "0b0100"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_NumericLiteral5(t *testing.T) {
	expected := []Token{
		Token{NumericLiteralToken, "0O0005"},
	}
	js := "0O0005"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_NumericLiteral6(t *testing.T) {
	expected := []Token{
		Token{NumericLiteralToken, "-6"},
	}
	js := "-6"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_NumericLiteral7(t *testing.T) {
	expected := []Token{
		Token{NumericLiteralToken, "0.0007"},
	}
	js := "0.0007"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_NumericLiteral8(t *testing.T) {
	expected := []Token{
		Token{NumericLiteralToken, "8.08"},
	}
	js := "8.08"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_NumericLiteral9(t *testing.T) {
	expected := []Token{
		Token{NumericLiteralToken, "3e2"},
	}
	js := "3e2"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_NumericLiteral10(t *testing.T) {
	expected := []Token{
		Token{ErrorToken, "bad number syntax: \"1o\""},
	}
	js := "1o"
	expectedTokens(t, expected, Lex("", js, true))
}
