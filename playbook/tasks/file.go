package tasks

import (
	"fmt"
	"github.com/santoshbachar/annabelle/playbook/helper"
	"gopkg.in/yaml.v2"
)

type fileState string

const (
	Directory fileState = "directory"
	_File               = "file"
)

type File struct {
	Path    string    `yaml:"path"`
	State   fileState `yaml:"state"`
	Recurse Boolean   `yaml:"recurse"`
	//Recurse string `yaml:"recurse"`
	Owner string `yaml:"owner"`
	Group string `yaml:"group"`
	Mode  string `yaml:"mode"`
}

func (f *File) Unmarshal(file []byte) {
	err := yaml.Unmarshal([]byte(file), &f)
	if err != nil {
		panic(err)
	}
}

func (f File) Execute(loop Loop) bool {
	fmt.Println("file execute()")
	if f.Path == "" {
		fmt.Println("Where to act?")
		return false
	}

	if f.State == "" {
		fmt.Println("State absent")
		return false
	}

	fmt.Println("File execute(). All well. Executing.")

	switch f.State {
	case Directory:
		fmt.Println("Directory needs to be created at", f.Path)
		makeDirectory(f.Path, f.Recurse)
	case _File:
		fmt.Println("File needs to be created at", f.Path)
	default:
		fmt.Println("unsupported state", f.State)
	}

	//if f.nextItem() != nil {
	//}

	return true
}

func makeDirectory(path string, recurse Boolean) {
	fmt.Println("makeDirectory()")
	switch recurse {
	case Yes:
		fmt.Println("creating dir with recursive pattern")
		helper.FindVariable(path)
	case No:
		fmt.Println("creating dir with non-recursive pattern")
	default:
		fmt.Println("Some unknown recursive value is found, defaulting to NO.")
	}
}

//unsafe
//func (f File) AddLoopItems(items interface{}) {
//	loop := Loop{}
//	loop.UnmarshallLoopItems(items)
//}

func (f File) nextItem() {

}

//func HandleFile(file []byte) {
//	//file := File{x}
//	fmt.Println("HandleFile*********")
//	fs := File{}
//	err := yaml.Unmarshal([]byte(file), &fs)
//	if err != nil {
//		fmt.Println("err while Unmarshall")
//	}
//	fmt.Println(fs)
//}
