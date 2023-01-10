# Collection

## go slice

go中`slice`存放的是一个地址。s储存在栈区，s在栈中的地址是&s，存放的为数组`[1,2,3]`在堆区的地址s

	```
    var s = []int{1, 2, 3}
	fmt.Println("s存放的地址和s的地址")
	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", &s)
	a := s
	fmt.Println("a存放的地址和s的地址")
	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", &a)
    ```

# byte buffer and strings builder

[byte buffer vs strings builder](https://geektutu.com/post/hpg-string-concat.html)

如果我们自己来实现strings.Builder,
大部分情况下我们完成前3步就觉得大功告成了。但是标准库做得要更近一步。我们知道Golang的堆栈在大部分情况下是不需要开发者关注的，如果能够在栈上完成的工作逃逸到了堆上，性能就大打折扣了。因此，copyCheck
加入了一行比较hack的代码来避免buf逃逸到堆上。关于这部分内容，你可以进一步阅读Dave
Cheney的关于[Go’s hidden #pragmas](https://dave.cheney.net/2018/01/08/gos-hidden-pragmas).

# go结构体

```text
package main

import "fmt"

func main() {
	type test struct {
		id   int
		id2  int32
		name string // 16个字节
		id3  int
		id4  int8
	}
	//使用new(T)函数为一个结构体分配内存，返回该结构体的指针。
	var a = test{}
	//结构体放的顺序不同，内存消耗不同，为了能让CPU减少一次获取的时间，Go编译器会帮你把struct结构体做数据的对齐.cpu64位一次读8个字节，
	//https://www.jb51.net/article/265457.htm
	fmt.Printf("%p\n", &(a.id))   // int，8字节。
	fmt.Printf("%p\n", &(a.id2))  // int32，4字节。但是打印地址还是差8字节，因为内存对齐，浪费了内存空间。
	fmt.Printf("%p\n", &(a.name)) // string，16字节（地址八位，长度8位）。差八个字节，内存对齐，不是自己以为的地址8字节。这里是连续的，存的就是值。
	fmt.Printf("%p\n", &(a.id3))  // 差16个字节
	fmt.Printf("%p\n", &(a.id4))  // 差8个字节
	//fmt.Printf("%p\n", a)         // 结构体的地址和第一个字段的地址一样
	//fmt.Printf("%p", &a)
	//0xc00007e4b0
	//0xc00007e4b8
	//0xc00007e4c0
	//0xc00007e4d0
	//0xc00007e4d8

	/*这样实例化和 new一样，且能取到具体地址，只定义地址未初始化0x0*/
	//var b = &test{}
	//fmt.Printf("%p\n", &(b.id))
	//fmt.Printf("%p\n", &(b.id2))
	//fmt.Printf("%p\n", &(b.name)) //差八个字节。
	////fmt.Printf("%p\n", b)         // 结构体的地址和第一个字段的地址一样
	//fmt.Printf("%p", &b)

}
```

## json.unmarshal and json decoder

* 第一种 根据输入来定，如果是一个流，选用decoder，如果是一个字符串选用unmarshal
*

第二种说法：我建议不要使用json.Decoder。它用于JSON对象流，而不是单个对象。对于单个JSON对象，它的效率并不高，因为它将整个对象读入内存。它的缺点是，如果在对象后面包含了垃圾，它就不会报错。取决于几个因素，json。解码器可能无法完全读取正文，连接将不符合重用的条件。

> It really depends on what your input is. If you look at the implementation of the Decode method of json.Decoder, <mark>**it buffers the entire JSON value in memory before unmarshalling it into a Go value. So in most cases it won't be any more memory efficient**</mark> (although this could easily change in a future version of the language).
>
>So a better rule of thumb is this:
>
>Use json.Decoder if your data is coming from an io.Reader stream, or you need to decode multiple values from a stream of data. Use json.Unmarshal if you already have the JSON data in memory. For the case of reading from an HTTP request, I'd pick json.Decoder since you're obviously reading from a stream.
>
> <mark>高亮话指的是和unmarshal相比（从数据流读到内存，再unmarshal）相比，decoder将io.Reader(res.body)整个值又额外缓存了一次，先缓存到decoder里的io.reader中,再进行（从数据流读到内存，再unmarshal）/mark>

## golang结构体

[golang空结构体不占据内存空间](https://zhuanlan.zhihu.com/p/420542865)

[golang空结构体](https://docker.blog.csdn.net/article/details/106917358?spm=1001.2101.3001.6650.2&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-2-106917358-blog-117222307.pc_relevant_3mothn_strategy_recovery&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-2-106917358-blog-117222307.pc_relevant_3mothn_strategy_recovery&utm_relevant_index=3)

[WaitGroup优雅退出](https://blog.csdn.net/inthat/article/details/107000010?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-1-107000010-blog-124380196.pc_relevant_multi_platform_whitelistv3&spm=1001.2101.3001.4242.2&utm_relevant_index=4)

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int
	var s struct{}
	{
	}
	fmt.Println(s)
	//fmt.Println(s == nil)
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(i))
	fmt.Printf("%p\n", &s)
	fmt.Printf("%p", &i)
}
```

## 切片存的是什么

[切片存的是什么](https://blog.csdn.net/qq_24599471/article/details/106000190?ops_request_misc=&request_id=&biz_id=102&utm_term=golang%E5%88%87%E7%89%87%E5%AD%98%E7%9A%84%E4%BB%80%E4%B9%88&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduweb~default-3-106000190.142^v67^control,201^v3^add_ask,213^v2^t3_control1&spm=1018.2226.3001.4187)

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a [8]int
	var b []int
	b = []int{1, 2, 3, 4, 4, 4, 4}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Printf("%p\n", a)
	fmt.Println(unsafe.Sizeof(b))
	fmt.Printf("%p\n", b)
}
```

## nil

```go
package main

import "fmt"

func main() {
	var i int = 0
	var j int
	var name []int
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", j)
	fmt.Printf("%p\n", &i)
	fmt.Printf("%p\n", &j)
	fmt.Printf("%p\n", name)
	fmt.Println(name == nil)
	fmt.Printf("%p\n", &name)
}
```

[nil](https://blog.csdn.net/zf766045962/article/details/81005037?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522166980205016800213039738%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=166980205016800213039738&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~baidu_landing_v2~default-3-81005037-null-null.142^v67^control,201^v3^add_ask,213^v2^t3_control1&utm_term=golang%20nil&spm=1018.2226.3001.4187)

## go错误处理

如何解决嵌套错误 github.com/pkg/errors errors.wrap

## unsafe.pointer uintptr

- A pointer value of any type can be converted to a Pointer.
- A Pointer can be converted to a pointer value of any type.
- A uintptr can be converted to a Pointer.
- A Pointer can be converted to a uintptr 所以对指针操作需先转换为uintptr。

## golang string rune byte

string实际上储存的是16位，八位地址八位长度

底层存的

- string len 结果是字节数，但是一个字符可能表达成1-3个字节
- rune 4个字节
- byte 1个字节

## 钩子函数

[钩子函数](https://blog.csdn.net/BF02jgtRS00XKtCx/article/details/110458293?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-110458293-blog-89481066.pc_relevant_3mothn_strategy_recovery&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-110458293-blog-89481066.pc_relevant_3mothn_strategy_recovery&utm_relevant_index=2)

钩子函数和回调函数 个人理解就是见链接的python代码，没什么区别，可以自己控制在事件发生前后回调函数

## mask and sweep 和三色标记法

[综合来看讲的最好的](https://golang.design/under-the-hood/zh-cn/part2runtime/ch08gc/barrier/)
强三色不变性和弱三色不变性

## var, new and make

[var, new and make](./var,%20new%20and%20make.go)

## Golang 深浅拷贝

golang中所有拷贝均为值拷贝，出现类似引用拷贝是因为如slice底层是struct的原因。

* 注意golang中的copy()也不是深拷贝，copy()拷贝的是切片中的元素。对于多维切片来说，比如二维`a :=[][]int`，`copy(b,a)`拷贝的是a中的元素，及a中切片的地址。a和b地址不同，但是a[0]和b[0]
  地址相同。

[Golang 深浅拷贝](https://allendaydayup.blog.csdn.net/article/details/124913809?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-124913809-blog-127450453.pc_relevant_recovery_v2&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-124913809-blog-127450453.pc_relevant_recovery_v2&utm_relevant_index=1)

## golang 遍历字符串(下标遍历，range遍历)

* 下标遍历得到的是`uint8`类型(对应`byte`类型)的数据，对`[]byte`,`string`，使用`len`获得的是字符串字节的长度
* 相对的，range遍历得到的是int32类型的数据(对应`rune`类型的数据)，对`[]rune`使用`len`获得的是获得的是字符的长度

## golang for

for循环条件，每次都重新计算循环条件

```
// for循环里面不固定每次重新计算len(string2)-len(string1)
string1 := "aa"
string2 := "aaaaaa"
for i := 0; i < len(string2)-len(string1); i++ {
	string1 = "0" + string1
}
fmt.Println(string1) //00aa
fmt.Println(string2) //aaaaaa
```

但是使用range 存在语法糖,如下代码不是死循环
`for_temp := range`和`len_temp := len(for_temp)`把长度和切片记录下来，所以：
1. 不会死循环
2. 打印`v[1]`为改动过后的值

```
func main() {
	v := []int{1, 2, 3}
	for _, data := range v {
		v = append(v, i) // 不是死循环
		v[1] = 100
		fmt.Println(data) // 打印的三个值为1，100，3
	}
}
```