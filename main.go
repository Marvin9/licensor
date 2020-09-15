package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/Marvin9/licensor/utils"

	"github.com/Marvin9/licensor/steps"
)

func main() {
	closeSignal()
	if runtime.GOOS == "windows" {
		utils.IsWindows = true
	}

	if !utils.IsWindows {
		fmt.Print("\033[s")    // save cursor position
		fmt.Print("\033[?25l") // hide cursor
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

	if utils.IsWindows {
		fmt.Print("Working...")
	}
	model.Start()

	// https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html#completeness
	if !utils.IsWindows {
		fmt.Print("\u001b[2K")
		fmt.Print("\u001b[0G")
	}
	fmt.Println("Done ✔️")
	utils.ShowCursor()
}

func closeSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		utils.ShowCursor()
		os.Exit(1)
	}()
}
