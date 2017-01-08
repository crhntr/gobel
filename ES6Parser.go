package gobel

import (
	"io"
	"io/ioutil"
	"log"
)

var Safe bool

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
