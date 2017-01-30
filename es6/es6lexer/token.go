package es6lexer

import "fmt"

// Token is a unit generated by the lexer whitch includes a type
// or value
type Token struct {
	Type  Type
	Value string
}

func (tok Token) String() string {
	val := ""
	if len(tok.Value) > 0 {
		val = " \"" + tok.Value + "\""
	}
	return fmt.Sprintf("<%s%s>", tok.Type.String(), val)
}

// Type represents a golang type
type Type int

// not handled Type's
// TODO missing LineTerminatorSequence
const (
	Error Type = iota
	EOF
	// Comment ::
	MultiLineComment
	SingleLineComment
	// WhiteSpace ::
	WhiteSpace
	// LineTerminator ::
	LineTerminator
	// ComomonToken ::
	IdentifierName
	ReservedWord
	//   Punctuator
	Punctuator
	RightBracePunctuator
	DivPunctuator

	NumericLiteral
	StringLiteral

	// Template ::
	NoSubstitutionTemplate
	TemplateHead
)

func (typ Type) String() string {
	switch typ {
	case Error:
		return "Error"
	case EOF:
		return "EOF"
	case MultiLineComment:
		return "MultiLineComment"
	case SingleLineComment:
		return "SingleLineComment"
	case WhiteSpace:
		return "WhiteSpace"
	case LineTerminator:
		return "LineTerminator"
	case IdentifierName:
		return "IdentifierName"
	case ReservedWord:
		return "ReservedWord"
	case Punctuator:
		return "Punctuator"
	case RightBracePunctuator:
		return "RightBracePunctuator"
	case DivPunctuator:
		return "DivPunctuator"
	case NumericLiteral:
		return "NumericLiteral"
	case StringLiteral:
		return "StringLiteral"
	case TemplateHead:
		return "TemplateHead"
	default:
		return "UNKNOWN_TOKEN_TYPE"
	}
}

// Equals checks if the token is equal to another token
func (tok Token) Equals(tok2 Token) bool {
	return tok.Type == tok2.Type && tok.Value == tok2.Value
}
