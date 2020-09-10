package tests

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/Marvin9/licensor/steps"
	"github.com/Marvin9/licensor/utils"
)

func TestIterateProject(t *testing.T) {
	model := steps.CommandModel{
		ProjectPath: "../../mock",
		Extensions:  []string{"go"},
		License:     "../../mock/License-1.txt",
	}

	model.Validate()
	lc := model.LoadLicense()
	lc = model.Pretty(lc)

	model.LicenseText = lc

	mainFile := "../../mock/main.go"
	barFile := "../../mock/foo/bar.go"

	mainMockFileBefore, err := ioutil.ReadFile(mainFile)
	if err != nil {
		t.Error(err)
	}
	barMockFileBefore, err := ioutil.ReadFile(barFile)
	if err != nil {
		t.Error(err)
	}

	model.Start()

	commentPrefix, commentPostfix := utils.Comment("go")

	licenseShouldBe := append([]byte(commentPrefix), []byte(utils.UniqueIdentifier)...)
	licenseShouldBe = append(licenseShouldBe, []byte("\n")...)
	licenseShouldBe = append(licenseShouldBe, lc...)
	licenseShouldBe = append(licenseShouldBe, []byte("\n")...)
	licenseShouldBe = append(licenseShouldBe, []byte(commentPostfix)...)
	licenseShouldBe = append(licenseShouldBe, []byte("\n\n")...)

	mainMockFileAfter, err := ioutil.ReadFile(mainFile)
	expect := append(licenseShouldBe, mainMockFileBefore...)

	if !bytes.Equal(expect, mainMockFileAfter) {
		t.Errorf("Expected /mock/main.go: \n%v\nGot: \n%v", string(expect), string(mainMockFileAfter))
	}

	barMockFileAfter, err := ioutil.ReadFile(barFile)
	expect = append(licenseShouldBe, barMockFileBefore...)

	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(expect, barMockFileAfter) {
		t.Errorf("Expected /mock/foo/bar.go: %v\nGot: %v", string(expect), string(barMockFileAfter))
	}
}
