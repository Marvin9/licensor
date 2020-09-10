package tests

import (
	"bytes"
	"testing"

	"github.com/Marvin9/licensor/steps"
)

func TestPrettyLicense(t *testing.T) {
	txt := []byte("Hi\nMy name is foo\nI am from bar.")
	model := steps.CommandModel{}

	op := model.Pretty(txt)
	expect := []byte("\n Hi\n My name is foo\n I am from bar.")

	if !bytes.Equal(op, expect) {
		t.Errorf("Pretty: %v\nExpected: %v\ngot: %v", string(txt), string(expect), string(op))
	}
}
