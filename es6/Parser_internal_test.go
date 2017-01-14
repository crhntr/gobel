package es6

import (
	"os"
	"testing"
)

func TestParseES601(t *testing.T) {
	r, err := os.Open("testdata/TestParseES601.js")
	if err != nil {
		t.Fatal(err)
	}
	_, err = ParseES6(r)
	if err != nil {
		t.Error(err)
	}
}
