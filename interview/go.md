# golang

## why gmp

> 并发效率问题，操作系统线程切换带来的巨大损失
>
>  怎么样解决，gmp模型
>
>  历史上，M:1结构的，M:N结构的
>
>  why M:N，一般会说M:1,
>
> 1. 协程阻塞影响并发（感觉重点其实是未提供processor类似的解绑P和M的能力）
> 2. 还有就是运用不了多核的能力
>
> 具体怎么实现的GMP，有哪些机制，Work-sharing，work-stealing，spinning面对系统调用阻塞时的处理，网络io，文件io

## golang Gc

垃圾回收的发展 标记-清楚法 -> 三色标记法

标记-清楚法的问题： 需要STW（不执行任何用户代码），暂停程序造成卡顿，于是三色标记法

怎么解决STW的 -> 三色标记法仍然是会进行STW，但是减少了其时间
为什么仍需STW，因为三色标记法本质用三种颜色来进行对象的筛选，

* 白色-代表潜在的被回收的垃圾
* 黑色-表示了从根部可到达的对象，且不引用外部指针的对象
* 灰色-代表了仍存在引用外部指针的对象

> * 白色：表示对象尚未被垃圾回收器访问到。所有对象在垃圾回收开始时都是白色的。
> * 灰色：表示对象已经被访问到，但其引用的对象还没有被完全扫描。灰色对象需要进一步处理。
> * 黑色：表示对象已经被访问到，并且其引用的所有对象也都已经被扫描。黑色对象不需要进一步处理。

三色标记法从灰色对象开始扫描，并且到灰色对象不存在时结束。

如果没有STW会出现的问题：
已经被扫瞄过的黑色对象，指向了原本被其它对象引用的一个白色对象，但是其原有引用断开导致最后出现黑色对象引用白色对象的情况，在清除白色对象时出现删除正在被引用对象的情况

为了解决STW时间长或对象丢的的问题，提出了屏障机制,下面这个描述很好，屏障机制更像钩子函数，在读对象，创建对象和更新对象时，插入一段代码影响着色
> The barrier technique in garbage collection is more like a hook method, which is a piece of code that is executed when
> the user program reads an object, creates a new object, and updates the object pointer.

https://www.sobyte.net/post/2021-12/golang-garbage-collector/
https://www.jianshu.com/p/4c5a303af470

 ```go
 // 强三色不变式:黑色对象不会直接引用白色对象
// Dijkstra’s写屏障(Inserting Write Barriers)
// 通过在赋值前着色，保证黑色节点不会指向白色节点
// 缺点，对于栈上的对象（一般被视为垃圾回收的根对象），必须对其添加写屏障，或者在扫描结束后STW，再次扫描栈。
writePointer(slot, ptr):
shade(ptr)
*slot = ptr

```

 ```go
 // 弱三色不变式:在某些情况下黑色对象引用白色对象，但通过在赋值操作前对旧引用的对象进行着色来避免问题。
// Yuasa's屏障
// GC开始时STW扫描堆栈来记录初始快照，这个过程会保护开始时刻的所有存活对象。
// 在删除节点slot时对其着色，保证不被错误回收，回收精度低
writePointer(slot, ptr)
shade(*slot)
*slot = ptr

```

 ```go
 // 混合写屏障
// 目的是通过两种混合，使得栈上对象不需要对其添加写屏障，且极少的STW时间,且较高的GC精度（融合了上述的优点，摒弃了缺点）
// GC开始将栈上的对象全部扫描并标记为黑色(之后不再进行第二次重复扫描，无需STW)，
// GC期间，任何在栈上创建的新对象，均为黑色。
// 被删除的对象标记为灰色。
// 添加的对象标记为灰色。

//1 
标记灰色(当前下游对象slot) //只要当前下游对象被移走，就标记灰色

//2 
标记灰色(新下游对象ptr)

//3
当前下游对象slot = 新下游对象ptr
}

```

## golang int concurrency

即使cpu为64位可以完全容纳int类型的长度，仍然有可能出现问题，这是系统没有保证的，比如CPU处理时内存是否对齐，编译器是否对其有未知的优化，重排序导致未知的情况发生
https://stackoverflow.com/questions/34750323/is-variable-assignment-atomic-in-go
https://gobyexample.com/atomic-counters
https://blog.csdn.net/ljfrocky/article/details/137677118
https://stackoverflow.com/questions/15466978/what-is-the-difference-between-sequential-consistency-and-atomicity

原子性和顺序一致性
其实golang1.22后保证了，对小于机器字大小的内存访问是原子的(但是好像不是我们说的原子性，比如上面对单个变量的累加)

## memory escape

In Go, memory escape refers to a situation where a variable is allocated on the heap instead of the stack. This happens
when the Go compiler determines that the variable's lifetime exceeds the scope of the function or goroutine in which it
is created. Escape analysis is the process the Go compiler uses to decide whether a variable should be allocated on the
stack or the heap.  
Common Causes of Memory Escapes

* Returning pointers to local variables:  
  If a local variable's address is returned, it escapes to the heap because the stack frame will be destroyed after the
  function returns.
* Interface conversions:  
  Assigning a value to an interface type may cause the value to escape if the interface outlives the function.
* Closures:Variables captured by closures escape to the heap because the closure may outlive the function.
* Slice or map growth:  
  When slices or maps grow, their underlying arrays may be reallocated on the heap.
* Large objects:
  Large objects may be allocated on the heap to avoid stack overflow.

## golang worker pool

用来管理多个任务的执行，在一个固定个数的goroutine下面，以防止资源的耗尽
限制并发工作的数量，以避免压倒系统资源

## What if GOMAXPROCS is too large?

there will be more context switches if there are more active threads than cores

## golang不会自己让出时间片防止饥饿问题吗

Golang 的调度器不会主动强制让出时间片来防止饥饿问题。Golang 的调度器是协作式调度（cooperative scheduling），这意味着
Goroutine 必须主动让出 CPU 时间片，调度器才能切换到其他 Goroutines。

如果一个 Goroutine 进入了死循环且没有任何阻塞操作（如 I/O、channel 操作）或显式调用 `runtime.Gosched()`，调度器将无法介入，这会导致该
Goroutine 独占 CPU，其他 Goroutines 无法被调度，进而可能导致饥饿问题。

### 解决方法

1. **显式调用 `runtime.Gosched()`**  
   在长时间运行的 Goroutine 中，定期调用 `runtime.Gosched()` 主动让出 CPU 时间片，允许调度器切换到其他 Goroutines。

   ```go
   package main

   import (
       "runtime"
   )

   func main() {
       go func() {
           for {
               // 主动让出时间片
               runtime.Gosched()
           }
       }()

       select {} // 阻塞主 Goroutine
   }
   ```

2. **使用阻塞操作**  
   在 Goroutine 中引入阻塞操作（如 channel 或 `time.Sleep`），调度器会在阻塞时自动切换到其他 Goroutines。

   ```go
   package main

   import (
       "time"
   )

   func main() {
       go func() {
           for {
               time.Sleep(1 * time.Millisecond) // 阻塞操作
           }
       }()

       select {} // 阻塞主 Goroutine
   }
   ```

3. **优化代码逻辑**  
   避免死循环或长时间运行的 Goroutine，确保 Goroutines 能够及时退出或让出 CPU。

Golang 的调度器设计是为了高效运行 Goroutines，但需要开发者注意避免死循环或饥饿问题。

## 那如果golang有一个长时间的数据处理任务，不人为让出时间片，会导致该goroutine长期占用P导致其他goroutine不可以执行？

Yes, if a goroutine is performing a long-running task without any blocking operations (e.g., I/O, channel operations) or
explicit calls to `runtime.Gosched()`, it can monopolize the **P** (Processor) it is running on. This prevents other
goroutines from being scheduled on that **P**, leading to starvation for those goroutines.

Golang's scheduler relies on cooperative scheduling, meaning that goroutines must yield control voluntarily. Without
yielding, the scheduler cannot preempt the long-running goroutine, and other goroutines will not get a chance to execute
on the same **P**.

### How to Mitigate This Issue

1. **Use `runtime.Gosched()`**: Explicitly yield control to allow the scheduler to run other goroutines.
2. **Introduce Blocking Operations**: Use operations like `time.Sleep`, channel communication, or I/O to allow the
   scheduler to switch to other goroutines.
3. **Break the Task into Smaller Chunks**: Divide the long-running task into smaller units of work and yield control
   between chunks.

### Example of a Long-Running Task with `runtime.Gosched()`

```go
package main

import (
	"fmt"
	"runtime"
)

func longRunningTask() {
	for i := 0; i < 1000000; i++ {
		if i%1000 == 0 {
			// Yield control to allow other goroutines to run
			runtime.Gosched()
		}
		fmt.Println(i)
	}
}

func main() {
	go longRunningTask()

	// Another goroutine
	go func() {
		for {
			fmt.Println("Other goroutine running")
			runtime.Gosched()
		}
	}()

	select {} // Block main goroutine
}
```

In this example, `runtime.Gosched()` ensures that the long-running task does not monopolize the **P**, allowing other
goroutines to execute.

## 为什么golang后面使用连续栈而不是分段栈

Golang switched from segmented stacks to contiguous stacks primarily due to performance and complexity reasons. Here's a
breakdown:

### 1. **Performance Issues with Segmented Stacks**

- **Frequent Stack Splitting and Joining**: Segmented stacks require splitting and joining stacks as goroutines grow or
  shrink. This introduces overhead during function calls, especially for deep recursion or large stack usage.
- **Pointer Adjustments**: When a stack segment is resized, all pointers to variables on the stack need to be adjusted,
  which is computationally expensive.
- **Cache Performance**: Segmented stacks can lead to poor CPU cache performance because stack segments are not
  contiguous in memory, causing more cache misses.

### 2. **Simpler Implementation with Contiguous Stacks**

- **Dynamic Growth**: Contiguous stacks in Go start small (e.g., 2 KB) and grow dynamically by allocating a larger stack
  and copying the existing stack into it. This avoids the need for managing multiple segments.
- **Better Cache Locality**: Since the stack is contiguous, it improves cache locality, leading to better performance.
- **Simplified Code**: Contiguous stacks simplify the runtime implementation, making it easier to maintain and optimize.

### 3. **Reduced Overhead**

- With contiguous stacks, the runtime avoids the frequent checks and adjustments required for segmented stacks, reducing
  the overhead during function calls.

### 4. **Compatibility with Modern Hardware**

- Modern hardware and operating systems are optimized for contiguous memory access, making contiguous stacks a better
  fit for performance-critical applications.

In summary, Golang moved to contiguous stacks to improve performance, simplify the runtime, and better utilize modern
hardware capabilities.

## 通道nil的相关问题

比如在合并两个通道（接受两个通道的消息的时候，简单close通道，此时一个通道尚未关闭正在阻塞中，会造成busy loop，造成性能损失）

## 两个线程轮流打印数字，一直到100

package main

import (
"fmt"
"sync"
)

func main() {
var mu sync.Mutex
cond := sync.NewCond(&mu)
counter := 1
max := 100

	// Function for thread A
	go func() {
		for {
			mu.Lock()
			for counter%2 == 0 {
				cond.Wait() // Wait for the condition to be signaled
			}
			if counter > max {
				mu.Unlock()
				cond.Broadcast()
				return
			}
			fmt.Println("Thread A:", counter)
			counter++
			mu.Unlock()
			cond.Signal() // Signal the other thread
		}
	}()

	// Function for thread B
	go func() {
		for {
			mu.Lock()
			for counter%2 != 0 {
				cond.Wait() // Wait for the condition to be signaled
			}
			if counter > max {
				mu.Unlock()
				cond.Broadcast()
				return
			}
			fmt.Println("Thread B:", counter)
			counter++
			mu.Unlock()
			cond.Signal() // Signal the other thread
		}
	}()

	// Wait for both threads to finish
	mu.Lock()
	for counter <= max {
		cond.Wait()
	}
	mu.Unlock()

}

## ch <- struct{}{} ready = true // Potentially unsafe modification 编译器不会使两条命令有可能不按照顺序执行？

Yes, without proper synchronization, the compiler or CPU may reorder these instructions for optimization purposes, which
could lead to unexpected behavior in concurrent programs. This is because Go, like many modern programming languages,
allows instruction reordering as long as it does not violate the program's single-threaded semantics.

In your example:

```go
ch <- struct{}{}
ready = true // Potentially unsafe modification
```

If another goroutine reads `ready` after receiving from the channel, it might observe `ready` as `false` because the
write to `ready` could be reordered to occur after the channel send operation.

### How to Prevent Reordering

To ensure proper ordering and visibility across goroutines, you need to use synchronization primitives like
`sync.Mutex`, `sync/atomic`, or channels. These primitives act as memory barriers, preventing reordering and ensuring
visibility of changes.

### Example Using `sync.Mutex`

```go
package main

import (
	"sync"
)

func main() {
	var mu sync.Mutex
	ready := false
	ch := make(chan struct{})

	go func() {
		<-ch
		mu.Lock()
		if ready {
			println("Condition met")
		}
		mu.Unlock()
	}()

	mu.Lock()
	ready = true
	mu.Unlock()
	ch <- struct{}{}
}
```

### Example Using `sync/atomic`

```go
package main

import (
	"sync/atomic"
)

func main() {
	var ready int32
	ch := make(chan struct{})

	go func() {
		<-ch
		if atomic.LoadInt32(&ready) == 1 {
			println("Condition met")
		}
	}()

	atomic.StoreInt32(&ready, 1)
	ch <- struct{}{}
}
```

Both approaches ensure that the `ready` variable is updated in a way that prevents reordering and guarantees visibility
to other goroutines.

## ready := false ch := make(chan struct{})  go func() { <-ch // Wait for the signal if ready { // Potentially unsafe access fmt.Println("Condition met") } }()  ready = true // Potentially unsafe modification ch <- struct{}{} 结合前面描述和当前代码，会同时出现上述 两个问题是吧

Yes, the provided code can exhibit **both issues**:

1. **Race Condition**: The `ready` variable is accessed by one goroutine (`if ready { ... }`) and modified by another (
   `ready = true`) without proper synchronization. This can lead to a race condition.

2. **Instruction Reordering**: Without synchronization, the compiler or CPU may reorder the instructions `ready = true`
   and `ch <- struct{}{}`. This could result in the receiving goroutine observing `ready` as `false` even after the
   channel send operation.

To avoid these issues, proper synchronization mechanisms like `sync.Mutex` or `sync/atomic` should be used.

Always release the lock (mu.Unlock()) before signaling (cond.Broadcast() or cond.Signal()) to ensure that waiting
goroutines can proceed without being blocked by the mutex.

## 如果wait的协程尝试获取锁很快，但是由于锁未解锁，导致获取锁失败进入阻塞状态，那signal就失效了？两个协程都被阻塞了？

`signal` 不会失效，即使 `wait` 的协程在尝试获取锁时失败并进入阻塞状态，`signal` 的效果仍然有效。以下是详细解释：

1. **`signal` 的作用**：
    - `signal` 唤醒一个正在 `wait` 的协程，使其从等待队列中移出并进入锁竞争状态。
    - 唤醒的协程会尝试获取锁，但如果锁未释放，它会再次进入阻塞状态，等待锁被释放。

2. **两个协程的状态**：
    - **原始协程**：在调用 `signal` 后，仍然持有锁，继续执行后续代码，直到调用 `unlock` 释放锁。
    - **被唤醒的协程**：被 `signal` 唤醒后，尝试获取锁。如果锁未释放，它会进入阻塞状态，等待锁被释放。

3. **不会导致死锁**：
    - 只要原始协程最终调用 `unlock` 释放锁，被唤醒的协程会成功获取锁并继续执行。
    - `signal` 的效果不会丢失，因为被唤醒的协程已经从等待队列中移出，只是暂时阻塞在锁的获取上。

### 总结

`signal` 的效果不会失效。即使 `wait` 的协程因锁未释放而阻塞，只要原始协程最终释放锁，`wait`
的协程会继续执行。这种机制确保了线程同步的正确性，不会导致两个协程永久阻塞。

## golang 锁的内部实现和atmic包的区别有哪些，优缺点有哪些

### **Differences Between `sync.Mutex` and `sync/atomic`**

#### **1. Internal Implementation**

- **`sync.Mutex`**:
    - Implements a **blocking lock** using operating system primitives (e.g., futex on Linux).
    - When contention occurs, goroutines are parked (blocked) and placed in a wait queue.
    - Relies on the Go runtime scheduler to manage goroutines waiting for the lock.
    - Suitable for protecting critical sections where multiple goroutines need exclusive access.

- **`sync/atomic`**:
    - Provides **low-level atomic operations** (e.g., `atomic.AddInt32`, `atomic.CompareAndSwap`).
    - Uses CPU instructions (e.g., `LOCK` prefix on x86) to perform atomic operations without blocking.
    - Does not involve the Go runtime scheduler.
    - Suitable for fine-grained synchronization or lock-free algorithms.

---

#### **2. Performance**

- **`sync.Mutex`**:
    - Higher overhead due to potential blocking and context switching.
    - Performs better when contention is low or when critical sections are large.
    - Can lead to goroutine parking and waking, which involves the Go runtime.

- **`sync/atomic`**:
    - Extremely fast for simple operations (e.g., incrementing counters).
    - No blocking or context switching, making it ideal for high-performance scenarios.
    - Limited to simple atomic operations; not suitable for complex critical sections.

---

#### **3. Usability**

- **`sync.Mutex`**:
    - Easier to use for protecting larger critical sections or shared resources.
    - Provides clear semantics for locking and unlocking.

- **`sync/atomic`**:
    - Requires careful design to avoid race conditions.
    - Limited to specific use cases (e.g., counters, flags, or single variables).
    - More error-prone for complex synchronization.

---

### **Advantages and Disadvantages**

#### **`sync.Mutex`**

- **Advantages**:
    - Simple to use for protecting shared resources.
    - Handles complex critical sections and multiple variables.
    - Avoids busy-waiting by blocking goroutines.

- **Disadvantages**:
    - Higher overhead due to blocking and context switching.
    - Can lead to contention and reduced performance under high load.

#### **`sync/atomic`**

- **Advantages**:
    - Extremely fast for simple operations.
    - No blocking or context switching, reducing overhead.
    - Ideal for high-performance, low-latency scenarios.

- **Disadvantages**:
    - Limited to simple atomic operations.
    - Difficult to use for complex synchronization.
    - Error-prone if not carefully designed.

---

### **When to Use**

- Use **`sync.Mutex`** when:
    - Protecting complex critical sections or multiple shared variables.
    - Contention is low, or blocking is acceptable.

- Use **`sync/atomic`** when:
    - Performing simple atomic operations (e.g., counters, flags).
    - High performance and low latency are critical.

## Go runtime scheduler 如何管理锁的

The Go runtime scheduler manages locks by coordinating goroutines and ensuring proper synchronization using mechanisms
like **`sync.Mutex`**, **`sync.Cond`**, and **atomic operations**. Here's how it works:

### 1. **Mutex Locking and Unlocking**

- **Lock Acquisition**:
    - When a goroutine calls `Lock()` on a `sync.Mutex`, it attempts to acquire the lock using atomic operations (e.g.,
      `compare-and-swap`).
    - If the lock is free, the goroutine acquires it and continues execution.
    - If the lock is already held, the goroutine is **parked** (blocked) and added to a wait queue.

- **Unlocking**:
    - When a goroutine calls `Unlock()`, it releases the lock and wakes up one or more goroutines waiting in the queue.
    - The Go runtime uses **`runtime_Semrelease`** to notify the waiting goroutines.

### 2. **Parking and Unparking Goroutines**

- The Go runtime uses **semaphores** to park and unpark goroutines:
    - **Parking**: A goroutine that cannot acquire a lock is put to sleep using `runtime_Semacquire`.
    - **Unparking**: When the lock is released, the runtime wakes up a waiting goroutine using `runtime_Semrelease`.

### 3. **Fairness and Spinning**

- The runtime ensures fairness by maintaining a queue of waiting goroutines.
- In some cases, the runtime allows **spinning** (active waiting) for a short time before parking a goroutine. This
  reduces the overhead of context switching for short lock contention.

### 4. **Condition Variables (`sync.Cond`)**

- The runtime uses **notify lists** to manage goroutines waiting on condition variables.
- When `Wait()` is called, the goroutine is parked and added to the notify list.
- `Signal()` or `Broadcast()` wakes up one or all waiting goroutines using `runtime_notifyListNotifyOne` or
  `runtime_notifyListNotifyAll`.

### 5. **Atomic Operations**

- The runtime uses atomic operations (e.g., `atomic.CompareAndSwap`) for lock-free synchronization.
- These operations are used to implement fast paths for acquiring and releasing locks without involving the kernel.

### 6. **Integration with the Scheduler**

- The Go scheduler ensures that runnable goroutines (e.g., those woken up after acquiring a lock) are scheduled on
  available threads.
- It balances goroutines across multiple threads (M) and processors (P) to maximize concurrency.

This combination of atomic operations, semaphores, and the scheduler ensures efficient and fair lock management in Go.  