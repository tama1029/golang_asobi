package interface_sample

import "fmt"

type Call interface {
	reply() string
}

type Dog struct {}

type Cat struct {}

func (d Dog) reply() string {
	return "wanwan"
}

func (c Cat) reply() string {
	return "nyaa"
}

func Run() {
	d := Dog{}
	fmt.Println(d.reply())

	c := Cat{}
	fmt.Println(c.reply())
}
