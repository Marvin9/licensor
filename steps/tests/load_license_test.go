package tests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/Marvin9/licensor/steps"
)

const url = "https://gist.githubusercontent.com/Marvin9/83781b0e416d0e554fea935498ae8ea2/raw/99dd6377b294ef99b287d2b36dbe83ffe25de0f5/foo"

func TestLoadLicenseTrue(t *testing.T) {
	model := steps.CommandModel{
		ProjectPath: "./",
		License:     "../../mock/License-1.txt",
		Extensions:  []string{"go"},
	}

	lc := model.LoadLicense()

	byteLicense, err := ioutil.ReadFile(model.License)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(lc, byteLicense) {
		t.Errorf("Provided path for license %v, Not loaded same.\nExpected: %v\nGot: %v", model.License, string(byteLicense), string(lc))
	}

	model = steps.CommandModel{
		ProjectPath: "./",
		License:     url,
		Extensions:  []string{"go"},
	}

	lc = model.LoadLicense()

	rt := []byte("foo")
	if !bytes.Equal(lc, rt) {
		t.Errorf("URL %v was passed as license, Expected %v, got %v.", url, string(rt), string(lc))
	}
}

func TestLoadLicenseFalse_invalidPath(t *testing.T) {
	model := steps.CommandModel{
		License:    "al;skjdf;asjkf;saj",
		Extensions: []string{"go"},
	}

	errMsg := fmt.Sprintf("%v is invalid path, it must throw error.", model.License)
	runErrorThrowingModelTest("TestLoadLicenseFalse_invalidPath", runLicense, model, t, errMsg, false)
}

func TestLoadLicenseFalse_invalidTextURL(t *testing.T) {
	model := steps.CommandModel{
		License:    "http://httpstat.us/200",
		Extensions: []string{"go"},
	}

	errMsg := fmt.Sprintf("%v returned empty text, it must throw error.", model.License)
	runErrorThrowingModelTest("TestLoadLicenseFalse_invalidTextURL", runLicense, model, t, errMsg, false)
}

func runLicense(m steps.CommandModel) {
	m.LoadLicense()
}
