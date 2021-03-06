package steps

// CommandModel is used to control input
type CommandModel struct {
	ProjectPath string            `yaml:"project"`
	Extensions  []string          `yaml:"extensions"`
	License     string            `yaml:"license"`
	Ignore      []string          `yaml:"ignore"`
	Template    map[string]string `yaml:"template"`
	LicenseText []byte
	RemoveFlag  bool
}
