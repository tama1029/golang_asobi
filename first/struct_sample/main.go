package struct_sample

import "fmt"

type Dog struct {
	Name string
	Age int
}

func Run() {
	d := Dog{Name: "poti", Age: 5}
	fmt.Println(fmt.Sprintf("Dog name is %s, and he is %d years old.", d.Name, d.Age))
}
