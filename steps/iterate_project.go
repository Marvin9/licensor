package steps

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/Marvin9/licensor/utils"
)

var spaceNextLineRegex = regexp.MustCompile(`\s+|\n`)

var wg sync.WaitGroup

// Start is main function to iterate in project and inject license
func (m *CommandModel) Start() {
	wg.Add(1)
	go m.iterateDirectory(m.ProjectPath)
	wg.Wait()
}

func (m *CommandModel) iterateDirectory(path string) {
	defer wg.Done()
	files, err := ioutil.ReadDir(path)
	if err != nil {
		utils.LogError(err)
	}

	for _, file := range files {
		filename := file.Name()
		fullpath := path + "/" + filename
		base := filepath.Base(fullpath)

		if utils.ShouldIgnoreDir(base) {
			continue
		}

		if utils.Exists(fullpath, m.Ignore) {
			fmt.Println(fullpath)
			continue
		}

		if file.IsDir() {
			wg.Add(1)
			go m.iterateDirectory(fullpath)
			continue
		}

		// FILE SPOTTED
		// PROCESS FILE ONLY IF IT HAS EXTENSIONS GIVEN IN COMMAND
		ext := utils.GetExtension(filename)
		if !utils.Exists(ext, m.Extensions) {
			continue
		}

		// PROCESS FILE

		// GET FILE CONTENT
		fileContent, err := ioutil.ReadFile(fullpath)
		if err != nil {
			utils.LogError(err)
		}

		// GENERATE COMMENT PREFIX & POSTFIX BASED ON EXTENSION
		commentPrefix, commentPostfix := utils.Comment(ext)

		uniqueHeader := append([]byte(commentPrefix), []byte(utils.UniqueIdentifier)...)

		// CHECK IF THERE IS ALREADY LICENSE GENERATED PREVIOUSLY
		licenseAlreadyExist := bytes.Index(fileContent, uniqueHeader)

		if licenseAlreadyExist != -1 {
			// PROCESS TO CHECK CURRENT LICENSE IS NOT EQUAL TO PREVIOUS ONE
			endOfComment := bytes.Index(fileContent, []byte(commentPostfix))
			oldLicenseText := bytes.TrimPrefix(fileContent[licenseAlreadyExist:endOfComment], uniqueHeader)

			if !m.RemoveFlag {
				null := []byte("")
				t1 := spaceNextLineRegex.ReplaceAll(oldLicenseText, null)
				t2 := spaceNextLineRegex.ReplaceAll(m.LicenseText, null)
				if bytes.Equal(t1, t2) {
					// BOTH ARE SAME LICENSE SO NO NEED TO CHANGE
					continue
				}
			}

			lastIdx := endOfComment + len(commentPostfix) - 1
			// REMOVE EXISTING LICENSE
			fileContent = append(fileContent[0:licenseAlreadyExist], fileContent[lastIdx+1:len(fileContent)]...)
			fileContent = bytes.TrimPrefix(fileContent, []byte("\n\n"))
		} else if m.RemoveFlag {
			continue
		}

		fileToInjectLicense, err := os.OpenFile(fullpath, os.O_WRONLY, os.ModePerm)
		if err != nil {
			utils.LogError(err)
		}
		fileToInjectLicense.Truncate(0)
		fileToInjectLicense.Seek(0, 0)

		// COMMENT OUT LICENSE TEXT
		// ---------------------- template --------------------------
		// commentPrefix uniqueIdentifier
		// license text
		// commentPostfix
		//
		//
		// actual code
		// -----------------------------------------------------------
		if !m.RemoveFlag {
			fileToInjectLicense.WriteString(fmt.Sprintf("%v%v\n%v\n%v\n\n", commentPrefix, utils.UniqueIdentifier, string(m.LicenseText), commentPostfix))
		}
		fileToInjectLicense.Write(fileContent)
		fileToInjectLicense.Close()

		// DONE!
		// fmt.Printf("\nFile updated: %v", fullpath)
	}
}
