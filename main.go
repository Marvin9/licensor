package main

import (
	"os"

	"github.com/Marvin9/licensor/steps"
)

func main() {
	var model steps.CommandModel
	model.MakeModel(os.Args)
	model.Validate()
	if !model.RemoveFlag {
		lc := model.LoadLicense()
		lc = model.InjectVariable(lc)
		lc = model.Pretty(lc)
		model.LicenseText = lc
	}

	model.Start()
}
