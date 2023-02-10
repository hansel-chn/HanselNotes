package main

import "fmt"

func main() {
	//s := "cd"
	//d := "c"
	//f := "cb"
	a := make([]int, 0)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	fmt.Println(len(a))
	a = a[:len(a)-1]
	fmt.Println(len(a))

	var array [2]int
	array[0] = 1
	array[0] = 2
	fmt.Println("array")
	fmt.Println(array)

	var ArrayInSlice [][2]int
	ArrayInSlice = append(ArrayInSlice, [2]int{1, 2})
	ArrayInSlice[0][0] = 111
	ArrayInSlice[0][0] = 222
	fmt.Println("ArrayInSlice")
	fmt.Println(ArrayInSlice)

	ArrayInMap := make(map[int][2]int)
	ArrayInMap[1] = [2]int{1, 2}
	// ArrayInMap[1][0]不可被寻址
	//map通过hash表来实现，可能出现扩容，地址会变
	// map映射值本质上是不可变的，因为它们是不可寻址的;因此，您不能只编辑其中的一部分（array[1]用地址赋值，切片不同底层存的是指针）。你必须复制出数组，修改它，然后赋值回去。
	//https://stackoverflow.com/questions/32751537/why-do-i-get-a-cannot-assign-error-when-setting-value-to-a-struct-as-a-value-i
	//http://www.wu.run/2021/11/12/not-addressable-in-golang/
	// 但是存在一个问题，为什么不能类似可寻址的值类型调用指针接收者方法一样，内部做一个处理 原因：Everything in Go returns a copy,
	//maps are no different. v is a copy of the contents of the value at m[k].是个值拷贝不能赋值
	//但是还有个问题，为什么不能在内部做，

	//ArrayInMap[1][0] = 111
	//ArrayInMap[1][0] = 222
	fmt.Println("ArrayInMap")
	fmt.Println(ArrayInMap[1])

	//SliceInMap := make(map[int][2]int) 数组不可以
	SliceInMap := make(map[int][]int)
	SliceInMap[1] = []int{1, 2}
	SliceInMap[1][0] = 111
	fmt.Println("SliceInMap")
	fmt.Println(SliceInMap[1])
	fmt.Println(len(SliceInMap[1]))
	fmt.Println(cap(SliceInMap[1]))
	fmt.Printf("%p\n", SliceInMap[1])
	//fmt.Printf("%p\n",&SliceInMap[1]) // 不能被寻址
	//当切片扩容
	for i := 0; i < 5; i++ {
		SliceInMap[1] = append(SliceInMap[1], i)
		fmt.Println("切片扩容", i)
		fmt.Println(SliceInMap[1])
		fmt.Println(len(SliceInMap[1]))
		fmt.Println(cap(SliceInMap[1]))
		fmt.Printf("%p\n", SliceInMap[1])
	}
	SliceInMap1 := make(map[int][2]int)
	SliceInMap1[1] = [2]int{1, 2}
	aqqqqq := SliceInMap1[1]
	aqqqqq[0] = 111
}
