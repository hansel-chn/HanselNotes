package main

import (
	"fmt"
	"sync"
)

/*
array 是相同类型值的集合，数组的长度是其类型的一部分。
数组赋值和传参都会拷贝整个数组的数据，所以数组不是引用类型。
数组的底层数据结构就是其本身，是一个相同类型不同值的顺序排列。所以如果数组位宽不大于 64 位且是 2 的整数次幂（8，16，32，64），那么其并发赋值其实也是安全的，只不过这个大部分情况并非如此，所以其并发赋值是不安全的。

下面以字节数组为例，看下位宽不大于 64 位的并发赋值安全的情况。
可以看到，位宽为 32 位的数组 [4]byte，虽然有四个元素，但是赋值时由一条机器指令完成，所以也是原子操作。
如果你把字节数组的长度换成下面这样子，即使没有超过 64 位，也需要多条指令完成赋值，因为 CPU 中并没有这样位宽的寄存器，需要拆分为多条指令来完成。
*/
func main() {
	var g [5]byte
	var errSet [][5]byte

	var i int
	for ; i < 10000000; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = [...]byte{1, 2, 3, 4, 5}
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = [...]byte{3, 4, 5, 6, 7}
		}()
		wg.Wait()

		// 赋值异常判断
		if !(g == [...]byte{1, 2, 3, 4, 5} || g == [...]byte{3, 4, 5, 6, 7}) {
			fmt.Printf("concurrent assignment error, i=%v g=%+v\n", i, g)
			errSet = append(errSet, g)
			if len(errSet) > 20 {
				break
			}
		}
	}
	//if i == 10000000 {
	//	fmt.Println("no error")
	//}
	if len(errSet) == 0 {
		fmt.Println("no error")
	}
}

// 其实这个已经能说明并发问题了就是没上面直观,只说明了string类型在未修改完全时可以访问,但是没有说明顺序问题
// 上面说明goroutine1的指令a,b和goroutine2的指令c,d,在任务结束并发写协程优雅退出的时刻,仍然会出现执行顺序(a->c->d->b)导致最后的结果出现问题,
// 并不是只在中间读取会有问题

// (编译过程拆分机器指令,一个goroutine和另一个goroutine的多条指令混合到一起了,单条机器指令是原子操作)
//var a = "0"
//
//func main() {
//
//	ch := make(chan string)
//
//	go func() {
//		i := 1
//
//		for {
//
//			if i%2 == 0 {
//				a = "0"
//			} else {
//				a = "aa"
//			}
//			// 如果将sleep去掉，则会读不到脏数据，原因在于编译器做了优化(不知道什么优化,不知道是不是时间片的原因)
//			//time.Sleep(1 * time.Millisecond)
//			i++
//		}
//	}()
//
//	go func() {
//		for {
//			b := a
//			if b != "0" && b != "aa" {
//				ch <- b
//			}
//		}
//	}()
//
//	for i := 0; i < 10; i++ {
//		fmt.Println("Got strange string: ", <-ch)
//	}
//
//}
