package tasks

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

type File struct {
	Path  string `yaml:"path"`
	State string `yaml:"state"`
	//Recurse Boolean `yaml:"recurse"`
	Recurse string `yaml:"recurse"`
	Owner   string `yaml:"owner"`
	Group   string `yaml:"group"`
	Mode    string `yaml:"mode"`
}

func HandleFile(file []byte) {
	//file := File{x}
	fmt.Println("HandleFile*********")
	fs := File{}
	err := yaml.Unmarshal([]byte(file), &fs)
	if err != nil {
		fmt.Println("err while Unmarshall")
	}
	fmt.Println(fs)
}
