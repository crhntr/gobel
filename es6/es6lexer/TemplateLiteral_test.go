package es6lexer

import "testing"

func TestLex_TemplateLiteral01(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "var"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "foo"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "="},
		Token{WhiteSpace, " "},
		Token{NoSubstitutionTemplate, "`foo`"},
	}
	js := "var foo = `foo`"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_TemplateLiteral02(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "var"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "foo"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "="},
		Token{WhiteSpace, " "},
		Token{NoSubstitutionTemplate, "``"},
	}
	js := "var foo = ``"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_TemplateLiteral03(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "var"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "foo"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "="},
		Token{WhiteSpace, " "},
		Token{Error, "did not reach end of template literal reached eof"},
	}
	js := "var foo = `"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_TemplateLiteral04(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "var"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "foo"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "="},
		Token{WhiteSpace, " "},
		Token{TemplateHead, "`Hello ${"},
		Token{IdentifierName, "friend"},
	}
	js := "var foo = `Hello ${friend"
	expectedTokens(t, expected, Lex("", js, true))
}
