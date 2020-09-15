package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Marvin9/licensor/utils"

	"github.com/Marvin9/licensor/steps"
)

func main() {
	if runtime.GOOS == "windows" {
		utils.IsWindows = true
	}

	if !utils.IsWindows {
		fmt.Print("\033[s")    // save cursor position
		fmt.Print("\033[?25l") // hide cursor
	} else {
		fmt.Print("Working...")
	}

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
	if !utils.IsWindows {
		fmt.Print("\u001b[2K")
		fmt.Print("\u001b[0G")
	}
	fmt.Println("✔️")
	utils.ShowCursor()
}
