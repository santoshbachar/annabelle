package tasks

import (
	"fmt"
	"github.com/santoshbachar/annabelle/playbook/helper"
	"gopkg.in/yaml.v2"
)

type Synchronize struct {
	Source      string   `yaml:"src"`
	Destination string   `yaml:"dest"`
	Delete      Boolean  `yaml:"delete"`
	RSyncOpts   []string `yaml:"rsync_opts"`
}

func (s *Synchronize) Unmarshal(file []byte) {
	err := yaml.Unmarshal([]byte(file), &s)
	if err != nil {
		panic(err)
	}
}

func (s Synchronize) Execute() bool {
	if s.Source == "" {
		fmt.Println("From where?")
		return false
	}

	if s.Destination == "" {
		fmt.Println("Where to?")
		return false
	}

	syncfiles(&s.Source, &s.Destination)

	return true
}

func syncfiles(source *string, destination *string) {
	ok, variable := helper.FindVariable(*destination)
	if ok == false {
		return
	}
	var values = []string{"val1", "val2"}

	ok, newVals := helper.FindVariableAndReplaceWithValue(destination, &variable, values)
	fmt.Println(newVals)
}

//func (s Synchronize) AddLoopItems(items interface{}) {
//
//}

//func HandleSynchronize(file []byte) {
//	//file := File{x}
//	fmt.Println("HandleFile*********")
//	s := Synchronize{}
//	err := yaml.Unmarshal([]byte(file), &s)
//	if err != nil {
//		fmt.Println("err while Unmarshall")
//	}
//	fmt.Println(s)
//}
