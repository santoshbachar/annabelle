package tasks

import (
	"fmt"
	"github.com/santoshbachar/annabelle/playbook/bash"
	"gopkg.in/yaml.v2"
)

type groupState string

const (
	Present groupState = "present"
	Absent             = "absent"
)

type Group struct {
	Name   string     `yaml:"name"`
	State  groupState `yaml:"state"`
	Gid    string     `yaml:"gid"`
	System string     `yaml:"system"`
}

func (g *Group) Unmarshal(file []byte) {
	err := yaml.Unmarshal(file, &g)
	if err != nil {
		panic(err)
	}
}

func (g *Group) Execute(loop Loop) bool {
	fmt.Println("group execute()")
	if g.Name == "" {
		fmt.Println("Group name is required")
		return false
	}
	if g.State == "" {
		fmt.Println("State is absent, taking \"present\" as default value")
		g.State = "present"
	}

	fmt.Println("Group Execute(). Validation done. Moving forward")

	switch g.State {
	case Present:
		fmt.Println("Group state is present")
		present(g.Name)
		break
	case Absent:
		fmt.Println("Group state is present")
		break
	}
	return true
}

func present(name string) {
	fmt.Println("name is ", name)

	if isAlreadyPresent(name) {
		return
	}

	ok := bash.Do("groupadd", name)
	if !ok {
		panic("error creating group")
	}
}

func isAlreadyPresent(name string) bool {
	ok := bash.Do("grep", "-q", name, "/etc/group")
	if ok {
		return true
	}
	return false
}
