package tests

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/Marvin9/licensor/steps"
)

func TestValidationTrue(t *testing.T) {
	model := steps.CommandModel{
		ProjectPath: ".",
		Extensions:  []string{"go"},
		License:     "foo.txt",
	}

	errMsg := fmt.Sprintf("%v is valid model, but still throw error.", model)
	runErrorThrowingValidationTest("TestValidationTrue", runValidation, model, t, errMsg, true)
}

func TestValidationTrue_NoProjectPath(t *testing.T) {
	model := steps.CommandModel{
		Extensions: []string{"go"},
		License:    "foo.txt",
	}

	errMsg := fmt.Sprintf("%v is valid model, ProjectPath is not provided, but it must set default value instead of throwing error.", model)
	runErrorThrowingValidationTest("TestValidationTrue_NoProjectPath", runValidation, model, t, errMsg, true)
}

func TestValidationFalse_NoExtension(t *testing.T) {
	model := steps.CommandModel{
		ProjectPath: ".",
		License:     "foo.txt",
		Extensions:  []string{},
	}

	errMsg := fmt.Sprintf("%v is invalid model, because it contains 0 extension, it must throw error.", model)
	runErrorThrowingValidationTest("TestValidationFalse_NoExtension", runValidation, model, t, errMsg, false)
}

func TestValidationFalse_ProjectPathDir(t *testing.T) {
	model := steps.CommandModel{
		ProjectPath: "./main.go",
		License:     "foo.txt",
		Extensions:  []string{"go"},
	}

	errMsg := fmt.Sprintf("%v is file, expected directory. It must throw error.", model.ProjectPath)
	runErrorThrowingValidationTest("TestValidationFalse_ProjectPathDir", runValidation, model, t, errMsg, false)
}

func TestValidationFalse_InvalidExtension(t *testing.T) {
	model := steps.CommandModel{
		ProjectPath: ".",
		Extensions:  []string{"invalid_ext"},
		License:     "foo.txt",
	}

	errMsg := fmt.Sprintf("%v does not contain valid extension(s). It should throw error.", model.Extensions)
	runErrorThrowingValidationTest("TestValidationFalse_InvalidExtension", runValidation, model, t, errMsg, false)
}

func runValidation(m steps.CommandModel) {
	m.Validate()
}

func runErrorThrowingValidationTest(testName string, f func(m steps.CommandModel), args steps.CommandModel, t *testing.T, errMsg string, shouldThisSuccess bool) {
	if os.Getenv("CRASHER") == "1" {
		f(args)
		return
	}
	cmd := exec.Command(os.Args[0], fmt.Sprintf("-test.run=%v", testName))
	cmd.Env = append(os.Environ(), "CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		if shouldThisSuccess {
			t.Errorf(errMsg)
		}
		return
	}
	if !shouldThisSuccess {
		t.Errorf(fmt.Sprintf("%v.\nInvalid model passed test: %v", errMsg, args))
	}
}
