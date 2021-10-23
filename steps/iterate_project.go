package steps

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Marvin9/licensor/utils"
)

var spaceNextLineRegex = regexp.MustCompile(`\s+|\n`)

// Start is main function to iterate in project and inject license
func (m *CommandModel) Start() {
	m.iterateDirectory(m.ProjectPath)
}

func (m *CommandModel) iterateDirectory(path string) {
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
			continue
		}

		if file.IsDir() {
			m.iterateDirectory(fullpath)
			continue
		}

		// FILE SPOTTED
		// PROCESS FILE ONLY IF IT HAS EXTENSIONS GIVEN IN COMMAND
		ext := utils.GetExtension(filename)
		if !utils.Exists(ext, m.Extensions) {
			continue
		}

		// PROCESS FILE
		// GENERATE COMMENT PREFIX & POSTFIX BASED ON EXTENSION
		commentPrefix, commentPostfix := utils.Comment(ext)
		var isSingleLine bool = (commentPostfix == utils.SINGLE_LINE_COMMENTS)

		uniqueHeader := append([]byte(commentPrefix), []byte(utils.DELIMITER)...)

		buffer := make([]byte, 512)
		file, err := os.OpenFile(fullpath, os.O_RDWR, os.ModePerm)
		if err != nil {
			utils.LogError(err)
		}

		var fileContent []byte
		var licenseAlreadyExist = -1
		for {
			read, err := file.Read(buffer)
			if err != nil {
				break
			}
			//read the file content
			content := buffer[:read]
			//the length of the file before append
			fileContentBeforeAppend := len(fileContent)
			//the file content with the content appended
			fileContent = append(fileContent, content...)

			if licenseAlreadyExist == -1 {
				exist := bytes.Index(content, uniqueHeader)
				if exist != -1 {
					licenseAlreadyExist = exist + fileContentBeforeAppend
				}
			}

			if read == 0 {
				break
			}
		}

		if licenseAlreadyExist != -1 {
			// PROCESS TO CHECK CURRENT LICENSE IS NOT EQUAL TO PREVIOUS ONE

			var endOfComment int
			uniqueHeaderLen := len(uniqueHeader)
			var postfix string
			//delimiter as the Index search
			if isSingleLine == true{
				postfix = commentPrefix + utils.DELIMITER 
			}else{
				postfix = utils.DELIMITER + commentPostfix 
			}

			endOfComment = bytes.Index(fileContent[licenseAlreadyExist+uniqueHeaderLen:], []byte(postfix))
			endOfComment += licenseAlreadyExist + uniqueHeaderLen
			oldLicenseText := bytes.TrimPrefix(fileContent[licenseAlreadyExist:endOfComment], uniqueHeader)

			if !m.RemoveFlag {
				null := []byte("")
				t1 := spaceNextLineRegex.ReplaceAll(oldLicenseText, null)
				t2 := spaceNextLineRegex.ReplaceAll(m.LicenseText, null)
				if bytes.Equal(t1, t2) {
					// BOTH ARE SAME LICENSE SO NO NEED TO CHANGE
					file.Close()
					continue
				}
			}

			lastIdx := endOfComment + len(postfix) - 1
			// REMOVE EXISTING LICENSE
			fileContent = append(fileContent[0:licenseAlreadyExist], fileContent[lastIdx+1:]...)
			fileContent = bytes.TrimPrefix(fileContent, []byte("\n"))
		} else if m.RemoveFlag {
			file.Close()
			continue
		}

		if !utils.IsWindows {
			// \u001b[2K => clear entire line
			// \033[u    => restore cursor position (position where porgram started)
			// \u001b[2K => clear entire line [to eliminate overflow issues]
			// \u001b[0G => place cursor to 0th position
			fmt.Printf("\u001b[2K\033[u\u001b[2K\u001b[0G%v\u001b[0G", fullpath)
		}
		
		fileToInjectLicense := file
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
		
		//copy the value of the license to a temporary string
		var finalString = string(m.LicenseText)
		println(finalString)
		if isSingleLine == true {
			//add delimeter to the end of the file and comment it out
			finalString = finalString + "\n" + commentPrefix + utils.DELIMITER	
			//add the unique identifier to the start of the license
			finalString =  commentPrefix + utils.UniqueIdentifier + "\n" + finalString
			//replace each newline with newline + comment
			finalString = strings.ReplaceAll(finalString, "\n ", "\n"+commentPrefix+" ")
		}else{
			//add the delimeter to the end of the license
			finalString = finalString + "\n" + utils.DELIMITER
			//make the unique identifier as part of the string
			finalString =  utils.UniqueIdentifier + "\n" + finalString


		}
		
		
		if !m.RemoveFlag {
			fileToInjectLicense.WriteString(fmt.Sprintf("%v%v%v\n", commentPrefix + utils.DELIMITER + "\n", finalString, commentPostfix))
		}
		fileToInjectLicense.Write(fileContent)
		file.Close()

		// DONE!
		// fmt.Printf("\nFile updated: %v", fullpath)
	}
}
