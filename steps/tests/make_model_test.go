package tests

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/Marvin9/licensor/steps"

	"github.com/Marvin9/licensor/utils"
)

func TestMakeModel(t *testing.T) {
	license := "foo.txt"
	inputExts := []string{"go", "c", "sh"}
	arg := []string{"", utils.PROJECT, "./mock", utils.EXT, inputExts[0], inputExts[1], inputExts[2], utils.LICENSE, license, utils.IGNORE, "a", "b", utils.TEMPLATE, "{\"name\": \"mayur\"}"}

	var model steps.CommandModel
	model.MakeModel(arg)

	if model.ProjectPath != "./mock" {
		t.Errorf("ProjectPath of model expected %v got %v", "./mock", model.ProjectPath)
	}

	extLen := len(model.Extensions)
	if extLen != 3 {
		t.Errorf("Extensions of model, expected length %v got %v", 3, extLen)
	}

	for i, ext := range inputExts {
		if model.Extensions[i] != ext {
			t.Errorf("At index %v, expected %v extension got %v", i, ext, model.Extensions[i])
		}
	}

	if model.License != license {
		t.Errorf("License of model expected %v got %v", license, model.License)
	}

	ignoreLen := len(model.Ignore)
	if ignoreLen != 2 {
		t.Errorf("Ignore of model, expected length %v got %v", 2, ignoreLen)
	}

	val, ok := model.Template["name"]
	if !ok {
		t.Errorf("name value in -template was passed as json, but not stored in model")
	}

	if val != "mayur" {
		t.Errorf("Expected value for key \"name\", %v, got %v", "mayur", val)
	}
}

func TestInvalidProjectTag_1(t *testing.T) {
	args := []string{"", utils.PROJECT}
	runErrorThrowingTest("TestInvalidProjectTag_1", invalidTags, args, t, "Empty flag value of -project should exit with error.")
}

func TestInvalidProjectTag_2(t *testing.T) {
	args := []string{"", utils.PROJECT, utils.LICENSE}
	runErrorThrowingTest("TestInvalidProjectTag_2", invalidTags, args, t, "Value of flag for flag should exit with error")
}

func TestInvalidLicenseTag(t *testing.T) {
	args := []string{"", utils.LICENSE}
	runErrorThrowingTest("TestInvalidLicenseTag", invalidTags, args, t, "Empty flag value of -license should exit with error.")
}

func TestInvalidTemplateTag(t *testing.T) {
	args := []string{"", utils.TEMPLATE, "invalid json"}
	runErrorThrowingTest("TestInvalidTemplateTag", invalidTags, args, t, "Invalid json value of -template should exit with error.")
}

func invalidTags(args []string) {
	var model steps.CommandModel
	model.MakeModel(args)
}

func runErrorThrowingTest(testName string, f func([]string), fArgs []string, t *testing.T, errMsg string) {
	if os.Getenv("CRASHER") == "1" {
		f(fArgs)
		return
	}
	cmd := exec.Command(os.Args[0], fmt.Sprintf("-test.run=%v", testName))
	cmd.Env = append(os.Environ(), "CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Errorf(fmt.Sprintf("%v.\nInvalid use of flag: %v", errMsg, fArgs))
}
