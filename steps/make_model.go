package steps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Marvin9/licensor/utils"
	"gopkg.in/yaml.v2"
)

// MakeModel will generate CommandModel based on given argument
// OR yml file
func (m *CommandModel) MakeModel(args []string) {
	mainArgs := args[1:]
	i := 0
	mainArgsLen := len(mainArgs)

	if mainArgsLen == 0 {
		yamlContent, err := ioutil.ReadFile(utils.LicensorYAML)
		if err != nil {
			utils.LogError(fmt.Sprintf(`
%v file does not exist. 
Read https://github.com/Marvin9/licensor/README.md or licensor -help`, utils.LicensorYAML))
		}
		err = yaml.Unmarshal(yamlContent, m)
		if err != nil {
			utils.LogError(err)
		}
		return
	}

	for i < mainArgsLen {
		arg := mainArgs[i]

		switch arg {
		case utils.PROJECT:
			i++
			if i >= mainArgsLen || utils.IsKeywordCommand(mainArgs[i]) {
				utils.InvalidFlagError(utils.PROJECT)
			}
			m.ProjectPath = mainArgs[i]
			i++
		case utils.EXT:
			i++
			for i < mainArgsLen && !utils.IsKeywordCommand(mainArgs[i]) {
				m.Extensions = append(m.Extensions, mainArgs[i])
				i++
			}
		case utils.LICENSE:
			i++
			if i >= mainArgsLen || utils.IsKeywordCommand(mainArgs[i]) {
				utils.InvalidFlagError(utils.LICENSE)
			}

			m.License = mainArgs[i]
			i++
		case utils.IGNORE:
			i++
			for i < mainArgsLen && !utils.IsKeywordCommand(mainArgs[i]) {
				m.Ignore = append(m.Ignore, mainArgs[i])
				i++
			}
		case utils.TEMPLATE:
			i++
			if i >= mainArgsLen || utils.IsKeywordCommand(mainArgs[i]) {
				utils.InvalidFlagError(utils.TEMPLATE)
			}
			err := json.Unmarshal([]byte(mainArgs[i]), &m.Template)
			if err != nil {
				utils.LogError(err)
			}
			i++
		case utils.REMOVE:
			m.RemoveFlag = true
			i++
		case utils.HELP:
			fmt.Print(utils.HelpDocs)
			os.Exit(0)
		default:
			i++
		}
	}
}
