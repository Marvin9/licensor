package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Marvin9/licensor/steps"
)

func main() {
	// find . | grep -i "\(\.go\|\.sh\)$" | wc -l
	fmt.Print("Working")
	go (func() {
		for {
			for i := 0; i < 3; i++ {
				fmt.Print(".")
				time.Sleep(time.Millisecond * 500)
			}
			for i := 0; i < 3; i++ {
				fmt.Print("\b \b")
			}
		}
	})()
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
	fmt.Print("\r          \r")
	fmt.Println("✔️")
}
