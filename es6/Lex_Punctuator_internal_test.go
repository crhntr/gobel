package es6

import "testing"

func TestLex_Punctuator1(t *testing.T) {
	expected := []Token{}
	js := ""
	ws := " "

	for _, punct := range punctuators {
		expected = append(expected, Token{PunctuatorToken, punct})
		expected = append(expected, Token{WhiteSpaceToken, ws})
		js += punct + ws
	}

	l := Lex("", js, false)
	expectedTokens(t, expected, l)
}

func TestLex_DivPunctuator1(t *testing.T) {
	expected := []Token{
		Token{IdentifierNameToken, "i"},
		Token{DivPunctuatorToken, "/="},
		Token{IdentifierNameToken, "j"},
		Token{DivPunctuatorToken, "/"},
		Token{NumericLiteralToken, "2"},
	}
	js := "i/=j/2"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLex_RightBracePunctuator1(t *testing.T) {
	expected := []Token{
		Token{PunctuatorToken, "{"},
		Token{RightBracePunctuatorToken, "}"},
	}
	js := "{}"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}
