package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestFormatEnvVarStatements(t *testing.T) {
	inputData := map[string]string{
		"foo": "foobar",
		"bar": "barfoo",
	}
	result := convertToEnvVarStatements(inputData)

	if !slices.Contains(result, "export FOO=foobar") {
		t.Errorf("result did not contain export FOO=foobar")
	}

	if !slices.Contains(result, "export BAR=barfoo") {
		t.Errorf("result did not contain export BAR=barfoo")
	}
}
