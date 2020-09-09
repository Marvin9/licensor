package steps

import (
	"fmt"
	"os"

	"github.com/Marvin9/licensor/utils"
)

func (m *CommandModel) Validate() {
	if len(m.Extensions) == 0 {
		utils.LogError("You must provide atleast one valid extension to -ext flag.")
	}

	// project path must exist
	projectDir, err := os.Stat(m.ProjectPath)
	if err != nil {
		utils.LogError(err)
	}

	// project path must be directory
	if !projectDir.IsDir() {
		utils.LogError(fmt.Sprintf("%v is not directory.", m.ProjectPath))
	}

	// extensions must be valid and implemented
	for _, ext := range m.Extensions {
		if !utils.IsValidExtension(ext) {
			utils.LogError(fmt.Sprintf("We do not support %v extension right now.", ext))
		}
	}
}
