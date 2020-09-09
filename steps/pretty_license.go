package steps

import "bytes"

func (m *CommandModel) Pretty(licenseText []byte) []byte {
	splitted := bytes.Split(licenseText, []byte("\n"))
	licenseText = []byte("")
	for _, line := range splitted {
		licenseText = append(licenseText, append([]byte("\n "), line...)...)
	}
	return licenseText
}
