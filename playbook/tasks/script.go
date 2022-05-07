package tasks

import (
	"fmt"
	"github.com/santoshbachar/annabelle/playbook/bash"
	"github.com/santoshbachar/annabelle/utility"
	"gopkg.in/yaml.v2"
)

type Script struct {
	//Name   string `yaml:"name"`
	Script string `yaml:"script"`
}

//type Script string

func (s *Script) Unmarshal(file []byte) {
	err := yaml.Unmarshal(file, &s)
	if err != nil {
		panic(err)
	}
}

func (s *Script) init(name string) {
	s.Script = name
}

func validateScript(s *Script) bool {
	if s.Script == "" {
		return false
	}
	path := "../files/" + s.Script
	return utility.FileExists(path)
}

func (s *Script) Execute(loop Loop) bool {
	if !validateScript(s) {
		fmt.Println("validation false, must be because of missing script in files directory")
		return false
	}

	runScript(s.Script)

	return true
}

func runScript(path string) {
	bash.Run(path)
}
