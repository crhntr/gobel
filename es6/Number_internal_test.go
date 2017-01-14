package es6

import "testing"

func TestLex_NumericLiteral0(t *testing.T) {
	expected := []Token{
		Token{NumericLiteral, "0"},
	}
	js := "0"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_NumericLiteral1(t *testing.T) {
	expected := []Token{
		Token{NumericLiteral, "1"},
	}
	js := "1"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_NumericLiteral2(t *testing.T) {
	expected := []Token{
		Token{NumericLiteral, "10"},
	}
	js := "10"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_NumericLiteral3(t *testing.T) {
	expected := []Token{
		Token{NumericLiteral, "0xAB10"},
	}
	js := "0xAB10"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_NumericLiteral4(t *testing.T) {
	expected := []Token{
		Token{NumericLiteral, "0b0100"},
	}
	js := "0b0100"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_NumericLiteral5(t *testing.T) {
	expected := []Token{
		Token{NumericLiteral, "0O0005"},
	}
	js := "0O0005"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_NumericLiteral6(t *testing.T) {
	expected := []Token{
		Token{NumericLiteral, "-6"},
	}
	js := "-6"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_NumericLiteral7(t *testing.T) {
	expected := []Token{
		Token{NumericLiteral, "0.0007"},
	}
	js := "0.0007"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_NumericLiteral8(t *testing.T) {
	expected := []Token{
		Token{NumericLiteral, "8.08"},
	}
	js := "8.08"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_NumericLiteral9(t *testing.T) {
	expected := []Token{
		Token{NumericLiteral, "3e2"},
	}
	js := "3e2"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_NumericLiteral10(t *testing.T) {
	expected := []Token{
		Token{Error, "bad number syntax: \"1o\""},
	}
	js := "1o"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}
