package main

import (
	"fmt"
	"io"
	"os"
)

//func main() {
//	fn := adder()
//	fmt.Println(&fn)
//
//	//fmt.Println(fn()) //11
//	//fmt.Println(fn()) //11
//	//
//	//fn1 := adder1()
//	//fmt.Println(fn1(10)) //20
//	//fmt.Println(fn1(10)) //30
//}
//
//func adder() func() int {
//	i := 10
//	return func() int {
//		return i + 1
//	}
//}
//
//func adder1() func(x int) int {
//	i := 10
//	return func(x int) int {
//		i = i + x
//		return i
//	}
//}

const (
	NAMELIST = 1 + iota
	PACKAGESTATUSDELETED
	PACKAGIST
)

func main() {
	var c []byte
	//a := new([]byte)
	//b := make([]byte, 5)
	//fmt.Printf("%p------%p\n", a, &a)
	//fmt.Printf("%v------%v\n", a, &a)
	//fmt.Printf("%p------%p\n", b, &b)
	//fmt.Printf("%v------%v\n", b, &b)
	//fmt.Printf("%p------%p\n", c, &c)
	//fmt.Printf("%v------%v\n", c, &c)
	//
	//fmt.Printf("%p", *a)

	file, err := os.Open("./aaa.txt")
	if err != nil {
		return
	}
	//n, err := file.Read(b)
	c, err = io.ReadAll(file)
	file.Close()
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	//fmt.Println(n)
	//fmt.Println(b)
	fmt.Println(c)
	fmt.Println("=============")
}
