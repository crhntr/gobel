package es6

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var errNotNull = fmt.Errorf("not null")

func Test_peek(t *testing.T) {
	l := Lex("", "101", false)
	l.CaptureWhitespaceTokens = true

	n1 := l.next()
	p0 := l.peek()
	n0 := l.next()

	if p0 != n0 && p0 == '0' && n1 != '1' {
		t.Error("peek is broken")
	}
}

func TestLex_Whitespace_AND_SingleLineComment(t *testing.T) {
	expected := []Token{
		Token{Type: WhiteSpaceToken, Value: " \t"},
		Token{Type: SingleLineCommentToken, Value: " Hello World!"},
	}
	js := " \t// Hello World!"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_Terminator_And_Whitespace(t *testing.T) {
	js := "\n\t\n"
	expected := []Token{
		Token{Type: LineTerminatorToken, Value: "\n"},
		Token{Type: WhiteSpaceToken, Value: "\t"},
		Token{Type: LineTerminatorToken, Value: "\n"},
	}

	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
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
	l.CaptureWhitespaceTokens = true
	for _, word := range l.reservedWords {
		expected = append(expected, Token{Type: ReservedWordToken, Value: word})
		expected = append(expected, Token{Type: WhiteSpaceToken, Value: ws})
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
	l.CaptureWhitespaceTokens = true
	for _, word := range l.reservedWords {
		expected = append(expected, Token{Type: ReservedWordToken, Value: word})
		expected = append(expected, Token{Type: WhiteSpaceToken, Value: ws})
	}

	expectedTokens(t, expected, l)
}

// func TestLex_EscapeSequence0(t *testing.T) {
// 	expected := []Token{
// 		Token{Type: IdentifierName, Value: "X"},
// 		Token{Type: Punctuator, Value: "&"},
// 		Token{Type: IdentifierName, Value: "ooooooooooooo___"},
// 		Token{Type: LineTerminator, Value: "\n"},
// 	}
// 	js := `"\n"`
// 	l := lex("", js, true)
// 	expectedTokens(t, expected, l)
// }

func TestLex1(t *testing.T) {
	expected := []Token{
		Token{Type: ReservedWordToken, Value: "function"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: PunctuatorToken, Value: ")"},
		Token{Type: PunctuatorToken, Value: "{"},
		Token{Type: RightBracePunctuatorToken, Value: "}"},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: PunctuatorToken, Value: ")"},
	}

	js := "function (){}()"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex2(t *testing.T) {
	expected := []Token{
		Token{Type: ReservedWordToken, Value: "function"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "foo"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: PunctuatorToken, Value: ")"},
		Token{Type: PunctuatorToken, Value: "{"},
		Token{Type: RightBracePunctuatorToken, Value: "}"},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: PunctuatorToken, Value: ")"},
	}

	js := "function foo (){}()"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex3(t *testing.T) {
	expected := []Token{
		Token{Type: ReservedWordToken, Value: "function"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "add"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: IdentifierNameToken, Value: "a"},
		Token{Type: PunctuatorToken, Value: ","},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "b"},
		Token{Type: PunctuatorToken, Value: ")"},
		Token{Type: PunctuatorToken, Value: "{"},
		Token{Type: LineTerminatorToken, Value: "\n"},
		Token{Type: WhiteSpaceToken, Value: "\t"},
		Token{Type: ReservedWordToken, Value: "return"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "a"},
		Token{Type: PunctuatorToken, Value: "+"},
		Token{Type: IdentifierNameToken, Value: "b"},
		Token{Type: LineTerminatorToken, Value: "\n"},
		Token{Type: RightBracePunctuatorToken, Value: "}"},
	}

	js := "function add (a, b){\n\treturn a+b\n}"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLexJS(t *testing.T) {
	expected := []Token{
		Token{Type: ReservedWordToken, Value: "function"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "fibonacci"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: IdentifierNameToken, Value: "n"},
		Token{Type: PunctuatorToken, Value: ")"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "{"},
		Token{Type: LineTerminatorToken, Value: "\n"},
		Token{Type: WhiteSpaceToken, Value: "  "},
		Token{Type: ReservedWordToken, Value: "if"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: IdentifierNameToken, Value: "n"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: ">="},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: NumericLiteralToken, Value: "2"},
		Token{Type: PunctuatorToken, Value: ")"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "{"},
		Token{Type: LineTerminatorToken, Value: "\n"},
		Token{Type: WhiteSpaceToken, Value: "    "},
		Token{Type: ReservedWordToken, Value: "return"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "fibonacci"},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: IdentifierNameToken, Value: "n"},
		Token{Type: NumericLiteralToken, Value: "-1"},
		Token{Type: PunctuatorToken, Value: ")"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "+"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "fibonacci"},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: IdentifierNameToken, Value: "n"},
		Token{Type: NumericLiteralToken, Value: "-2"},
		Token{Type: PunctuatorToken, Value: ")"},
		Token{Type: LineTerminatorToken, Value: "\n"},
		Token{Type: WhiteSpaceToken, Value: "  "},
		Token{Type: RightBracePunctuatorToken, Value: "}"},
		Token{Type: LineTerminatorToken, Value: "\n"},
		Token{Type: WhiteSpaceToken, Value: "  "},
		Token{Type: ReservedWordToken, Value: "return"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: NumericLiteralToken, Value: "1"},
		Token{Type: LineTerminatorToken, Value: "\n"},
		Token{Type: RightBracePunctuatorToken, Value: "}"},
		Token{Type: LineTerminatorToken, Value: "\n"},
		Token{Type: LineTerminatorToken, Value: "\n"},
		Token{Type: IdentifierNameToken, Value: "console"},
		Token{Type: PunctuatorToken, Value: "."},
		Token{Type: IdentifierNameToken, Value: "log"},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: IdentifierNameToken, Value: "fibonacci"},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: NumericLiteralToken, Value: "7"},
		Token{Type: PunctuatorToken, Value: ")"},
		Token{Type: PunctuatorToken, Value: ")"},
		Token{Type: LineTerminatorToken, Value: "\n"},
	}

	testData, err := ioutil.ReadFile("testdata/index01.js")
	if err != nil {
		t.Fatal(err)
	}

	js := string(testData)
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLexJS2(t *testing.T) {
	expected := []Token{
		Token{Type: ReservedWordToken, Value: "function"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "fibonacci"},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: IdentifierNameToken, Value: "n"},
		Token{Type: PunctuatorToken, Value: ")"},
		Token{Type: PunctuatorToken, Value: "{"},
		Token{Type: ReservedWordToken, Value: "if"},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: IdentifierNameToken, Value: "n"},
		Token{Type: PunctuatorToken, Value: ">="},
		Token{Type: NumericLiteralToken, Value: "2"},
		Token{Type: PunctuatorToken, Value: ")"},
		Token{Type: PunctuatorToken, Value: "{"},
		Token{Type: ReservedWordToken, Value: "return"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "fibonacci"},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: IdentifierNameToken, Value: "n"},
		Token{Type: NumericLiteralToken, Value: "-1"},
		Token{Type: PunctuatorToken, Value: ")"},
		Token{Type: PunctuatorToken, Value: "+"},
		Token{Type: IdentifierNameToken, Value: "fibonacci"},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: IdentifierNameToken, Value: "n"},
		Token{Type: NumericLiteralToken, Value: "-2"},
		Token{Type: PunctuatorToken, Value: ")"},
		Token{Type: RightBracePunctuatorToken, Value: "}"},
		Token{Type: ReservedWordToken, Value: "return"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: NumericLiteralToken, Value: "1"},
		Token{Type: RightBracePunctuatorToken, Value: "}"},
		Token{Type: PunctuatorToken, Value: ";"},
		Token{Type: IdentifierNameToken, Value: "console"},
		Token{Type: PunctuatorToken, Value: "."},
		Token{Type: IdentifierNameToken, Value: "log"},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: IdentifierNameToken, Value: "fibonacci"},
		Token{Type: PunctuatorToken, Value: "("},
		Token{Type: NumericLiteralToken, Value: "7"},
		Token{Type: PunctuatorToken, Value: ")"},
		Token{Type: PunctuatorToken, Value: ")"},
		Token{Type: LineTerminatorToken, Value: "\n"},
	}

	testData, err := ioutil.ReadFile("testdata/index02.js")
	if err != nil {
		t.Fatal(err)
	}

	js := string(testData)
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLexJS3(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{Type: ReservedWordToken, Value: "function"}, InputElementRegExp},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementRegExp},
		TokenTest{Token{Type: IdentifierNameToken, Value: "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{Type: PunctuatorToken, Value: "("}, InputElementRegExp},
		TokenTest{Token{Type: IdentifierNameToken, Value: "sequence"}, InputElementRegExp},
		TokenTest{Token{Type: PunctuatorToken, Value: ")"}, InputElementRegExp},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementRegExp},
		TokenTest{Token{Type: PunctuatorToken, Value: "{"}, InputElementRegExp},
		TokenTest{Token{Type: LineTerminatorToken, Value: "\n"}, InputElementRegExp},
		TokenTest{Token{Type: WhiteSpaceToken, Value: "  "}, InputElementRegExp},
		TokenTest{Token{Type: ReservedWordToken, Value: "return"}, InputElementRegExp},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementRegExp},
		TokenTest{Token{Type: IdentifierNameToken, Value: "valid"}, InputElementRegExp},
		TokenTest{Token{Type: PunctuatorToken, Value: "."}, InputElementRegExp},
		TokenTest{Token{Type: IdentifierNameToken, Value: "match"}, InputElementRegExp},
		TokenTest{Token{Type: PunctuatorToken, Value: "("}, InputElementRegExp},
		TokenTest{Token{Type: RegExToken, Value: "/([CGAT]{3}){1,}/g"}, InputElementRegExp},
		TokenTest{Token{Type: PunctuatorToken, Value: ")"}, InputElementRegExp},
		TokenTest{Token{Type: LineTerminatorToken, Value: "\n"}, InputElementRegExp},
		TokenTest{Token{Type: RightBracePunctuatorToken, Value: "}"}, InputElementRegExp},
		TokenTest{Token{Type: LineTerminatorToken, Value: "\n"}, InputElementRegExp},
		TokenTest{Token{Type: IdentifierNameToken, Value: "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{Type: PunctuatorToken, Value: "("}, InputElementRegExp},
		TokenTest{Token{Type: StringLiteralToken, Value: "\"ATATTGGTGTTCATGTGCGCGGGGCCGACGAGCTACTGGCAGAACCACGAGGACAAGAGGTGA\""}, InputElementRegExp},
		TokenTest{Token{Type: PunctuatorToken, Value: ")"}, InputElementRegExp},
		TokenTest{Token{Type: LineTerminatorToken, Value: "\n"}, InputElementRegExp},
		TokenTest{Token{Type: IdentifierNameToken, Value: "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{Type: PunctuatorToken, Value: "("}, InputElementRegExp},
		TokenTest{Token{Type: StringLiteralToken, Value: "\"FAIL\""}, InputElementRegExp},
		TokenTest{Token{Type: PunctuatorToken, Value: ")"}, InputElementRegExp},
		TokenTest{Token{Type: LineTerminatorToken, Value: "\n"}, InputElementRegExp},
		TokenTest{Token{Type: IdentifierNameToken, Value: "anylize_dna_sequence"}, InputElementRegExp},
		TokenTest{Token{Type: PunctuatorToken, Value: "("}, InputElementRegExp},
		TokenTest{Token{Type: StringLiteralToken, Value: "\"Alanine\""}, InputElementRegExp},
		TokenTest{Token{Type: PunctuatorToken, Value: ")"}, InputElementRegExp},
		TokenTest{Token{Type: LineTerminatorToken, Value: "\n"}, InputElementRegExp},
		TokenTest{Token{Type: EOFToken, Value: ""}, InputElementDiv},
	}
	testData, err := ioutil.ReadFile("testdata/TestLexJS3.js")
	if err != nil {
		t.Fatal(err)
	}

	js := string(testData)
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokensTable(t, expected, l)
}

func TestToken_String(t *testing.T) {
	t1 := Token{Type: TokenType(-1), Value: ""}
	if t1.String() == "" {
		t.Error("t1.String() string returns empty string")
	}
	t2 := Token{Type: LineTerminatorToken, Value: ""}
	if t2.String() == "" {
		t.Error("t2.String() string returns empty string")
	}
	t3 := Token{Type: MultiLineCommentToken, Value: " some string value "}
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
		Token{Type: MultiLineCommentToken, Value: "Hello World!"},
	}
	js := "/*Hello World!*/"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_MultiLineComment2(t *testing.T) {
	expected := []Token{
		Token{Type: ErrorToken, Value: "no multi line comment terminator \"*/\""},
	}
	js := "/*Hello World!"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_MultiLineComment3(t *testing.T) {
	expected := []Token{
		Token{Type: ErrorToken, Value: "no multi line comment terminator \"*/\""},
	}
	js := "/* \""
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_SingleLineComment1(t *testing.T) {
	expected := []Token{
		Token{Type: SingleLineCommentToken, Value: " Hello World!"},
	}
	js := "// Hello World!"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}
func TestLex_SingleLineComment2(t *testing.T) {
	expected := []Token{
		Token{Type: SingleLineCommentToken, Value: " Hello World!"},
	}
	js := "// Hello World!\n"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}
func TestLex_Comments(t *testing.T) {
	expected := []Token{
		Token{Type: SingleLineCommentToken, Value: " Hello World!"},
		Token{Type: MultiLineCommentToken, Value: "This is a multi\nline comment "},
	}
	js := "// Hello World!\n/*This is a multi\nline comment */"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

// Test EscapeSequence
func TestLex_lexEscapeSequence01(t *testing.T) {
	expected := []Token{
		Token{Type: StringLiteralToken, Value: "\"\\u0074\\x61z\nzz\""},
	}
	js := "\"\\u0074\\x61z\nzz\""
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

//
// Test Identifier
//
func TestLex_Identifier1(t *testing.T) {
	expected := []Token{
		Token{Type: IdentifierNameToken, Value: "$"},
		Token{Type: PunctuatorToken, Value: "="},
		Token{Type: IdentifierNameToken, Value: "_"},
		Token{Type: PunctuatorToken, Value: "="},
		Token{Type: IdentifierNameToken, Value: "foo"},
	}
	js := "$=_=foo"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_Identifier2(t *testing.T) {
	expected := []Token{
		Token{Type: IdentifierNameToken, Value: "X"},
		Token{Type: PunctuatorToken, Value: "&"},
		Token{Type: IdentifierNameToken, Value: "ooooooooooooo___"},
		Token{Type: LineTerminatorToken, Value: "\n"},
	}
	js := "X&ooooooooooooo___\n"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

//
// Test LineTerminator
//

func TestLex_LineTerminator1(t *testing.T) {
	js := "\n"
	expected := []Token{
		Token{Type: LineTerminatorToken, Value: js},
	}
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_LineTerminator2(t *testing.T) {
	js := "\n\n"
	expected := []Token{
		Token{Type: LineTerminatorToken, Value: "\n"},
		Token{Type: LineTerminatorToken, Value: "\n"},
	}
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_LineTerminator3(t *testing.T) {
	js := "\u000A\u000D\u2028\u2029"
	expected := []Token{
		Token{Type: LineTerminatorToken, Value: "\u000A"},
		Token{Type: LineTerminatorToken, Value: "\u000D"},
		Token{Type: LineTerminatorToken, Value: "\u2028"},
		Token{Type: LineTerminatorToken, Value: "\u2029"},
	}
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

//
// Test Number Literal
//

func TestLex_NumericLiteral0(t *testing.T) {
	expected := []Token{
		Token{Type: NumericLiteralToken, Value: "0"},
	}
	js := "0"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_NumericLiteral1(t *testing.T) {
	expected := []Token{
		Token{Type: NumericLiteralToken, Value: "1"},
	}
	js := "1"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_NumericLiteral2(t *testing.T) {
	expected := []Token{
		Token{Type: NumericLiteralToken, Value: "10"},
	}
	js := "10"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_NumericLiteral3(t *testing.T) {
	expected := []Token{
		Token{Type: NumericLiteralToken, Value: "0xAB10"},
	}
	js := "0xAB10"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_NumericLiteral4(t *testing.T) {
	expected := []Token{
		Token{Type: NumericLiteralToken, Value: "0b0100"},
	}
	js := "0b0100"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_NumericLiteral5(t *testing.T) {
	expected := []Token{
		Token{Type: NumericLiteralToken, Value: "0O0005"},
	}
	js := "0O0005"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_NumericLiteral6(t *testing.T) {
	expected := []Token{
		Token{Type: NumericLiteralToken, Value: "-6"},
	}
	js := "-6"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_NumericLiteral7(t *testing.T) {
	expected := []Token{
		Token{Type: NumericLiteralToken, Value: "0.0007"},
	}
	js := "0.0007"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_NumericLiteral8(t *testing.T) {
	expected := []Token{
		Token{Type: NumericLiteralToken, Value: "8.08"},
	}
	js := "8.08"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_NumericLiteral9(t *testing.T) {
	expected := []Token{
		Token{Type: NumericLiteralToken, Value: "3e2"},
	}
	js := "3e2"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_NumericLiteral10(t *testing.T) {
	expected := []Token{
		Token{Type: ErrorToken, Value: "bad number syntax: \"1o\""},
	}
	js := "1o"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

//
// Test Punctuator
//

func TestLex_Punctuator1(t *testing.T) {
	expected := []Token{}
	js := ""
	ws := " "

	for _, punct := range punctuators {
		expected = append(expected, Token{Type: PunctuatorToken, Value: punct})
		expected = append(expected, Token{Type: WhiteSpaceToken, Value: ws})
		js += punct + ws
	}

	l := Lex("", js, false)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_DivPunctuator1(t *testing.T) {
	expected := []Token{
		Token{Type: IdentifierNameToken, Value: "i"},
		Token{Type: DivPunctuatorToken, Value: "/="},
		Token{Type: IdentifierNameToken, Value: "j"},
		Token{Type: DivPunctuatorToken, Value: "/"},
		Token{Type: NumericLiteralToken, Value: "2"},
	}
	js := "i/=j/2"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_RightBracePunctuator1(t *testing.T) {
	expected := []Token{
		Token{Type: PunctuatorToken, Value: "{"},
		Token{Type: RightBracePunctuatorToken, Value: "}"},
	}
	js := "{}"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

//
// Test RegEx
//

func TestLex_RegEx00(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{Type: ReservedWordToken, Value: "var"}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: IdentifierNameToken, Value: "foo"}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: PunctuatorToken, Value: "="}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: RegExToken, Value: "/abc/i"}, InputElementRegExp},
	}
	l := Lex("", "var foo = /abc/i", true)
	l.CaptureWhitespaceTokens = true
	expectedTokensTable(t, expected, l)
}

func TestLex_RegEx01(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{Type: ReservedWordToken, Value: "var"}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: IdentifierNameToken, Value: "foo"}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: PunctuatorToken, Value: "="}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: ErrorToken, Value: "regex did not close with '/' "}, InputElementRegExp},
	}
	l := Lex("", "var foo = /abc", true)
	l.CaptureWhitespaceTokens = true
	expectedTokensTable(t, expected, l)
}

//
// Test String
//

func TestLex_StringLiteralSingleQuote1(t *testing.T) {
	expected := []Token{
		Token{Type: ReservedWordToken, Value: "var"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "foo"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "="},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: StringLiteralToken, Value: "'foo'"},
	}
	js := "var foo = 'foo'"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_StringLiteralSingleQuote2(t *testing.T) {
	expected := []Token{
		Token{Type: ReservedWordToken, Value: "var"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "foo"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "="},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: ErrorToken, Value: "did not reach end of string literal reached eof"},
		Token{Type: StringLiteralToken, Value: "'foo"},
	}
	js := "var foo = 'foo"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_StringLiteralDoubleQuote1(t *testing.T) {
	expected := []Token{
		Token{Type: ReservedWordToken, Value: "var"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "foo"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "="},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: StringLiteralToken, Value: "\"foo\""},
	}
	js := "var foo = \"foo\""
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_StringLiteralDoubleQuote2(t *testing.T) {
	expected := []Token{
		Token{Type: ReservedWordToken, Value: "var"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "foo"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "="},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: ErrorToken, Value: "did not reach end of string literal reached eof"},
		Token{Type: StringLiteralToken, Value: "\"foo"},
	}
	js := "var foo = \"foo"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

//
// Test Template literal
//

func TestLex_TemplateLiteral01(t *testing.T) {
	expected := []Token{
		Token{Type: ReservedWordToken, Value: "var"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "foo"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "="},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: NoSubstitutionTemplateToken, Value: "`foo`"},
	}
	js := "var foo = `foo`"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_TemplateLiteral02(t *testing.T) {
	expected := []Token{
		Token{Type: ReservedWordToken, Value: "var"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "foo"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "="},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: NoSubstitutionTemplateToken, Value: "``"},
	}
	js := "var foo = ``"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_TemplateLiteral03(t *testing.T) {
	expected := []Token{
		Token{Type: ReservedWordToken, Value: "var"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "foo"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "="},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: ErrorToken, Value: "did not reach end of template literal reached eof"},
	}
	js := "var foo = `"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_TemplateLiteral04(t *testing.T) {
	expected := []Token{
		Token{Type: ReservedWordToken, Value: "var"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: IdentifierNameToken, Value: "foo"},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: PunctuatorToken, Value: "="},
		Token{Type: WhiteSpaceToken, Value: " "},
		Token{Type: TemplateHeadToken, Value: "`Hello ${"},
		Token{Type: IdentifierNameToken, Value: "friend"},
	}
	js := "var foo = `Hello ${friend"
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_TemplateLiteral05(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{Type: ReservedWordToken, Value: "var"}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: IdentifierNameToken, Value: "foo"}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: PunctuatorToken, Value: "="}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: TemplateHeadToken, Value: "`Hello ${"}, InputElementDiv},
		TokenTest{Token{Type: IdentifierNameToken, Value: "friend"}, InputElementDiv},
		TokenTest{Token{Type: TemplateTailToken, Value: "}!`"}, InputElementRegExpOrTemplateTail},
	}
	l := Lex("", "var foo = `Hello ${friend}!`", true)
	l.CaptureWhitespaceTokens = true
	expectedTokensTable(t, expected, l)
}

func TestLex_TemplateLiteral06(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{Type: ReservedWordToken, Value: "var"}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: IdentifierNameToken, Value: "foo"}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: PunctuatorToken, Value: "="}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: TemplateHeadToken, Value: "`Hello ${"}, InputElementDiv},
		TokenTest{Token{Type: IdentifierNameToken, Value: "friend"}, InputElementDiv},
		TokenTest{Token{Type: ErrorToken, Value: "did not reach TemplateMiddle or TemplateTail but reached eof"}, InputElementRegExpOrTemplateTail},
	}
	l := Lex("", "var foo = `Hello ${friend}! ", true)
	l.CaptureWhitespaceTokens = true
	expectedTokensTable(t, expected, l)
}

func TestLex_TemplateLiteral07(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{Type: ReservedWordToken, Value: "var"}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: IdentifierNameToken, Value: "foo"}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: PunctuatorToken, Value: "="}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: TemplateHeadToken, Value: "`Hello ${"}, InputElementDiv},
		TokenTest{Token{Type: IdentifierNameToken, Value: "friend"}, InputElementDiv},
		TokenTest{Token{Type: TemplateMiddleToken, Value: "}! ${"}, InputElementRegExpOrTemplateTail},
		TokenTest{Token{Type: TemplateHeadToken, Value: "` ${"}, InputElementDiv},
		TokenTest{Token{Type: NumericLiteralToken, Value: "4"}, InputElementDiv},
		TokenTest{Token{Type: TemplateTailToken, Value: "}`"}, InputElementRegExpOrTemplateTail},
		TokenTest{Token{Type: PunctuatorToken, Value: "+"}, InputElementDiv},
		TokenTest{Token{Type: TemplateHeadToken, Value: "`${"}, InputElementDiv},
		TokenTest{Token{Type: NumericLiteralToken, Value: "2"}, InputElementDiv},
		TokenTest{Token{Type: TemplateTailToken, Value: "} `"}, InputElementRegExpOrTemplateTail},
		TokenTest{Token{Type: TemplateTailToken, Value: "}`"}, InputElementRegExpOrTemplateTail},
	}
	l := Lex("", "var foo = `Hello ${friend}! ${` ${4}`+`${2} `}`", true)
	l.CaptureWhitespaceTokens = true
	expectedTokensTable(t, expected, l)
}

func TestLex_TemplateLiteral08(t *testing.T) {
	expected := []TokenTest{
		TokenTest{Token{Type: ReservedWordToken, Value: "var"}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: IdentifierNameToken, Value: "foo"}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: PunctuatorToken, Value: "="}, InputElementDiv},
		TokenTest{Token{Type: WhiteSpaceToken, Value: " "}, InputElementDiv},
		TokenTest{Token{Type: TemplateHeadToken, Value: "`Hello ${"}, InputElementDiv},
		TokenTest{Token{Type: IdentifierNameToken, Value: "friend"}, InputElementDiv},
		TokenTest{Token{Type: TemplateMiddleToken, Value: "}! ${"}, InputElementTemplateTail},
		TokenTest{Token{Type: TemplateHeadToken, Value: "` ${"}, InputElementDiv},
		TokenTest{Token{Type: NumericLiteralToken, Value: "4"}, InputElementDiv},
		TokenTest{Token{Type: TemplateTailToken, Value: "}`"}, InputElementTemplateTail},
		TokenTest{Token{Type: PunctuatorToken, Value: "+"}, InputElementDiv},
		TokenTest{Token{Type: TemplateHeadToken, Value: "`${"}, InputElementDiv},
		TokenTest{Token{Type: NumericLiteralToken, Value: "2"}, InputElementDiv},
		TokenTest{Token{Type: TemplateTailToken, Value: "} `"}, InputElementTemplateTail},
		TokenTest{Token{Type: TemplateTailToken, Value: "}`"}, InputElementTemplateTail},
	}
	l := Lex("", "var foo = `Hello ${friend}! ${` ${4}`+`${2} `}`", true)
	l.CaptureWhitespaceTokens = true
	expectedTokensTable(t, expected, l)
}

//
// Test WhiteSpaceToken
//

func TestLex_Whitespace1(t *testing.T) {
	js := " \t"
	expected := []Token{
		Token{Type: WhiteSpaceToken, Value: js},
	}
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}

func TestLex_Whitespace2(t *testing.T) {
	js := " "
	expected := []Token{
		Token{Type: WhiteSpaceToken, Value: js},
	}
	l := Lex("", js, true)
	l.CaptureWhitespaceTokens = true
	expectedTokens(t, expected, l)
}
