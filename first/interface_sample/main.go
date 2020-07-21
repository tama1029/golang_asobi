package interface_sample

import "fmt"

type Call interface {
	reply() string
}

type dog struct {}

type cat struct {}

func (d dog) reply() string {
	return "wanwan"
}

func (c cat) reply() string {
	return "nyaa"
}

func NewDog () Call {
	return dog{}
}

func NewCat () Call {
	return cat{}
}

func Run() {
	d := NewDog()
	fmt.Println(d.reply())

	c := NewCat()
	fmt.Println(c.reply())
}
