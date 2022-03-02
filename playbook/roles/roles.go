package roles

import (
	"fmt"
	"github.com/santoshbachar/annabelle/playbook/constants"
	"os"
)

func LoadRoles(roles []string) []string {
	var okRoles []string
	for _, role := range roles {
		ok := mainFileExists(role)
		if ok != true {
			fmt.Println("no playbook for " + role)
		} else {
			fmt.Println(role + " ready to parse")
			okRoles = append(okRoles, role)
		}
	}
	return okRoles
}

func mainFileExists(name string) bool {
	var path = constants.ResourceDir + "roles/" + name + "/tasks/main.yaml"
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
