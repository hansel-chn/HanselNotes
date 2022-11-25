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
* 第二种说法：我建议不要使用json.Decoder。它用于JSON对象流，而不是单个对象。对于单个JSON对象，它的效率并不高，因为它将整个对象读入内存。它的缺点是，如果在对象后面包含了垃圾，它就不会报错。取决于几个因素，json。解码器可能无法完全读取正文，连接将不符合重用的条件。

>It really depends on what your input is. If you look at the implementation of the Decode method of json.Decoder, <mark>**it buffers the entire JSON value in memory before unmarshalling it into a Go value. So in most cases it won't be any more memory efficient**</mark> (although this could easily change in a future version of the language). 
> 
>So a better rule of thumb is this:
>
>Use json.Decoder if your data is coming from an io.Reader stream, or you need to decode multiple values from a stream of data.
Use json.Unmarshal if you already have the JSON data in memory.
For the case of reading from an HTTP request, I'd pick json.Decoder since you're obviously reading from a stream.
> 
> <mark>高亮话指的是和unmarshal相比（从数据流读到内存，再unmarshal）相比，decoder将io.Reader(res.body)整个值又额外缓存了一次，先缓存到decoder里的io.reader中,再进行（从数据流读到内存，再unmarshal）/mark>