package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

func main() {
	//ticker := time.NewTicker(time.Second)
	//defer ticker.Stop()
	//done := make(chan bool)
	//go func() {
	//	time.Sleep(10 * time.Second)
	//	done <- true
	//}()
	//for {
	//	select {
	//	case <-done:
	//		fmt.Println("Done!")
	//		return
	//	case t := <-ticker.C:
	//		fmt.Println("Current time: ", t)
	//	}
	//}
	//type a struct {
	//	id   int
	//	DsFf string
	//}
	//s := a{
	//	id:   2,
	//	DsFf: "1111",
	//}
	////b := reflect.TypeOf(s)
	////c := reflect.ValueOf(s)
	////fmt.Println(b.Field(1).Name)
	////fmt.Println(c.Field(1).Kind())
	////fmt.Printf("%v\n", b)
	////fmt.Println(c)
	//type test struct {
	//	name   string
	//	policy []interface{}
	//}
	//data := []a{s, s, s, s, s, s, s}
	//dataTest := func(data interface{}) test {
	//	return test{
	//		name:   "test",
	//		policy: data,
	//	}
	//}(data)
	//fmt.Println(reflect.TypeOf(dataTest.policy))
	//fmt.Println(dataTest.policy)
	//dataInter, ok := dataTest.policy.([]interface{})
	//fmt.Println(ok)
	//fmt.Println(reflect.TypeOf(dataInter))
	//fmt.Println(dataInter)
	a := make([]int, 10, 10)
	a[1] = 1
	b := a
	//b := make([]int, len(a))
	//copy(b, a)
	fmt.Println(a)
	fmt.Println(b)
	a[2] = 1
	fmt.Println(a)
	fmt.Println(b)

	c := make([]int, 0)
	c = append(c, 1)
	d := c
	fmt.Println(c)
	fmt.Println(d)
	c = append(c, 1)
	c[0] = 2
	fmt.Println(c)
	fmt.Println(d)

	//c := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	//fmt.Println(len(c[:len(c)-1]))
	//d := c[:len(c)-1]
	//fmt.Println(len(c))
	//fmt.Println(len(d))

	fmt.Println("二维切片比较")
	q := make([][]int, 0)
	w := make([][]int, 0)
	e := make([]int, 0)
	for i := 0; i < 10; i++ {
		e = append(e, i)
		q = append(q, e)
		w = append(w, e)
	}
	fmt.Println(q)
	fmt.Println(w)
	fmt.Println(e)

	r := [][]int{{1, 2}}
	t := [][]int{{1, 2}}
	fmt.Printf("%p\n", r)
	fmt.Printf("%p\n", t)
	fmt.Printf("%p\n", r[0])
	fmt.Printf("%p\n", t[0])

	fmt.Println("数组边界")
	fmt.Println(1 / 3)

	fmt.Println("int长度")
	fmt.Println(strconv.IntSize)
	fmt.Println(runtime.GOARCH)

	var maxInt int = 1
	maxInt = maxInt << 62
	fmt.Println(maxInt)
	maxInt = maxInt << 1
	fmt.Println(maxInt)
	maxInt--
	fmt.Println(maxInt)
	maxInt = maxInt << 1
	fmt.Println(maxInt)

	fmt.Println("字符串相减")
	fmt.Println('3' - '2')
	fmt.Println([]byte("3")[0] - []byte("2")[0])
	fmt.Println()

	fmt.Println("遍历字符串得到的是")
	fmt.Println("aaaa")
	strings := "aaaaaaa"
	for i1, i2 := range strings {
		fmt.Println(reflect.TypeOf(strings[i1]))
		fmt.Println(reflect.TypeOf(i2))
		fmt.Println(reflect.ValueOf(strings[i1]))
		fmt.Println(reflect.ValueOf(i2))
		break
	}

	var stringTest string = "aaaaaaaaa"
	//var runeTest rune = 'a'
	//var byteTest uint8 = 'a'
	//fmt.Println(runeTest - byteTest)
	var runeTest int32 = '我'
	var byteTest uint8 = 'a'
	//rune(byteTest)
	//byte(runeTest)
	fmt.Printf("%b\n", byteTest)
	fmt.Printf("%b\n", runeTest)
	fmt.Printf("%b\n", uint8(runeTest)) // 做了个截断
	fmt.Println(uint8(runeTest) - byteTest)
	//_ = stringTest[0] - '我'
	fmt.Println(reflect.TypeOf(stringTest[0]))
	fmt.Println(reflect.TypeOf('0'))

	fmt.Println("byte convert")
	fmt.Println(byte(10))
	fmt.Println(reflect.TypeOf(10))
	fmt.Println(reflect.TypeOf(byte(10)))

	// go不像python可以直接这样扩充
	fmt.Println('3' * 3)
	fmt.Println('3')
	fmt.Println("a" + "b")

	// for循环里面不固定每次重新计算len(string2)-len(string1)
	string1 := "aa"
	string2 := "aaaaaa"
	for i := 0; i < len(string2)-len(string1); i++ {
		string1 = "0" + string1
	}
	fmt.Println(string1) //00aa
	fmt.Println(string2) //aaaaaa

	// for 问题
	v := []int{1, 2, 3}
	for _, data := range v {
		//v = append(v, i)
		v[1] = 100
		fmt.Println(data)
	}

	//maxPosition := 5 //无限循环
	//for i := 1; i <= maxPosition; i++ {
	//	maxPosition++
	//	fmt.Println("========")
	//}

	h := man{}
	h.speak()
	h.sing()
	//var j human = man{}
	//j.speak()
	//j.sing()

	fmt.Println("================")
	listnode := &ListNode{}
	listnode.Next = &ListNode{}
	fmt.Println(&listnode)
	fmt.Println(listnode)
	listnode = &ListNode{} // 会重新分配地址
	listnode.Next = &ListNode{}
	fmt.Println(&listnode)
	fmt.Println(listnode)

	fmt.Println("================")
	listnode1 := new(ListNode)
	listnode1.Next = new(ListNode)
	fmt.Println(&listnode1)
	fmt.Println(listnode1)
	listnode1 = new(ListNode)
	listnode1.Next = new(ListNode)
	fmt.Println(&listnode1)
	fmt.Println(listnode1)

	fmt.Println("================")
	var listnode2 *ListNode
	listnode2.Next = new(ListNode)
	fmt.Println(&listnode2)
	fmt.Println(listnode2)
	//var listnode2 *ListNode
	listnode2.Next = new(ListNode)
	fmt.Println(&listnode2)
	fmt.Println(listnode2)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type human interface {
	speak()
	sing()
}

type man struct {
	name string
}

func (m man) speak() {
	fmt.Println("speaking")
}

func (m *man) sing() {
	fmt.Println("singing")
}
