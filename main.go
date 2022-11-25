package main

import (
	"fmt"
)

func main() {
	type parent struct {
		note string
	}

	parentInstance := &parent{note: "1111"}

	type child struct {
		*parent
		id   int
		name string
	}
	a := new(child)
	a.parent = parentInstance
	a.id = 1

	fmt.Printf("%p\n", parentInstance)
	fmt.Println(*a)

	//json.Decoder{}
}
