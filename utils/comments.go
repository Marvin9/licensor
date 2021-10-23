package utils

const commonPrefix = "/* "
const commonPostfix = "*/"

// SupportedComments is mapping of language extension with syntax
var SupportedComments = map[string]([2]string){
	"html":     [2]string{"<!-- ", "-->"},              //html files
	"js":       [2]string{commonPrefix, commonPostfix}, //javascript files
	"jsx":      [2]string{commonPrefix, commonPostfix}, //jsx files
	"tsx":      [2]string{commonPrefix, commonPostfix}, //tsx files
	"css":      [2]string{commonPrefix, commonPostfix}, //css files
	"py":       [2]string{"\"\"\"\n", "\"\"\""},        //python files
	"java":     [2]string{commonPrefix, commonPostfix}, //java files
	"rb":       [2]string{"=begin ", "=end"},           //ruby files
	"c":        [2]string{commonPrefix, commonPostfix}, //c files
	"cpp":      [2]string{commonPrefix, commonPostfix}, //c++ files
	"h":        [2]string{commonPrefix, commonPostfix}, //c/c++ header files
	"hpp":      [2]string{commonPrefix, commonPostfix}, //c++ header files
	"cs":       [2]string{commonPrefix, commonPostfix}, //c# files
	"m":        [2]string{commonPrefix, commonPostfix}, //objective c files
	"go":       [2]string{commonPrefix, commonPostfix}, //go files
	"swift":    [2]string{commonPrefix, commonPostfix}, //swift files
	"clj":      [2]string{"(comment ", ")"},            //clojure files
	"ts":       [2]string{commonPrefix, commonPostfix}, //typescript files
	"dart":     [2]string{commonPrefix, commonPostfix}, //dart files
	"elm":      [2]string{"{- ", "-}"},                 //elm files
	"groovy":   [2]string{"/* ", "*/"},                 //groovy files
	"hs":       [2]string{"{- ", "-}"},                 //haskel files
	"kt":       [2]string{commonPrefix, commonPostfix}, //kotlin files
	"rs":       [2]string{commonPrefix, commonPostfix}, //rust files
	"sh":       [2]string{"#", SINGLE_LINE_COMMENTS},   //shell files
	"Makefile": [2]string{"#", SINGLE_LINE_COMMENTS},   //TODO makefiles still don't work due to them not actually having a file extension
	"f":        [2]string{"!", SINGLE_LINE_COMMENTS},   //fortran files ext. 1
	"for":      [2]string{"!", SINGLE_LINE_COMMENTS},   //fortran files ext. 2
}

// Comment will give comment prefix and postfix based on programming language
func Comment(programmingLanExtension string) (string, string) {
	cmnt := SupportedComments[programmingLanExtension]
	return cmnt[0], cmnt[1]
}
