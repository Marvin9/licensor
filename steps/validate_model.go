package steps

import (
	"fmt"
	"net/url"
	"os"

	"github.com/Marvin9/licensor/utils"
)

// Validate command model
// 1. -ext should have atleast 1 value, otherwise command will ignore all files
// 2. path given in -project flag must be directory and must exist
// 3. for given -ext, they should be valid and implemented in our app
func (m *CommandModel) Validate() {
	if len(m.Extensions) == 0 {
		utils.LogError(`
You must provide atleast one valid extension to -ext flag.`)
	}

	if m.ProjectPath == "" {
		m.ProjectPath = "."
	}

	// project path must exist
	projectDir, err := os.Stat(m.ProjectPath)
	if err != nil {
		utils.LogError(err)
	}

	// project path must be directory
	if !projectDir.IsDir() {
		utils.LogError(fmt.Sprintf(`
%v is not directory.`, m.ProjectPath))
	}

	// extensions must be valid and implemented
	for _, ext := range m.Extensions {
		if !utils.IsValidExtension(ext) {
			utils.LogError(fmt.Sprintf(`
We do not support %v extension right now. 
Open issue: https://github.com/Marvin9/licensor/issues/new?title=%v%v`, ext, url.QueryEscape("Support Extension "), ext))
		}
	}
}
