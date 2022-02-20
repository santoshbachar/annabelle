package tasks

import (
	"fmt"
	"github.com/santoshbachar/annabelle/playbook/constants"
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
)

//type TaskType string

const (
	FILE        = "file"
	SYNCHRONIZE = "synchronize"
)

type HandleTask interface {
	Unmarshal(raw []byte)
	Execute() bool
}

type Task struct {
	Name string                   `yaml:"name"`
	Loop []string                 `yaml:"loop"`
	Tags []string                 `yaml:"tags"`
	Kind []map[string]interface{} `yaml:"-"`
	//File        File        `yaml:"file"`
	//Synchronize Synchronize `yaml:"synchronize"`
}

func LoadTasks(name string) {
	var path = constants.ResourceDir + "roles/" + name + "/tasks/main.yaml"
	var tasks []Task

	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(dat), &tasks)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(dat), &tasks[0].Kind)
	if err != nil {
		panic(err)
	}

	//fmt.Println("len", len(tasks[0].Kind))
	//fmt.Println("len of kind maps", len(tasks[0].Kind[0]))

	for _, task := range tasks {
		for _, kind := range task.Kind {
			for k, v := range kind {
				fmt.Println("v", v)

				switch k {
				case "name":
					fmt.Println("name found")
				case "loop":
					fmt.Println("loop found")
				case "tags":
					fmt.Println("tags found")
				case FILE, SYNCHRONIZE:
					raw, err := yaml.Marshal(v)
					if err != nil {
						fmt.Println("err while Marshal")
					}
					handleTasks(k, raw)
				default:
					fmt.Println("======================================")
					fmt.Println("v=", v)
					fmt.Println("v= type is", reflect.TypeOf(v))
					switch v.(type) {
					case File:
						fmt.Println("file")
					case []File:
						fmt.Println("file array")
					case map[string]interface{}:
						fmt.Println("map[string]interface{}")
					case map[interface{}]interface{}:
						fmt.Println("found map[interface{}]interface{}")
					default:
						fmt.Println("unable to find type")
					}
					fmt.Println("======================================")
				}
			}
			fmt.Println("**********Attention************")
			fmt.Println("One task is done")
			fmt.Println("**********Attention************")
		}
	}
}

func handleTasks(task string, raw []byte) {
	switch task {
	case FILE:
		File{}.Unmarshal(raw)
	case SYNCHRONIZE:
		Synchronize{}.Unmarshal(raw)
	default:
		fmt.Println("Are you lost sweetheart?")
	}
}
