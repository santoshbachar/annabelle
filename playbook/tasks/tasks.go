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
	Unmarshal(file []byte) (bool, error)
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
					if k == FILE {
						u, err := yaml.Marshal(v)
						if err != nil {
							fmt.Println("err while Marshal")
						}
						f := File{}
						ok, err := f.Unmarshal(u)
						if ok {
							fmt.Println("Interface function works")
						}
						//fok, err := File{}.Unmarshal(u)

						//HandleFile(u)

						//fi := v.(map[interface{}]interface{})
						//f := File{}
						//if name, o := fi["group"].(string); o {
						//	f.Group = name
						//}
						//fmt.Println("group", name)

						file, ok := v.(File)
						if !ok {
							fmt.Println("not ok")
						} else {
							fmt.Println("printing file", file)
						}
					} else if k == SYNCHRONIZE {
						u, err := yaml.Marshal(v)
						if err != nil {
							fmt.Println("err while Marshal")
						}
						HandleSynchronize(u)
					}

					fmt.Println("======================================")
					typeofstruct(v)
					teststruct(v)
					fmt.Println("v=", v)
					fmt.Println("v= type is", reflect.TypeOf(v))

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
