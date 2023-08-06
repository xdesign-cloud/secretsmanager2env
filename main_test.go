package main

import (
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func TestFormatEnvVarStatements(t *testing.T) {
	inputData := map[string]string{
		"foo": "foobar",
		"bar": "barfoo",
	}
	result := convertToEnvVarStatements(inputData)

	want := "export FOO=\"foobar\""
	if !slices.Contains(result, want) {
		t.Errorf("wanted to find '%s', got: '%s'", want, result)
	}

	want = "export BAR=\"barfoo\""
	if !slices.Contains(result, want) {
		t.Errorf("wanted to find '%s', got: '%s'", want, result)
	}
}

func TestValuesWithSpacesAreQuoted(t *testing.T) {
	inputData := map[string]string{
		"foo": "foo bar",
	}

	want := "export FOO=\"foo bar\""
	got := convertToEnvVarStatements(inputData)[0]

	if want != got {
		t.Errorf("wanted '%s', got '%s'", want, got)
	}
}

func TestKeysWithSpacesAreUnderscored(t *testing.T) {
	inputData := map[string]string{
		"foo bar": "foo",
	}
	want := "export FOO_BAR="
	got := convertToEnvVarStatements(inputData)[0]

	if !strings.HasPrefix(got, want) {
		t.Errorf("wanted prefix '%s', got '%s'", want, got)
	}
}
