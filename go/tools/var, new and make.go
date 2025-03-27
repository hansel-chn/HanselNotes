package main

import (
	"fmt"
)

func main() {
	s1 := new(*int)
	//**s1 = 10
	fmt.Println(*s1)
	fmt.Println(s1)

	var s2 *int
	//*s2 = 10
	fmt.Println(s2)
	fmt.Println(&s2)
	//fmt.Println(*s2)

	a := new([]int)
	b := make([]int, 0)

	fmt.Printf("%p\n", *a)
	fmt.Printf("%p\n", b)

	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", &b)

	fmt.Println(*a)
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(&b)

	fmt.Printf("%#v\n", a)
}
