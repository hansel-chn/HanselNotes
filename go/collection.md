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

[golang GC](https://www.cnblogs.com/flippedxyy/p/15558742.html)

[这里图片错误，Yuasa写屏障图片画错，正确应该是将删除的节点上色，图片画成了将指向删除节点的节点上色](https://golang.design/under-the-hood/zh-cn/part2runtime/ch08gc/barrier/)

强三色不变性和弱三色不变性

插入写屏障”机制,对于栈中的对象是不生效的，“插入写屏障” 仅仅使用在堆中生效。所以在结束时需要STW来重新扫描栈，执行三色标记法回收白色垃圾

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

## Goroutine栈空间无限大原因

[http://dave.cheney.net/2013/06/02/why-is-a-goroutines-stack-infinite](http://dave.cheney.net/2013/06/02/why-is-a-goroutines-stack-infinite)

[https://blog.csdn.net/weixin_52690231/article/details/124476954](https://blog.csdn.net/weixin_52690231/article/details/124476954)

[swap交换内存主要是指当物理内存不够用时，系统会启用硬盘的一部分空间来充当服务器内存，而默认情况下swap内存会有一些设置标准，它与物理内存的大小也是有关系的](https://blog.csdn.net/Listen2You/article/details/108205275?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522167343982316800186555927%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fall.%2522%257D&request_id=167343982316800186555927&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~first_rank_ecpm_v1~rank_v31_ecpm-2-108205275-null-null.142^v70^control,201^v4^add_ask&utm_term=%E5%A0%86%E5%86%85%E5%AD%98%E6%89%A9%E5%AE%B9%E5%90%97&spm=1018.2226.3001.4187)

这里好像还讲的是分段栈 when new stack pages are needed, they are allocated from the heap.

如果物理内存不足了，数据会在主存和磁盘之间频繁交换，命中率很低，性能出现急剧下降，我们称这种现象叫内存颠簸。这时你会发现系统的 swap 空间利用率开始增高， CPU 利用率中 iowait 占比开始增高。

## Golang 栈--分段栈和连续栈

[分段栈和连续栈](https://www.jianshu.com/p/7ec9acca6480)

Go 1.4 开始使用的是连续栈，而这之前使用的分段栈。 分段栈：分段栈是指开始时只有一个stack，当需要更多的 stack 时，就再去申请一个，然后将多个stack 之间用双向链接连接在一起。当使用完成后，再将无用的 stack
从链接中删除释放内存。 连续栈：创建一个两倍于原stack大小的新stack，并将旧栈拷贝到其中

## 设计模式

单例模式，工厂模式

## 指针接收者

一个问题： 值类型可被寻址可以调用指针接收者的方法，编译器会做`(&fv).pointerMethod()`的操作. 那么问题来了，值类型不能寻址，比如函数返回值，其不能调用指针接收者方法，

当指针接收者实现接口时，值类型不可寻址自然不能实现接口。但是，为什么<font color=LightCoral>即使golang值类型可被寻址也不能实现指针类型的接口</font>，

```
h := man{} // grammatically correct
h.speak()
h.sing()
var j human = man{} // grammatical mistake

j.speak()
j.sing()

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
```

[指针接收者](https://stackoverflow.com/questions/38166925/why-cant-i-assign-types-value-to-an-interface-implementing-methods-with-receiv)

## How does GoLand of JetBrains find the implementations of interface?

## Dependency Inversion Principle, Dependency Injection (DI), Inversion of Control (IoC)

[stackoverflow](https://stackoverflow.com/questions/6550700/inversion-of-control-vs-dependency-injection)

![IOC,DI,DIP关系](./ioc.png "IOC,DI,DIP关系")

* Inversion of Control 思考java的ioc容器，对象的控制权上交ioc容器。原先A依赖B需要主动创建B,现在通过IOC容器将对象B注入对象A
* Dependency Injection (DI)  依赖注入(DI)模式是IoC模式的一个更具体的版本，实现通过constructors/setters/Interface ，对象将“依赖”这些以正确地行为。
* Dependency Inversion Principle 高级模块不应该依赖于低级模块。两者都应该依赖于抽象。抽象不应该依赖于细节。细节应该依赖于抽象。

## 逃逸分析

## go 并发-锁与channel的选择

[go 并发-锁与channel的选择](https://lessisbetter.site/2019/01/14/golang-channel-and-mutex/)

channel底层牵扯到了锁

* channel的核心是关注的是数据的流动 传递数据的所有，即把某个数据发送给其他协程,分发任务，每个任务都是一个数据
  交流异步结果，结果是一个数据
* mutex lock核心是关注公共的数据,mutex的能力是数据不动，某段时间只给一个协程访问数据的权限擅长数据位置固定的场景
```go
package main

import "time"
import "fmt"

func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		time.Sleep(5 * time.Second)
		for _, n := range nums {
			time.Sleep(time.Second)
			out <- n
		}
	}()
	return out
}
func main() {
	ch := producer(1, 2, 3, 4, 5, 6, 7, 8, 910)
	for i2 := range ch {
		fmt.Println(i2)
	}
	a := <-ch
	fmt.Println(a)
}
```

[通信共享内存的含义：](https://medium.com/@buzoumei/%E4%B8%BA%E4%BB%80%E4%B9%88%E4%BD%BF%E7%94%A8%E9%80%9A%E4%BF%A1%E6%9D%A5%E5%85%B1%E4%BA%AB%E5%86%85%E5%AD%98-do-not-communicate-by-sharing-memory-instead-share-memory-by-communicating-5827bd3c4a77)
通信共享内存的含义：
不论通信模式，线程、协程最终都是从内存中获取数据，所以本质上都是通过共享内存来通信。

“通信共享内存”，应该说是使用发消息来同步信息，而不是多个线程或者协程直接共享内存。

发送消息是不同语言的高级抽象，实现这一机制时都会使用操作系统提供的锁机制来实现，共享内存也是使用锁这种并发机制实现的。

更为高级和抽象的信息传递方式其实也只是对低抽象级别接口的组合和封装。

Go的channel和Goroutine之间传递消息的方式内部实现时就广泛用到了共享内存和锁，通过对两者进行的组合提供了更高级别的同步机制。