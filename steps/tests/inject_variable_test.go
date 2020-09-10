package tests

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/Marvin9/licensor/steps"
)

func TestInjectVariable(t *testing.T) {
	model := steps.CommandModel{
		Template: map[string]string{
			"foo": "bar",
		},
	}

	template := []byte("I am {{foo}}")
	op := model.InjectVariable(template)
	expect := []byte("I am bar")
	if !bytes.Equal(op, expect) {
		t.Errorf("For template string %v, expected %v after variable injection but got %v", string(template), string(expect), string(op))
	}
}

func TestInjectVariableFalse(t *testing.T) {
	model := steps.CommandModel{}

	template := []byte("I am {{foo}}")
	model.LicenseText = template

	errMsg := fmt.Sprintf("For template %v, variable was not provided in model. It must throw error", string(template))
	runErrorThrowingModelTest("TestInjectVariableFalse", runInjectLicense, model, t, errMsg, false)
}

func runInjectLicense(m steps.CommandModel) {
	m.InjectVariable(m.LicenseText)
}
