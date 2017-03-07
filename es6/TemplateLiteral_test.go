package es6lexer

import (
	"testing"
)

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

func TestLex_TemplateLiteral05(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWord, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{IdentifierName, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{Punctuator, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{TemplateHead, "`Hello ${"}, InputElementDiv},
		TokenTest{Token{IdentifierName, "friend"}, InputElementDiv},
		TokenTest{Token{TemplateTail, "}!`"}, InputElementRegExpOrTemplateTail},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = `Hello ${friend}!`", true))
}

func TestLex_TemplateLiteral06(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWord, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{IdentifierName, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{Punctuator, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{TemplateHead, "`Hello ${"}, InputElementDiv},
		TokenTest{Token{IdentifierName, "friend"}, InputElementDiv},
		TokenTest{Token{Error, "did not reach TemplateMiddle or TemplateTail but reached eof"}, InputElementRegExpOrTemplateTail},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = `Hello ${friend}! ", true))
}

func TestLex_TemplateLiteral07(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWord, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{IdentifierName, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{Punctuator, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{TemplateHead, "`Hello ${"}, InputElementDiv},
		TokenTest{Token{IdentifierName, "friend"}, InputElementDiv},
		TokenTest{Token{TemplateMiddle, "}! ${"}, InputElementRegExpOrTemplateTail},
		TokenTest{Token{TemplateHead, "` ${"}, InputElementDiv},
		TokenTest{Token{NumericLiteral, "4"}, InputElementDiv},
		TokenTest{Token{TemplateTail, "}`"}, InputElementRegExpOrTemplateTail},
		TokenTest{Token{Punctuator, "+"}, InputElementDiv},
		TokenTest{Token{TemplateHead, "`${"}, InputElementDiv},
		TokenTest{Token{NumericLiteral, "2"}, InputElementDiv},
		TokenTest{Token{TemplateTail, "} `"}, InputElementRegExpOrTemplateTail},
		TokenTest{Token{TemplateTail, "}`"}, InputElementRegExpOrTemplateTail},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = `Hello ${friend}! ${` ${4}`+`${2} `}`", true))
}

func TestLex_TemplateLiteral08(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWord, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{IdentifierName, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{Punctuator, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpace, " "}, InputElementDiv},
		TokenTest{Token{TemplateHead, "`Hello ${"}, InputElementDiv},
		TokenTest{Token{IdentifierName, "friend"}, InputElementDiv},
		TokenTest{Token{TemplateMiddle, "}! ${"}, InputElementTemplateTail},
		TokenTest{Token{TemplateHead, "` ${"}, InputElementDiv},
		TokenTest{Token{NumericLiteral, "4"}, InputElementDiv},
		TokenTest{Token{TemplateTail, "}`"}, InputElementTemplateTail},
		TokenTest{Token{Punctuator, "+"}, InputElementDiv},
		TokenTest{Token{TemplateHead, "`${"}, InputElementDiv},
		TokenTest{Token{NumericLiteral, "2"}, InputElementDiv},
		TokenTest{Token{TemplateTail, "} `"}, InputElementTemplateTail},
		TokenTest{Token{TemplateTail, "}`"}, InputElementTemplateTail},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = `Hello ${friend}! ${` ${4}`+`${2} `}`", true))
}
