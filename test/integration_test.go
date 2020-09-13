package test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/Marvin9/licensor/steps"
)

func TestIntegration(t *testing.T) {
	args := []string{
		"",
		"-project", "../mock/integration",
		"-ext", "go", "py",
		"-license", "https://gist.githubusercontent.com/Marvin9/ea5bb6b292cdccd7deb21d11509fcc56/raw/0ea1894244767b8f4ebf3644647524eb664807c9/Apache-2",
		"-template", "{\"year\": \"2020\", \"owner\": \"mayursinh\"}",
	}

	goFilePath := "../mock/integration/integration.go"
	pyFilePath := "../mock/integration/integration.py"

	filesToTest := []string{goFilePath, pyFilePath}

	for _, fileToTest := range filesToTest {

		var model steps.CommandModel
		model.MakeModel(args)
		model.Validate()
		lc := model.LoadLicense()
		lc = model.InjectVariable(lc)
		lc = model.Pretty(lc)
		model.LicenseText = lc

		model.Start()

		file, err := ioutil.ReadFile(fileToTest)
		if err != nil {
			t.Error(err)
		}

		isLicenseInGoFile := bytes.Index(file, lc)
		if isLicenseInGoFile == -1 {
			t.Errorf("License is not included in file %v.\n%v", fileToTest, string(file))
		}

		// it should not overwrite
		model.Start()

		file, err = ioutil.ReadFile(fileToTest)
		if err != nil {
			t.Error(err)
		}

		licenseCountInFile := bytes.Count(file, lc)
		if licenseCountInFile != 1 {
			t.Errorf("License in %v was added %v times. For existing license, it should not append", fileToTest, licenseCountInFile)
		}

		// update license
		model.Template["year"] = "2021"
		model.Template["owner"] = "FOO"

		model.Validate()
		updatedLc := model.LoadLicense()
		updatedLc = model.InjectVariable(updatedLc)
		updatedLc = model.Pretty(updatedLc)
		model.LicenseText = updatedLc

		model.Start()

		file, err = ioutil.ReadFile(fileToTest)
		if err != nil {
			t.Error(err)
		}

		// it should not contain previous license
		prevLicenseCount := bytes.Count(file, lc)
		if prevLicenseCount != 0 {
			t.Errorf("Previous license should be replaced. but found.\n%v", string(file))
		}

		// it should contain new license
		newLicenseCount := bytes.Count(file, updatedLc)
		if newLicenseCount != 1 {
			t.Errorf("License was not updated.\n%v", string(file))
		}

		// remove license
		model.RemoveFlag = true
		model.Start()

		file, err = ioutil.ReadFile(fileToTest)
		if err != nil {
			t.Error(err)
		}

		licenseCountInFile = bytes.Count(file, updatedLc)
		if licenseCountInFile != 0 {
			t.Errorf("License was not removed from file %v.", fileToTest)
		}
	}
}
