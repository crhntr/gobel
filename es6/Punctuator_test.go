package es6

import "testing"

func TestLex_Punctuator1(t *testing.T) {
	expected := []Token{}
	js := ""
	ws := " "

	for _, punct := range punctuators {
		expected = append(expected, Token{Punctuator, punct})
		expected = append(expected, Token{WhiteSpace, ws})
		js += punct + ws
	}

	l := Lex("", js, false)
	expectedTokens(t, expected, l)
}

func TestLex_DivPunctuator1(t *testing.T) {
	expected := []Token{
		Token{IdentifierName, "i"},
		Token{DivPunctuator, "/="},
		Token{IdentifierName, "j"},
		Token{DivPunctuator, "/"},
		Token{NumericLiteral, "2"},
	}
	js := "i/=j/2"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLex_RightBracePunctuator1(t *testing.T) {
	expected := []Token{
		Token{Punctuator, "{"},
		Token{RightBracePunctuator, "}"},
	}
	js := "{}"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}
