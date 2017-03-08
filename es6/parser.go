package es6

import (
	"io"
	"io/ioutil"
)

// DecodeES6Script ...
func DecodeES6Script(r io.Reader) (ASTNode, error) {
	b, _ := ioutil.ReadAll(r)
	l := Lex("", string(b), true)
	l.SkipWhitespace = true
	return ParseScriptNode(l)
}
