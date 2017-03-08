package es6

import "testing"

func TestLex_RegEx00(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWordToken, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{PunctuatorToken, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{RegExToken, "/abc/i"}, InputElementRegExp},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = /abc/i", true))
}

func TestLex_RegEx01(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWordToken, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{PunctuatorToken, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{ErrorToken, "regex did not close with '/' "}, InputElementRegExp},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = /abc", true))
}
