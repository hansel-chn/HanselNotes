# Context

[https://tech.ipalfish.com/blog/2020/03/30/golang-context/](https://tech.ipalfish.com/blog/2020/03/30/golang-context/)

-> blog耗时操作部分存在memory leak

[https://www.arangodb.com/2020/09/a-story-of-a-memory-leak-in-go-how-to-properly-use-time-after/](https://www.arangodb.com/2020/09/a-story-of-a-memory-leak-in-go-how-to-properly-use-time-after/)
或者用计时器在 `<-ctx.Done()`后处理（人为关闭计时器，关闭时计时器已经停止就消费掉 `<-delay.C`）

```go
package main

import "time"

func main() {
	//memory leak
	//select {
	//  case <-time.After(time.Second):
	//     // do something after 1 second.
	//  case <-ctx.Done():
	//     // do something when context is finished.
	//     // resources created by the time.After() will not be garbage collected
	//  }

	delay := time.NewTimer(time.Second)

	select {
	case <-delay.C:
		// do something after one second.
	case <-ctx.Done():
		// do something when context is finished and stop the timer.
		if !delay.Stop() {
			// if the timer has been stopped then read from the channel.
			<-delay.C
		}
	}
}
```

[https://stackoverflow.com/questions/73611358/timeout-on-function-and-goroutine-leak](https://stackoverflow.com/questions/73611358/timeout-on-function-and-goroutine-leak)
采用buffer为1的channel，不需要关闭channel，因为不阻塞，channel随gc清理；超时退出 `defer cancel()`处理

* 例子

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func foo(ctx context.Context) (string, error) {
	ch := make(chan string, 1)

	go func() {
		fmt.Println("Sleeping...")
		time.Sleep(time.Second * 1)
		fmt.Println("Wake up...")
		ch <- "foo"
	}()

	select {
	case <-ctx.Done():
		return "", fmt.Errorf("context cancelled: %w", ctx.Err())
	case result := <-ch:
		return result, nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, err := foo(ctx)
	if err != nil {
		log.Fatalf("foo failed: %v", err)
	}

	log.Printf("res: %s", res)
}
```

同时解决了select可能会出现的内存泄露，

1. 正常情况，ctx.done退出，阻塞的 `ch := make(chan string, 1)`,通过buffer为1的channel通过gc处理
2. 异常情况超时退出，阻塞的 `ctx.Done`，通过 `defer cancel()`关闭通道。

## 通过ctx控制所有协程退出

* 通过share memory

```go
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	parent, pCancel := context.WithCancel(context.Background())
	child, _ := context.WithCancel(parent)
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go work(wg, child)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	defer signal.Stop(c)

	select {
	case <-c:
		pCancel()
		fmt.Println("Waiting everyone to finish...")
		wg.Wait()
		fmt.Println("Exiting")
		os.Exit(0)
	}
}

func work(wg *sync.WaitGroup, ctx context.Context) {
	done := false
	wg.Add(1)
	for !done {
		fmt.Println("Doing something...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			done = true
		default:

		}
	}
	wg.Done()
}
```

* 通过communicating

```go
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	parent, pCancel := context.WithCancel(context.Background())
	child, _ := context.WithCancel(parent)
	done := make(chan struct{})
	jobsCount := 10

	for i := 0; i < jobsCount; i++ {
		go work(child, done)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	defer signal.Stop(c)

	select {
	case <-c:
		pCancel()
		fmt.Println("Waiting everyone to finish...")
		for i := 0; i < jobsCount; i++ {
			<-done
		}
		fmt.Println("Exiting")
		os.Exit(0)
	}
}

func work(ctx context.Context, doneChan chan struct{}) {
	done := false
	for !done {
		fmt.Println("Doing something...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			done = true
		default:

		}
	}
	doneChan <- struct{}{}
}
```

## 源码解析

### parentCancelCtx

```go
package context

func parentCancelCtx(parent Context) (*cancelCtx, bool) {
	done := parent.Done()
	if done == closedchan || done == nil {
		return nil, false
	}
	p, ok := parent.Value(&cancelCtxKey).(*cancelCtx)
	if !ok {
		return nil, false
	}
	pdone, _ := p.done.Load().(chan struct{})
	if pdone != done {
		return nil, false
	}
	return p, true
}
```

* 判断parent是否为 `*cancelCtx`，通过 `Value`判断。`Value`传参 `cancelCtxKey`时，只有返回 `cancelCtx`满足类型断言

```go
package context

func (c *cancelCtx) Value(key any) any {
	if key == &cancelCtxKey {
		return c
	}
	return value(c.Context, key)
}
```

```go
package context

func (c *valueCtx) Value(key any) any {
	if c.key == key {
		return c.val
	}
	return value(c.Context, key)
}
```

```go
package context

func value(c Context, key any) any {
	for {
		switch ctx := c.(type) {
		case *valueCtx:
			if key == ctx.key {
				return ctx.val
			}
			c = ctx.Context
		case *cancelCtx:
			if key == &cancelCtxKey {
				return c
			}
			c = ctx.Context
		case *timerCtx:
			if key == &cancelCtxKey {
				return &ctx.cancelCtx
			}
			c = ctx.Context
		case *emptyCtx:
			return nil
		default:
			return c.Value(key)
		}
	}
}
```

* `Value`代码如上

```go
package context

func propagateCancel(parent Context, child canceler) {
	/*
		...省略...
	*/

	if p, ok := parentCancelCtx(parent); ok {
		p.mu.Lock()
		if p.err != nil {
			// parent has already been canceled
			child.cancel(false, p.err)
		} else {
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
	} else {
		atomic.AddInt32(&goroutines, +1)
		go func() {
			select {
			case <-parent.Done():
				child.cancel(false, parent.Err())
			case <-child.Done():
			}
		}()
	}
}
```

* 若 `parent`类型断言为 `*cancelCtx`，添加 `children`
* 若否，新启一个协程，监控子ctx和父ctx关闭

