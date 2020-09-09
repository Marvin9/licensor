package steps

type CommandModel struct {
	ProjectPath string
	Extensions  []string
	License     string
	Ignore      []string
	LicenseText []byte
	Template    map[string]string
}
