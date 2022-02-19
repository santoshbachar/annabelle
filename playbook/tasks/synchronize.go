package tasks

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

type Synchronize struct {
	Source      string   `yaml:"src"`
	Destination string   `yaml:"dest"`
	Delete      Boolean  `yaml:"delete"`
	RSyncOpts   []string `yaml:"rsync_opts"`
}

func HandleSynchronize(file []byte) {
	//file := File{x}
	fmt.Println("HandleFile*********")
	s := Synchronize{}
	err := yaml.Unmarshal([]byte(file), &s)
	if err != nil {
		fmt.Println("err while Unmarshall")
	}
	fmt.Println(s)
}
