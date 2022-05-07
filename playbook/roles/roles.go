package roles

import (
	"fmt"
	"github.com/santoshbachar/annabelle/playbook/constants"
	"github.com/santoshbachar/annabelle/utility"
)

func LoadRoles(roles []string) []string {
	var okRoles []string
	for _, role := range roles {
		path := getTaskMainYamlPath(role)
		ok := utility.FileExists(path)
		if ok != true {
			fmt.Println("no playbook for " + role)
		} else {
			fmt.Println(role + " ready to parse")
			okRoles = append(okRoles, role)
		}
	}
	return okRoles
}

func getTaskMainYamlPath(name string) string {
	var path = constants.ResourceDir + "roles/" + name + "/tasks/main.yaml"
	return path
}
