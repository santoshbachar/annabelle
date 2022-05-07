package tasks

import (
	"github.com/fatih/color"
	"github.com/santoshbachar/annabelle/playbook/bash"
	"github.com/santoshbachar/annabelle/playbook/constants"
	"github.com/santoshbachar/annabelle/utility"
	"gopkg.in/yaml.v2"
)

type Script struct {
	//Name   string `yaml:"name"`
	Role   string
	Script string //`yaml:"script"`
	Path   string
}

//type Script string

func (s *Script) Unmarshal(file []byte) {
	err := yaml.Unmarshal(file, &s)
	if err != nil {
		panic(err)
	}
}

func (s *Script) init(roleName, fileName string) {
	s.Role = roleName
	s.Script = fileName
	s.Path = constants.ResourceDir + "roles/" + s.Role + "/files/" + s.Script
}

func validateScript(s *Script) bool {
	if s.Script == "" {
		return false
	}
	return utility.FileExists(s.Path)
}

func (s *Script) Execute(loop Loop) bool {
	if !validateScript(s) {
		color.Red("validation false, must be because of missing script in files directory")
		return false
	}

	runScript(s.Path)

	return true
}

func runScript(path string) {
	bash.Run(path)
}
