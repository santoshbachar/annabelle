package tasks

import "gopkg.in/yaml.v2"

type DockerImage struct {
	Name     string `yaml:"name"`
	State    string `yaml:"state"`
	Tag      string `yaml:"tag"`
	LoadPath string `yaml:"load_path"`
	Source   string `yaml:"source"`
}

func (di *DockerImage) Unmarshal(file []byte) {
	err := yaml.Unmarshal(file, &di)
	if err != nil {
		panic(err)
	}
}

func (di *DockerImage) Execute(loop Loop) bool {
	return true
}
