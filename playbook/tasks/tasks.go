package tasks

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/santoshbachar/annabelle/playbook/constants"
	"gopkg.in/yaml.v2"
	"os"
	"sync"
)

//type TaskType string

const (
	NAME        = "name"
	LOOP        = "loop"
	TAGS        = "tags"
	FILE        = "file"
	SYNCHRONIZE = "synchronize"
	GROUP       = "group"
	USER        = "user"
	SCRIPT      = "script"
)

type HandleTask interface {
	Unmarshal(raw []byte)
	Execute(loop Loop) bool
	//AddLoopItems(items interface{})
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

	var wg sync.WaitGroup
	var handle HandleTask
	for _, task := range tasks {
		for _, kind := range task.Kind {
			loop := Loop{}
			for k, v := range kind {
				//fmt.Println("value =>", v)
				switch k {
				case NAME:
					fmt.Println("name value =>", v)
				case LOOP:
					fmt.Println("loop value =>", v)
					//handle.AddLoopItems(v)
					loop.UnmarshallLoopItems(v)
				case TAGS:
					fmt.Println("tags found")
				case FILE, SYNCHRONIZE, GROUP, USER:
					raw, err := yaml.Marshal(v)
					if err != nil {
						fmt.Println("err while Marshal")
					}
					wg.Add(1)
					//go func() {
					handle = unMarshallTasks(k, raw)
				//}()
				case SCRIPT:
					wg.Add(1)
					handle = unMarshallSpecialTasks(name, v.(string))
					fmt.Println("This is a script with value", v)
				default:
					fmt.Println("Are you lost sweetheart?")
				}
			}
			fmt.Println("**********Attention************")
			color.Green("One task is done")
			fmt.Println("**********Attention************")
			//for i := 0
			//i < len(loo))
			handle.Execute(loop)
			wg.Done()
		}
	}
}

func unMarshallTasks(task string, raw []byte) HandleTask {
	switch task {
	case FILE:
		file := File{}
		file.Unmarshal(raw)
		return &file
		//File{}.Unmarshal(raw)
	case SYNCHRONIZE:
		sync := Synchronize{}
		sync.Unmarshal(raw)
		return &sync
	//return Synchronize{}.Unmarshal(raw)
	case GROUP:
		group := Group{}
		group.Unmarshal(raw)
		return &group
	case USER:
		user := User{}
		user.Unmarshal(raw)
		return &user
	default:
		fmt.Println("someone call the cops")
	}
	return nil
}

func unMarshallSpecialTasks(roleName, fileName string) HandleTask {
	script := Script{}
	script.init(roleName, fileName)
	return &script
}
