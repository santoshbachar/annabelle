package playbook

import (
	"github.com/santoshbachar/annabelle/playbook/constants"
	"github.com/santoshbachar/annabelle/playbook/roles"
	tasks2 "github.com/santoshbachar/annabelle/playbook/tasks"
	"gopkg.in/yaml.v2"
	"os"
)

type tasks struct {
	Name string `yaml:"name"`
	Fail struct {
		Msg string `yaml:"msg"`
	} `yaml:"fail"`
	When      string   `yaml:"when"`
	WithItems []string `yaml:with_items`
}

type Playbook struct {
	Name         string   `yaml:"name"`
	Hosts        string   `yaml:"hosts"`
	GatherFacts  bool     `yaml:"gather_facts"`
	Become       bool     `yaml:"become"`
	BecomeMethod string   `yaml:"become_method"`
	Tasks        []tasks  `yaml:"tasks"`
	Roles        []string `yaml:"roles"`
}

func LoadPlaybook() {
	// playbooks := Playbook{}

	var playbooks []Playbook

	dat, err := os.ReadFile(constants.ResourceDir + "roles-def.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(dat), &playbooks)
	if err != nil {
		panic(err)
	}

	//fmt.Println("entire playbooks: ", playbooks[0].Tasks)

	for _, playbook := range playbooks {
		if len(playbook.Roles) > 0 {
			for _, task := range roles.LoadRoles(playbook.Roles) {
				tasks2.LoadTasks(task)
			}
		}
	}
}
