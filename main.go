package main

import (
	"fmt"
	"os"

	"github.com/Marvin9/licensor/steps"
)

func main() {
	fmt.Print("\033[s")    // save cursor position
	fmt.Print("\033[?25l") // hide cursor

	// find . | grep -i "\(\.go\|\.sh\)$" | wc -l

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

	// https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html#completeness
	fmt.Print("\u001b[2K")
	fmt.Print("\u001b[0G")
	fmt.Println("✔️")
	fmt.Print("\033[?25h")
}
