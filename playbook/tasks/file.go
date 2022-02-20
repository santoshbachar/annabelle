package tasks

import (
	"gopkg.in/yaml.v2"
)

type File struct {
	Path  string `yaml:"path"`
	State string `yaml:"state"`
	//Recurse Boolean `yaml:"recurse"`
	Recurse string `yaml:"recurse"`
	Owner   string `yaml:"owner"`
	Group   string `yaml:"group"`
	Mode    string `yaml:"mode"`
}

func (f File) Unmarshal(file []byte) (bool, error) {
	err := yaml.Unmarshal([]byte(file), &f)
	if err != nil {
		return false, err
	}
	return true, nil
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
