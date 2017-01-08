package gobel

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var errNotNull = fmt.Errorf("not null")

func Test_peek(t *testing.T) {
	l, _ := lex("", "101", false)

	n1 := l.next()
	p0 := l.peek()
	n0 := l.next()

	if p0 != n0 && p0 == '0' && n1 != '1' {
		t.Error("peek is broken")
	}
}

func TestLex_MultiLineComment1(t *testing.T) {
	expected := []Token{
		Token{MultiLineComment, "Hello World!"},
	}
	js := "/*Hello World!*/"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_MultiLineComment2(t *testing.T) {
	expected := []Token{
		Token{Error, "no multi line comment terminator \"*/\""},
	}
	js := "/*Hello World!"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_MultiLineComment3(t *testing.T) {
	expected := []Token{
		Token{Error, "no multi line comment terminator \"*/\""},
	}
	js := "/* \""
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_SingleLineComment1(t *testing.T) {
	expected := []Token{
		Token{SingleLineComment, " Hello World!"},
	}
	js := "// Hello World!"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}
func TestLex_SingleLineComment2(t *testing.T) {
	expected := []Token{
		Token{SingleLineComment, " Hello World!"},
	}
	js := "// Hello World!\n"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}
func TestLex_Comments(t *testing.T) {
	expected := []Token{
		Token{SingleLineComment, " Hello World!"},
		Token{MultiLineComment, "This is a multi\nline comment "},
	}
	js := "// Hello World!\n/*This is a multi\nline comment */"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_Whitespace1(t *testing.T) {
	js := " \t"
	expected := []Token{
		Token{WhiteSpace, js},
	}
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_Whitespace2(t *testing.T) {
	js := " "
	expected := []Token{
		Token{WhiteSpace, js},
	}

	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_Whitespace_AND_SingleLineComment(t *testing.T) {
	expected := []Token{
		Token{WhiteSpace, " \t"},
		Token{SingleLineComment, " Hello World!"},
	}
	js := " \t// Hello World!"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_LineTerminator1(t *testing.T) {
	js := "\n"
	expected := []Token{
		Token{LineTerminator, js},
	}

	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_LineTerminator2(t *testing.T) {
	js := "\n\n"
	expected := []Token{
		Token{LineTerminator, "\n"},
		Token{LineTerminator, "\n"},
	}

	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_Terminator_And_Whitespace(t *testing.T) {
	js := "\n\t\n"
	expected := []Token{
		Token{LineTerminator, "\n"},
		Token{WhiteSpace, "\t"},
		Token{LineTerminator, "\n"},
	}

	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_ReservedWord1(t *testing.T) {
	expected := []Token{}
	js := ""
	ws := " "

	lJs := lexer{}
	lJs.setStrict()
	for _, word := range lJs.reservedWords {
		js += word + ws
	}
	l, tokens := lex("", js, true)
	for _, word := range l.reservedWords {
		expected = append(expected, Token{ReservedWord, word})
		expected = append(expected, Token{WhiteSpace, ws})
	}

	expectedTokens(t, expected, tokens)
}

func TestLex_ReservedWord2(t *testing.T) {
	expected := []Token{}
	js := ""
	ws := " "

	lJs := lexer{}
	lJs.unsetStrict()
	for _, word := range lJs.reservedWords {
		js += word + ws
	}

	l, tokens := lex("", js, false)
	for _, word := range l.reservedWords {
		expected = append(expected, Token{ReservedWord, word})
		expected = append(expected, Token{WhiteSpace, ws})
	}

	expectedTokens(t, expected, tokens)
}

func TestLex_Punctuator1(t *testing.T) {
	expected := []Token{}
	js := ""
	ws := " "

	for _, punct := range punctuators {
		expected = append(expected, Token{Punctuator, punct})
		expected = append(expected, Token{WhiteSpace, ws})
		js += punct + ws
	}

	_, tokens := lex("", js, false)
	expectedTokens(t, expected, tokens)
}

func TestLex_DivPunctuator1(t *testing.T) {
	expected := []Token{
		Token{DivPunctuator, "/"},
		Token{WhiteSpace, " "},
		Token{DivPunctuator, "/="},
	}
	js := "/ /="
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_RightBracePunctuator1(t *testing.T) {
	expected := []Token{
		Token{Punctuator, "{"},
		Token{RightBracePunctuator, "}"},
	}
	js := "{}"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_StringLiteralSingleQuote1(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "var"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "foo"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "="},
		Token{WhiteSpace, " "},
		Token{StringLiteral, "'foo'"},
	}
	js := "var foo = 'foo'"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_StringLiteralSingleQuote2(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "var"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "foo"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "="},
		Token{WhiteSpace, " "},
		Token{Error, "did not reach end of string literal reached eof"},
		Token{StringLiteral, "'foo"},
	}
	js := "var foo = 'foo"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_StringLiteralDoubleQuote1(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "var"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "foo"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "="},
		Token{WhiteSpace, " "},
		Token{StringLiteral, "\"foo\""},
	}
	js := "var foo = \"foo\""
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_StringLiteralDoubleQuote2(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "var"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "foo"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "="},
		Token{WhiteSpace, " "},
		Token{Error, "did not reach end of string literal reached eof"},
		Token{StringLiteral, "\"foo"},
	}
	js := "var foo = \"foo"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

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

func TestLex_Identifier1(t *testing.T) {
	expected := []Token{
		Token{IdentifierName, "$"},
		Token{Punctuator, "="},
		Token{IdentifierName, "_"},
		Token{Punctuator, "="},
		Token{IdentifierName, "foo"},
	}
	js := "$=_=foo"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex_Identifier2(t *testing.T) {
	expected := []Token{
		Token{IdentifierName, "X"},
		Token{Punctuator, "&"},
		Token{IdentifierName, "ooooooooooooo___"},
		Token{LineTerminator, "\n"},
	}
	js := "X&ooooooooooooo___\n"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

// func TestLex_EscapeSequence0(t *testing.T) {
// 	expected := []Token{
// 		Token{IdentifierName, "X"},
// 		Token{Punctuator, "&"},
// 		Token{IdentifierName, "ooooooooooooo___"},
// 		Token{LineTerminator, "\n"},
// 	}
// 	js := `"\n"`
// 	_, tokens := lex("", js, true)
// 	expectedTokens(t, expected, tokens)
// }

func TestLex1(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "function"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "("},
		Token{Punctuator, ")"},
		Token{Punctuator, "{"},
		Token{RightBracePunctuator, "}"},
		Token{Punctuator, "("},
		Token{Punctuator, ")"},
	}

	js := "function (){}()"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex2(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "function"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "foo"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "("},
		Token{Punctuator, ")"},
		Token{Punctuator, "{"},
		Token{RightBracePunctuator, "}"},
		Token{Punctuator, "("},
		Token{Punctuator, ")"},
	}

	js := "function foo (){}()"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLex3(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "function"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "add"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "("},
		Token{IdentifierName, "a"},
		Token{Punctuator, ","},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "b"},
		Token{Punctuator, ")"},
		Token{Punctuator, "{"},
		Token{LineTerminator, "\n"},
		Token{WhiteSpace, "\t"},
		Token{ReservedWord, "return"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "a"},
		Token{Punctuator, "+"},
		Token{IdentifierName, "b"},
		Token{LineTerminator, "\n"},
		Token{RightBracePunctuator, "}"},
	}

	js := "function add (a, b){\n\treturn a+b\n}"
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLexJS(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "function"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "fibonacci"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{Punctuator, ")"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "{"},
		Token{LineTerminator, "\n"},
		Token{WhiteSpace, "  "},
		Token{ReservedWord, "if"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{WhiteSpace, " "},
		Token{Punctuator, ">="},
		Token{WhiteSpace, " "},
		Token{NumericLiteral, "2"},
		Token{Punctuator, ")"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "{"},
		Token{LineTerminator, "\n"},
		Token{WhiteSpace, "    "},
		Token{ReservedWord, "return"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "fibonacci"},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{NumericLiteral, "-1"},
		Token{Punctuator, ")"},
		Token{WhiteSpace, " "},
		Token{Punctuator, "+"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "fibonacci"},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{NumericLiteral, "-2"},
		Token{Punctuator, ")"},
		Token{LineTerminator, "\n"},
		Token{WhiteSpace, "  "},
		Token{RightBracePunctuator, "}"},
		Token{LineTerminator, "\n"},
		Token{WhiteSpace, "  "},
		Token{ReservedWord, "return"},
		Token{WhiteSpace, " "},
		Token{NumericLiteral, "1"},
		Token{LineTerminator, "\n"},
		Token{RightBracePunctuator, "}"},
		Token{LineTerminator, "\n"},
		Token{LineTerminator, "\n"},
		Token{IdentifierName, "console"},
		Token{Punctuator, "."},
		Token{IdentifierName, "log"},
		Token{Punctuator, "("},
		Token{IdentifierName, "fibonacci"},
		Token{Punctuator, "("},
		Token{NumericLiteral, "7"},
		Token{Punctuator, ")"},
		Token{Punctuator, ")"},
		Token{LineTerminator, "\n"},
	}

	testData, err := ioutil.ReadFile("testdata/index01.js")
	if err != nil {
		t.Fatal(err)
	}

	js := string(testData)
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestLexJS2(t *testing.T) {
	expected := []Token{
		Token{ReservedWord, "function"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "fibonacci"},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{Punctuator, ")"},
		Token{Punctuator, "{"},
		Token{ReservedWord, "if"},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{Punctuator, ">="},
		Token{NumericLiteral, "2"},
		Token{Punctuator, ")"},
		Token{Punctuator, "{"},
		Token{ReservedWord, "return"},
		Token{WhiteSpace, " "},
		Token{IdentifierName, "fibonacci"},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{NumericLiteral, "-1"},
		Token{Punctuator, ")"},
		Token{Punctuator, "+"},
		Token{IdentifierName, "fibonacci"},
		Token{Punctuator, "("},
		Token{IdentifierName, "n"},
		Token{NumericLiteral, "-2"},
		Token{Punctuator, ")"},
		Token{RightBracePunctuator, "}"},
		Token{ReservedWord, "return"},
		Token{WhiteSpace, " "},
		Token{NumericLiteral, "1"},
		Token{RightBracePunctuator, "}"},
		Token{Punctuator, ";"},
		Token{IdentifierName, "console"},
		Token{Punctuator, "."},
		Token{IdentifierName, "log"},
		Token{Punctuator, "("},
		Token{IdentifierName, "fibonacci"},
		Token{Punctuator, "("},
		Token{NumericLiteral, "7"},
		Token{Punctuator, ")"},
		Token{Punctuator, ")"},
		Token{LineTerminator, "\n"},
	}

	testData, err := ioutil.ReadFile("testdata/index02.js")
	if err != nil {
		t.Fatal(err)
	}

	js := string(testData)
	_, tokens := lex("", js, true)
	expectedTokens(t, expected, tokens)
}

func TestToken_String(t *testing.T) {
	t1 := Token{tokenType(-1), ""}
	if t1.String() == "" {
		t.Error("t1.String() string returns empty string")
	}
	t2 := Token{LineTerminator, ""}
	if t2.String() == "" {
		t.Error("t2.String() string returns empty string")
	}
	t3 := Token{MultiLineComment, " some string value "}
	if t3.String() == "" {
		t.Error("t3.String() string returns empty string")
	}
}

func expectedTokens(t *testing.T, expectedTokens []Token, tokens chan Token) {
	i := 0
	for tok := range tokens {
		if i >= len(expectedTokens) {
			t.Errorf("expected fewer tokens (expected: %d, got %d %s)", len(expectedTokens), i, tok)
		} else {
			// t.Logf("%d: %s %s\n", i, expectedTokens[i], tok)
			if !expectedTokens[i].equals(tok) {
				t.Errorf("expected and recived tokens do not match [%d](%s != %s)", i, tok, expectedTokens[i])
			}
			i++
		}
	}
	if i < len(expectedTokens) {
		t.Errorf("expected more tokens (expected: %d, got %d)", len(expectedTokens), i)
	}
}
