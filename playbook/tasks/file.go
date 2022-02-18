package tasks

type File struct {
	Path string `yaml:"path"`
	State string `yaml:"state"`
	Recurse Boolean `yaml:"recurse"`
	Owner string `yaml:"owner"`
	Group string `yaml:"group"`
	Mode string `yaml:"mode"`
}