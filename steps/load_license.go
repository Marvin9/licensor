package steps

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Marvin9/licensor/utils"
)

// LoadLicense will store license text in CommandModel
// given value for -license flag
// 1. It will open and read file, if exist
// 2. If not then it will request on http
// if both not exist, give error
func (m *CommandModel) LoadLicense() []byte {
	var licenseText []byte
	_, errPath := os.Stat(m.License)

	licenseFileError := fmt.Sprintf("%v is either valid path nor valid url.", m.License)

	if errPath != nil {
		res, errURL := http.Get(m.License)
		if errURL != nil {
			utils.LogError(licenseFileError)
		}

		rd, errURL := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if errURL != nil {
			utils.LogError(errURL)
		}
		licenseText = rd
	} else {
		licenseFile, err := os.Open(m.License)
		if err != nil {
			utils.LogError(licenseFileError)
		}
		licenseTextLc, errFile := ioutil.ReadAll(licenseFile)
		licenseFile.Close()
		if errFile != nil {
			utils.LogError(licenseFileError)
		}

		licenseText = licenseTextLc
	}

	return licenseText
}
