package main

import (
	"fmt"
	"github.com/tama1029/golang_asobi/first/func_sample"
	"github.com/tama1029/golang_asobi/first/interface_sample"
	"github.com/tama1029/golang_asobi/first/struct_sample"
)

func main() {
	fmt.Println("Hello Golang!")

	fmt.Println("---------- func sample ------------")
	func_sample.Run()
	fmt.Println("---------- struct sample ------------")
	struct_sample.Run()
	fmt.Println("---------- interface sample -----------")
	interface_sample.Run()

}
