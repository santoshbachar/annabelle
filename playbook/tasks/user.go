package tasks

import (
	"fmt"
	"github.com/santoshbachar/annabelle/playbook/bash"
	"gopkg.in/yaml.v2"
)

type User struct {
	Name    string `yaml:"name"`
	Comment string `yaml:"comment"`
	Group   string `yaml:"group"`
}

func (u *User) Unmarshal(file []byte) {
	err := yaml.Unmarshal(file, &u)
	if err != nil {
		panic(err)
	}
}

func validate(u *User) bool {
	if u.Name == "" {
		fmt.Println("User must specify a name")
		return false
	}
	if u.Group == "" {
		fmt.Println("Need to specify the group to add the user")
		return false
	}
	return true
}

func (u *User) Execute(loop Loop) bool {
	if !validate(u) {
		return false
	}

	if userExists(u.Name) {
		return true
	}

	userAdd(u.Name, u.Group)

	return true
}

func userAdd(user, group string) {
	//ok := bash.Do("useradd", name)
	ok := bash.Do("useradd", "-g", group, user)
	if !ok {
		fmt.Println("error creating user")
	}
}

func userExists(name string) bool {
	ok := bash.Do("id", name)
	if ok {
		return true
	}
	return false
}
