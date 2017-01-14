package es6

import (
	"io"
	"io/ioutil"
	"log"
)

// Safe is a bool
var Safe bool

// ParseES6 is a func
func ParseES6(r io.Reader) (*ASTNode, error) {
	root := ASTNode{}
	all, err := ioutil.ReadAll(r)
	if err != nil {
		return &root, err
	}

	_, tokens := lex("", string(all), Safe)

	for token := range tokens {
		log.Print(token)
	}
	return &root, nil
}
