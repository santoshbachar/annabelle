package tasks

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

type Loop struct {
	items   []string
	counter int
}

//var items []string

func (l Loop) UnmarshallLoopItems(items interface{}) {
	raw, err := yaml.Marshal(items)
	if err != nil {
		fmt.Println("err while Marshal")
	}

	err = yaml.Unmarshal([]byte(raw), &l.items)
	if err != nil {
		panic(err)
	}

	fmt.Println("Loop items =>", l.items)
}

func (l Loop) nextItem() string {
	return l.items[l.counter]
}
