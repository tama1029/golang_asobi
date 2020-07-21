package interface_sample

import "fmt"

type Call interface {
	Reply() string
}

type Dog struct {}

type Cat struct {}

func (d Dog) Reply() string {
	return "wanwan"
}

func (c Cat) Reply() string {
	return "nyaa"
}

func Run() {
	d := Dog{}
	fmt.Println(d.Reply())

	c := Cat{}
	fmt.Println(c.Reply())
}
