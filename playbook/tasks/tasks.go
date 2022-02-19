package tasks

import (
	"fmt"
	"github.com/santoshbachar/annabelle/playbook/constants"
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
)

type TaskType string

const (
	FILE        TaskType = "file"
	SYNCHRONIZE          = "synchronize"
)

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

	//fmt.Println("entire task", tasks[0])
	//fmt.Println("task Kind: ", tasks[0].Kind)
	//fmt.Println("task file: ", tasks[0].File)

	//fmt.Println("task name: ", tasks[0].Name)
	//fmt.Println("task tags: ", tasks[0].Tags)
	//fmt.Println("task loop: ", tasks[0].Loop)

	//fmt.Println("try: ", Keys(tasks[0].Kind))

	fmt.Println("len", len(tasks[0].Kind))
	fmt.Println("len of kind maps", len(tasks[0].Kind[0]))

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
				default:
					if k == "file" {
						fi := v.(map[interface{}]interface{})
						f := File{}
						if name, o := fi["group"].(string); o {
							f.Group = name
						}
						fmt.Println("group", name)
						file, ok := v.(File)
						if !ok {
							fmt.Println("not ok")
						} else {
							fmt.Println("printing file", file)
						}
					}

					fmt.Println("======================================")
					typeofstruct(v)
					teststruct(v)
					fmt.Println(v)
					switch v.(type) {
					case string:
						fmt.Println("String found")
					case []string:
						fmt.Println("array of string")
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
		}
	}
}

func typeofstruct(x interface{}) {
	m := make(map[string]string)
	m["a"] = "b"
	fmt.Println(reflect.TypeOf(m))
	fmt.Println("typeofstruct() -> ",
		reflect.ValueOf(x).Kind())
	fmt.Println("typeofstruct() -> ",
		reflect.TypeOf(x))
}

func teststruct(x interface{}) {
	fmt.Println("teststruct func()======")
	switch x.(type) {
	case File:
		fmt.Println("I am file ")

	}
}
