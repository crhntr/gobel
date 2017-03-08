package es6

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var errNotNull = fmt.Errorf("not null")

func Test_peek(t *testing.T) {
	l := Lex("", "101", false)

	n1 := l.next()
	p0 := l.peek()
	n0 := l.next()

	if p0 != n0 && p0 == '0' && n1 != '1' {
		t.Error("peek is broken")
	}
}

func TestLex_Whitespace_AND_SingleLineComment(t *testing.T) {
	expected := []Token{
		Token{WhiteSpaceToken, " \t"},
		Token{SingleLineCommentToken, " Hello World!"},
	}
	js := " \t// Hello World!"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLex_Terminator_And_Whitespace(t *testing.T) {
	js := "\n\t\n"
	expected := []Token{
		Token{LineTerminatorToken, "\n"},
		Token{WhiteSpaceToken, "\t"},
		Token{LineTerminatorToken, "\n"},
	}

	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLex_ReservedWord1(t *testing.T) {
	expected := []Token{}
	js := ""
	ws := " "

	lJs := Lexer{}
	lJs.setStrict()
	for _, word := range lJs.reservedWords {
		js += word + ws
	}
	l := Lex("", js, true)
	for _, word := range l.reservedWords {
		expected = append(expected, Token{ReservedWordToken, word})
		expected = append(expected, Token{WhiteSpaceToken, ws})
	}

	expectedTokens(t, expected, l)
}

func TestLex_ReservedWord2(t *testing.T) {
	expected := []Token{}
	js := ""
	ws := " "

	lJs := Lexer{}
	lJs.unsetStrict()
	for _, word := range lJs.reservedWords {
		js += word + ws
	}

	l := Lex("", js, false)
	for _, word := range l.reservedWords {
		expected = append(expected, Token{ReservedWordToken, word})
		expected = append(expected, Token{WhiteSpaceToken, ws})
	}

	expectedTokens(t, expected, l)
}

// func TestLex_EscapeSequence0(t *testing.T) {
// 	expected := []Token{
// 		Token{IdentifierName, "X"},
// 		Token{Punctuator, "&"},
// 		Token{IdentifierName, "ooooooooooooo___"},
// 		Token{LineTerminator, "\n"},
// 	}
// 	js := `"\n"`
// 	l := lex("", js, true)
// 	expectedTokens(t, expected, l)
// }

func TestLex1(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "function"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "("},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, "{"},
		Token{RightBracePunctuatorToken, "}"},
		Token{PunctuatorToken, "("},
		Token{PunctuatorToken, ")"},
	}

	js := "function (){}()"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLex2(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "function"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "("},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, "{"},
		Token{RightBracePunctuatorToken, "}"},
		Token{PunctuatorToken, "("},
		Token{PunctuatorToken, ")"},
	}

	js := "function foo (){}()"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLex3(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "function"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "add"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "a"},
		Token{PunctuatorToken, ","},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "b"},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, "{"},
		Token{LineTerminatorToken, "\n"},
		Token{WhiteSpaceToken, "\t"},
		Token{ReservedWordToken, "return"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "a"},
		Token{PunctuatorToken, "+"},
		Token{IdentifierNameToken, "b"},
		Token{LineTerminatorToken, "\n"},
		Token{RightBracePunctuatorToken, "}"},
	}

	js := "function add (a, b){\n\treturn a+b\n}"
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLexJS(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "function"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "fibonacci"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{PunctuatorToken, ")"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "{"},
		Token{LineTerminatorToken, "\n"},
		Token{WhiteSpaceToken, "  "},
		Token{ReservedWordToken, "if"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, ">="},
		Token{WhiteSpaceToken, " "},
		Token{NumericLiteralToken, "2"},
		Token{PunctuatorToken, ")"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "{"},
		Token{LineTerminatorToken, "\n"},
		Token{WhiteSpaceToken, "    "},
		Token{ReservedWordToken, "return"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "fibonacci"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{NumericLiteralToken, "-1"},
		Token{PunctuatorToken, ")"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "+"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "fibonacci"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{NumericLiteralToken, "-2"},
		Token{PunctuatorToken, ")"},
		Token{LineTerminatorToken, "\n"},
		Token{WhiteSpaceToken, "  "},
		Token{RightBracePunctuatorToken, "}"},
		Token{LineTerminatorToken, "\n"},
		Token{WhiteSpaceToken, "  "},
		Token{ReservedWordToken, "return"},
		Token{WhiteSpaceToken, " "},
		Token{NumericLiteralToken, "1"},
		Token{LineTerminatorToken, "\n"},
		Token{RightBracePunctuatorToken, "}"},
		Token{LineTerminatorToken, "\n"},
		Token{LineTerminatorToken, "\n"},
		Token{IdentifierNameToken, "console"},
		Token{PunctuatorToken, "."},
		Token{IdentifierNameToken, "log"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "fibonacci"},
		Token{PunctuatorToken, "("},
		Token{NumericLiteralToken, "7"},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, ")"},
		Token{LineTerminatorToken, "\n"},
	}

	testData, err := ioutil.ReadFile("testdata/index01.js")
	if err != nil {
		t.Fatal(err)
	}

	js := string(testData)
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLexJS2(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "function"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "fibonacci"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, "{"},
		Token{ReservedWordToken, "if"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{PunctuatorToken, ">="},
		Token{NumericLiteralToken, "2"},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, "{"},
		Token{ReservedWordToken, "return"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "fibonacci"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{NumericLiteralToken, "-1"},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, "+"},
		Token{IdentifierNameToken, "fibonacci"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "n"},
		Token{NumericLiteralToken, "-2"},
		Token{PunctuatorToken, ")"},
		Token{RightBracePunctuatorToken, "}"},
		Token{ReservedWordToken, "return"},
		Token{WhiteSpaceToken, " "},
		Token{NumericLiteralToken, "1"},
		Token{RightBracePunctuatorToken, "}"},
		Token{PunctuatorToken, ";"},
		Token{IdentifierNameToken, "console"},
		Token{PunctuatorToken, "."},
		Token{IdentifierNameToken, "log"},
		Token{PunctuatorToken, "("},
		Token{IdentifierNameToken, "fibonacci"},
		Token{PunctuatorToken, "("},
		Token{NumericLiteralToken, "7"},
		Token{PunctuatorToken, ")"},
		Token{PunctuatorToken, ")"},
		Token{LineTerminatorToken, "\n"},
	}

	testData, err := ioutil.ReadFile("testdata/index02.js")
	if err != nil {
		t.Fatal(err)
	}

	js := string(testData)
	l := Lex("", js, true)
	expectedTokens(t, expected, l)
}

func TestLexJS3(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWordToken, "function"}, InputElementRegExp},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementRegExp},
		TokenTest{Token{IdentifierNameToken, "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, "("}, InputElementRegExp},
		TokenTest{Token{IdentifierNameToken, "sequence"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, ")"}, InputElementRegExp},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, "{"}, InputElementRegExp},
		TokenTest{Token{LineTerminatorToken, "\n"}, InputElementRegExp},
		TokenTest{Token{WhiteSpaceToken, "  "}, InputElementRegExp},
		TokenTest{Token{ReservedWordToken, "return"}, InputElementRegExp},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementRegExp},
		TokenTest{Token{IdentifierNameToken, "valid"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, "."}, InputElementRegExp},
		TokenTest{Token{IdentifierNameToken, "match"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, "("}, InputElementRegExp},
		TokenTest{Token{RegExToken, "/([CGAT]{3}){1,}/g"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, ")"}, InputElementRegExp},
		TokenTest{Token{LineTerminatorToken, "\n"}, InputElementRegExp},
		TokenTest{Token{RightBracePunctuatorToken, "}"}, InputElementRegExp},
		TokenTest{Token{LineTerminatorToken, "\n"}, InputElementRegExp},
		TokenTest{Token{IdentifierNameToken, "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, "("}, InputElementRegExp},
		TokenTest{Token{StringLiteralToken, "\"ATATTGGTGTTCATGTGCGCGGGGCCGACGAGCTACTGGCAGAACCACGAGGACAAGAGGTGA\""}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, ")"}, InputElementRegExp},
		TokenTest{Token{LineTerminatorToken, "\n"}, InputElementRegExp},
		TokenTest{Token{IdentifierNameToken, "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, "("}, InputElementRegExp},
		TokenTest{Token{StringLiteralToken, "\"FAIL\""}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, ")"}, InputElementRegExp},
		TokenTest{Token{LineTerminatorToken, "\n"}, InputElementRegExp},
		TokenTest{Token{IdentifierNameToken, "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, "("}, InputElementRegExp},
		TokenTest{Token{StringLiteralToken, "\"Alanine\""}, InputElementRegExp},
		TokenTest{Token{PunctuatorToken, ")"}, InputElementRegExp},
		TokenTest{Token{LineTerminatorToken, "\n"}, InputElementRegExp},
		TokenTest{Token{EOFToken, ""}, InputElementDiv},
	}
	testData, err := ioutil.ReadFile("testdata/TestLexJS3.js")
	if err != nil {
		t.Fatal(err)
	}

	js := string(testData)
	l := Lex("", js, true)
	expectedTokensTable(t, expected, l)
}

func TestToken_String(t *testing.T) {
	t1 := Token{TokenType(-1), ""}
	if t1.String() == "" {
		t.Error("t1.String() string returns empty string")
	}
	t2 := Token{LineTerminatorToken, ""}
	if t2.String() == "" {
		t.Error("t2.String() string returns empty string")
	}
	t3 := Token{MultiLineCommentToken, " some string value "}
	if t3.String() == "" {
		t.Error("t3.String() string returns empty string")
	}
}

func expectedTokens(t *testing.T, expectedTokens []Token, l *Lexer) {
	i := 0
	for _, expected := range expectedTokens {
		tok := l.Next(InputElementDiv)
		t.Logf("%d: %s %s\n", i, expectedTokens[i], tok)
		if !expected.Equals(tok) {
			t.Errorf("expected and recived tokens do not match [%d](%s != %s)", i, tok, expectedTokens[i])
		}
		i++
	}
	if i < len(expectedTokens) {
		t.Errorf("expected more tokens (expected: %d, got %d)", len(expectedTokens), i)
	}
}

type TokenTest struct {
	Token Token
	Goal  LexerGoal
}

func expectedTokensTable(t *testing.T, expectedTokenRows []TokenTest, l *Lexer) {
	i := 0
	for _, expected := range expectedTokenRows {
		tok := l.Next(expected.Goal)
		t.Logf("%d: %s %s\n", i, expected.Token, tok)
		if !expected.Token.Equals(tok) {
			t.Errorf("expected and recived tokens do not match [%d](%s != %s)", i, expected.Token, tok)
		}
		i++
	}
	if i < len(expectedTokenRows) {
		t.Errorf("expected more tokens (expected: %d, got %d)", len(expectedTokenRows), i)
	}
}

//
// Test Comment
//

func TestLex_MultiLineComment1(t *testing.T) {
	expected := []Token{
		Token{MultiLineCommentToken, "Hello World!"},
	}
	js := "/*Hello World!*/"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_MultiLineComment2(t *testing.T) {
	expected := []Token{
		Token{ErrorToken, "no multi line comment terminator \"*/\""},
	}
	js := "/*Hello World!"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_MultiLineComment3(t *testing.T) {
	expected := []Token{
		Token{ErrorToken, "no multi line comment terminator \"*/\""},
	}
	js := "/* \""
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_SingleLineComment1(t *testing.T) {
	expected := []Token{
		Token{SingleLineCommentToken, " Hello World!"},
	}
	js := "// Hello World!"
	expectedTokens(t, expected, Lex("", js, true))
}
func TestLex_SingleLineComment2(t *testing.T) {
	expected := []Token{
		Token{SingleLineCommentToken, " Hello World!"},
	}
	js := "// Hello World!\n"
	expectedTokens(t, expected, Lex("", js, true))
}
func TestLex_Comments(t *testing.T) {
	expected := []Token{
		Token{SingleLineCommentToken, " Hello World!"},
		Token{MultiLineCommentToken, "This is a multi\nline comment "},
	}
	js := "// Hello World!\n/*This is a multi\nline comment */"
	expectedTokens(t, expected, Lex("", js, true))
}

// Test EscapeSequence
func TestLex_lexEscapeSequence01(t *testing.T) {
	expected := []Token{
		Token{StringLiteralToken, "\"\\u0074\\x61z\nzz\""},
	}
	js := "\"\\u0074\\x61z\nzz\""
	expectedTokens(t, expected, Lex("", js, true))
}

//
// Test Identifier
//
func TestLex_Identifier1(t *testing.T) {
	expected := []Token{
		Token{IdentifierNameToken, "$"},
		Token{PunctuatorToken, "="},
		Token{IdentifierNameToken, "_"},
		Token{PunctuatorToken, "="},
		Token{IdentifierNameToken, "foo"},
	}
	js := "$=_=foo"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_Identifier2(t *testing.T) {
	expected := []Token{
		Token{IdentifierNameToken, "X"},
		Token{PunctuatorToken, "&"},
		Token{IdentifierNameToken, "ooooooooooooo___"},
		Token{LineTerminatorToken, "\n"},
	}
	js := "X&ooooooooooooo___\n"
	expectedTokens(t, expected, Lex("", js, true))
}

//
// Test LineTerminator
//

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

//
// Test Number Literal
//

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

//
// Test Punctuator
//

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

//
// Test RegEx
//

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

//
// Test String
//

func TestLex_StringLiteralSingleQuote1(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{StringLiteralToken, "'foo'"},
	}
	js := "var foo = 'foo'"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_StringLiteralSingleQuote2(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{ErrorToken, "did not reach end of string literal reached eof"},
		Token{StringLiteralToken, "'foo"},
	}
	js := "var foo = 'foo"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_StringLiteralDoubleQuote1(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{StringLiteralToken, "\"foo\""},
	}
	js := "var foo = \"foo\""
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_StringLiteralDoubleQuote2(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{ErrorToken, "did not reach end of string literal reached eof"},
		Token{StringLiteralToken, "\"foo"},
	}
	js := "var foo = \"foo"
	expectedTokens(t, expected, Lex("", js, true))
}

//
// Test Template literal
//

func TestLex_TemplateLiteral01(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{NoSubstitutionTemplateToken, "`foo`"},
	}
	js := "var foo = `foo`"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_TemplateLiteral02(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{NoSubstitutionTemplateToken, "``"},
	}
	js := "var foo = ``"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_TemplateLiteral03(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{ErrorToken, "did not reach end of template literal reached eof"},
	}
	js := "var foo = `"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_TemplateLiteral04(t *testing.T) {
	expected := []Token{
		Token{ReservedWordToken, "var"},
		Token{WhiteSpaceToken, " "},
		Token{IdentifierNameToken, "foo"},
		Token{WhiteSpaceToken, " "},
		Token{PunctuatorToken, "="},
		Token{WhiteSpaceToken, " "},
		Token{TemplateHeadToken, "`Hello ${"},
		Token{IdentifierNameToken, "friend"},
	}
	js := "var foo = `Hello ${friend"
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_TemplateLiteral05(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWordToken, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{PunctuatorToken, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{TemplateHeadToken, "`Hello ${"}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "friend"}, InputElementDiv},
		TokenTest{Token{TemplateTailToken, "}!`"}, InputElementRegExpOrTemplateTail},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = `Hello ${friend}!`", true))
}

func TestLex_TemplateLiteral06(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWordToken, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{PunctuatorToken, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{TemplateHeadToken, "`Hello ${"}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "friend"}, InputElementDiv},
		TokenTest{Token{ErrorToken, "did not reach TemplateMiddle or TemplateTail but reached eof"}, InputElementRegExpOrTemplateTail},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = `Hello ${friend}! ", true))
}

func TestLex_TemplateLiteral07(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWordToken, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{PunctuatorToken, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{TemplateHeadToken, "`Hello ${"}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "friend"}, InputElementDiv},
		TokenTest{Token{TemplateMiddleToken, "}! ${"}, InputElementRegExpOrTemplateTail},
		TokenTest{Token{TemplateHeadToken, "` ${"}, InputElementDiv},
		TokenTest{Token{NumericLiteralToken, "4"}, InputElementDiv},
		TokenTest{Token{TemplateTailToken, "}`"}, InputElementRegExpOrTemplateTail},
		TokenTest{Token{PunctuatorToken, "+"}, InputElementDiv},
		TokenTest{Token{TemplateHeadToken, "`${"}, InputElementDiv},
		TokenTest{Token{NumericLiteralToken, "2"}, InputElementDiv},
		TokenTest{Token{TemplateTailToken, "} `"}, InputElementRegExpOrTemplateTail},
		TokenTest{Token{TemplateTailToken, "}`"}, InputElementRegExpOrTemplateTail},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = `Hello ${friend}! ${` ${4}`+`${2} `}`", true))
}

func TestLex_TemplateLiteral08(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{ReservedWordToken, "var"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "foo"}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{PunctuatorToken, "="}, InputElementDiv},
		TokenTest{Token{WhiteSpaceToken, " "}, InputElementDiv},
		TokenTest{Token{TemplateHeadToken, "`Hello ${"}, InputElementDiv},
		TokenTest{Token{IdentifierNameToken, "friend"}, InputElementDiv},
		TokenTest{Token{TemplateMiddleToken, "}! ${"}, InputElementTemplateTail},
		TokenTest{Token{TemplateHeadToken, "` ${"}, InputElementDiv},
		TokenTest{Token{NumericLiteralToken, "4"}, InputElementDiv},
		TokenTest{Token{TemplateTailToken, "}`"}, InputElementTemplateTail},
		TokenTest{Token{PunctuatorToken, "+"}, InputElementDiv},
		TokenTest{Token{TemplateHeadToken, "`${"}, InputElementDiv},
		TokenTest{Token{NumericLiteralToken, "2"}, InputElementDiv},
		TokenTest{Token{TemplateTailToken, "} `"}, InputElementTemplateTail},
		TokenTest{Token{TemplateTailToken, "}`"}, InputElementTemplateTail},
	}
	expectedTokensTable(t, expected, Lex("", "var foo = `Hello ${friend}! ${` ${4}`+`${2} `}`", true))
}

//
// Test WhiteSpaceToken
//

func TestLex_Whitespace1(t *testing.T) {
	js := " \t"
	expected := []Token{
		Token{WhiteSpaceToken, js},
	}
	expectedTokens(t, expected, Lex("", js, true))
}

func TestLex_Whitespace2(t *testing.T) {
	js := " "
	expected := []Token{
		Token{WhiteSpaceToken, js},
	}
	expectedTokens(t, expected, Lex("", js, true))
}
