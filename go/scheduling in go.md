# Scheduling in go

[[7]](#ref-7)
[Concurrency in Go](https://edu.anarcho-copy.org/Programming%20Languages/Go/Concurrency%20in%20Go.pdf)
讲的很广，但是似乎缺失一些细节

阅读顺序[[5]](#ref-5) -> [[1]](#ref-1), [[3]](#ref-3) -> [[2]](#ref-2) -> [[8]](#ref-8), [[9]](#ref-9)  ->  [[4]](#ref-4)
，参考文献很多已经讲的很清楚，就不赘述了，针对自己想到的一些问题进行记录

## 从不同概念的线程说起

[[5]](#ref-5)中清晰地讲述了OS-level Threads和Green Threads，Kernel threads vs User Threads的关系。~~Kernel threads vs User
Threads都属于OS-level Threads~~ (这样说更好，User threads and Kernel threads are exactly the same. (You can see by
looking in
/proc/ and see that the kernel threads are there too.))，下文goroutines in Go(包括coroutines in C, fibers in Ruby)等关注的就是，
如何利用Green Threads代替直接使用OS-level Threads实现对并发性能的提升(Goroutine最终仍依赖于OS Thread，即常说的G(
goroutine)，M(
OS thread)，P(processor)三个概念，通过P实现对G的分配和管理)

## 为什么这样设计go scheduler？

goroutine(G)和OS thread(M)通过processor(P)来协调，呈现(M:N关系)

参考总结：

1. 方便实现并发，并且提高Cpu的使用率(参考[[1]](#ref-1)中Practical Example节)
    * 多协程和多线程相比，线程间的切换时间花费更高。对于协程来说，从操作系统来看，两个协程使用了相同的OS Thread and
      Core，内核线程并未陷入waiting state(
      goroutine的上下文切换发生在application level)
   > In the case of using Goroutines, the same OS Thread and Core is being used for all the processing. This means that,
   from the OS’s perspective, the OS Thread never moves into a waiting state; not once. As a result all those
   instructions we lost to context switches when using Threads are not lost when using Goroutines.
2. 与线程相比，goroutine占用内存更小

### work steal机制

最不希望看到M进入waiting state，这会造成core上下文切换导致效率损失。

为了处理哪些情况： 如果没有work
steal机制，当Processor上无待执行和正在执行的Goroutine，P绑定的M会进入wait state(
即使其他P上仍有处于runnable状态的Goroutine) ，OS会在内核中(switch M off core)
。这时即使有一个处于runnable状态的Goroutine，P也不能完成任何工作，直到M(switch back on the core)。

换句话说，保证资源不会发生无意义的浪费，当M空闲的时候抢夺其他P上或global queue的G

### 为什么Goroutine可以进行阻塞调用

[[8]](#ref-8)和[[9]](#ref-9)产生的疑问

> User-level threads requires non-blocking systems call i.e., a multithreaded kernel. Otherwise, entire process will
> blocked in the kernel, even if there are runable threads left in the processes. For example, if one thread causes a
> page
> fault, the process blocks.

这里User-level指的更像是Green thread: 解释如下

> User-level threads require non-blocking system calls because they are managed by the user-level runtime rather than
> the
> operating system. If a user-level thread makes a blocking system call, the entire process, including all other
> runnable
> threads, will be blocked by the kernel. This happens because the kernel is unaware of the user-level threads and
> treats
> the process as a single execution unit. Therefore, when one thread blocks, the kernel blocks the entire process,
> preventing other runnable threads from executing. This is why a multithreaded kernel, which can handle non-blocking
> system calls, is necessary to avoid blocking the entire process.

Green threads是JVM早期版本支持的一种技术，其中N个Java线程可以多路复用到单个操作OS Thread上。需要确保没有阻塞调用。

其实，Goroutine绑定的M也发生阻塞了，只不过GMP的机制解决了该问题，见下方

### System calls带来的Goroutine context-switch

* Asynchronous System Calls
    * 比如 network poller(猜想go net/http库是这么做的？)，通过使用network poller可以避免G阻塞M，M仍然可用，P无需寻找新的M。
      > Goroutine-1 wants to make a network system call, so Goroutine-1 is moved to the network poller and the
      asynchronous network system call is processed. Once Goroutine-1 is moved to the network poller, the M is now
      available to execute a different Goroutine from the LRQ. In this case, Goroutine-2 is context-switched on the M.
* Synchronous System Calls
    * Synchronous System Calls会使G阻塞M,P脱离原先M1，绑定新的M2继续处理LRQ中的G(原M1和G1阻塞)
      > the scheduler is able to identify that Goroutine-1 has caused the M to block. At this point, the scheduler
      detaches M1 from the P with the blocking Goroutine-1 still attached. Then the scheduler brings in a new M2 to
      service the P. At that point, Goroutine-2 can be selected from the LRQ and context-switched on M2. If an M already
      exists because of a previous swap, this transition is quicker than having to create a new M.

注意，[[1]](#ref-1)介绍的是Asynchronous System Calls和Synchronous System Calls;[[3]](#ref-3)介绍的是blocking and
non-blocking System
Calls

### Processor数量影响

```go
package main

import "fmt"
import "runtime"

func main() {
	fmt.Println(runtime.NumCPU())      // 12
	fmt.Println(runtime.GOMAXPROCS(0)) // 12( 默认setting)
}
```

目前未找到明确的依据，在Scheduling in go，processor默认与CPU的逻辑核数相等，直觉如此。 大多数解释如[[4]](#ref-4)
描述，设计是为了平衡并行性、效率和资源利用，如果GOMAXPROCS大于CPU逻辑核数会发生更多的上下文切换，缓存争用，但并不完全如此，有可能会出现提升性能的情况

备注：GOMAXPROCS限制的是processor(context)的数量，当p绑定的thread被阻塞，processor会解绑阻塞的thread并绑定新的thread。
> The GOMAXPROCS variable limits the number of operating system threads that can execute user-level Go code
> simultaneously. There is no limit to the number of threads that can be blocked in system calls on behalf of Go code;
> those do not count against the GOMAXPROCS limit. This package's GOMAXPROCS function queries and changes the limit.

## References

<span id="ref-1"></span>[1] [Scheduling In Go : Part II - Go Scheduler](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)

<span id="ref-2"></span>[2] [Go's work-stealing scheduler](https://rakyll.org/scheduler/)

<span id="ref-3"></span>[3] [[Golang三关-典藏版] Golang 调度器 GMP 原理与调度全分析](https://learnku.com/articles/41728)

<span id="ref-4"></span>[4] [GOMAXPROCS in golang](https://stackoverflow.com/questions/57215184/what-if-gomaxprocs-is-too-large)

<span id="ref-5"></span>[5] [OS-level Threads vs Green Threads, Kernel threads vs User Threads](https://stackoverflow.com/questions/15983872/difference-between-user-level-and-kernel-supported-threads)

<span id="ref-6"></span>[6] [Why blocking system calls blocks entire procedure with user-level threads?](https://stackoverflow.com/questions/40877998/why-blocking-system-calls-blocks-entire-procedure-with-user-level-threads)

<span id="ref-7"></span>[7] [Book: Concurrency in Go](https://edu.anarcho-copy.org/Programming%20Languages/Go/Concurrency%20in%20Go.pdf)

<span id="ref-8"></span>[8] [User-level threads requires non-blocking systems call](http://www.cs.iit.edu/~cs561/cs450/ChilkuriDineshThreads/dinesh%27s%20files/User%20and%20Kernel%20Level%20Threads.html#:~:text=User%2Dlevel%20threads%20requires%20non,page%20fault%2C%20the%20process%20blocks.)

<span id="ref-9"></span>[9] [User-level threads requires non-blocking systems call](https://www.reddit.com/r/golang/comments/14lwmnx/goroutine/)

