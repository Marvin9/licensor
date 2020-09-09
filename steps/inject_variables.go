package steps

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/Marvin9/licensor/utils"
)

// InjectVariable will replace template string with variables
// for example string: I am {{name}}
// now inside CommandModel m: m.Template["name"] = "Mayur"
// output: I am Mayur
func (m *CommandModel) InjectVariable(licenseText []byte) []byte {
	templateReg := regexp.MustCompile(`{{[[:alpha:]]+}}`)
	templateMatches := templateReg.FindAll(licenseText, -1)

	for _, match := range templateMatches {
		// {{foo}}
		matchStr := string(match)

		// foo
		variableName := matchStr[2 : len(matchStr)-2]
		variableValue, defined := m.Template[variableName]
		if !defined {
			utils.LogError(fmt.Sprintf("%v is not defined in template.", variableName))
		}
		licenseText = bytes.Replace(licenseText, match, []byte(variableValue), 1)
	}

	return licenseText
}
