package utils

const commonPrefix = "/* "
const commonPostfix = "*/"

// SupportedComments is mapping of language extension with syntax
var SupportedComments = map[string]([2]string){
	"h":      [2]string{commonPrefix, commonPostfix},
	"hpp":    [2]string{commonPrefix, commonPostfix},
	"html":   [2]string{"<!-- ", "-->"},
	"js":     [2]string{commonPrefix, commonPostfix},
	"jsx":    [2]string{commonPrefix, commonPostfix},
	"tsx":    [2]string{commonPrefix, commonPostfix},
	"css":    [2]string{commonPrefix, commonPostfix},
	"py":     [2]string{"\"\"\"\n", "\"\"\""},
	"java":   [2]string{commonPrefix, commonPostfix},
	"rb":     [2]string{"=begin ", "=end"},
	"c":      [2]string{commonPrefix, commonPostfix},
	"cpp":    [2]string{commonPrefix, commonPostfix},
	"cs":     [2]string{commonPrefix, commonPostfix},
	"m":      [2]string{commonPrefix, commonPostfix},
	"go":     [2]string{commonPrefix, commonPostfix},
	"swift":  [2]string{commonPrefix, commonPostfix},
	"clj":    [2]string{"(comment ", ")"},
	"ts":     [2]string{commonPrefix, commonPostfix},
	"dart":   [2]string{commonPrefix, commonPostfix},
	"elm":    [2]string{"{- ", "-}"},
	"groovy": [2]string{"/* ", "*/"},
	"hs":     [2]string{"{- ", "-}"},
	"kt":     [2]string{commonPrefix, commonPostfix},
	"rs":     [2]string{commonPrefix, commonPostfix},
}

// Comment will give comment prefix and postfix based on programming language
func Comment(programmingLanExtension string) (string, string) {
	cmnt := SupportedComments[programmingLanExtension]
	return cmnt[0], cmnt[1]
}
