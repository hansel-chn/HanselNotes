package main

import (
	"sync/atomic"
	"time"
	"unsafe"
)

// LFStack lock free stack
type LFStack struct {
	Next unsafe.Pointer
	Item int
}

var lfHead unsafe.Pointer // 记录栈头信息

func (head *LFStack) Push(i int) *LFStack { // 强制逃逸
	newStack := &LFStack{Item: i}
	newPtr := unsafe.Pointer(newStack)
	for {
		old := atomic.LoadPointer(&lfHead)
		newStack.Next = old

		if atomic.CompareAndSwapPointer(&lfHead, old, newPtr) {
			break
		}
	}
	return newStack
}

func (head *LFStack) Pop() int {
	for {
		time.Sleep(time.Nanosecond) // 可以让CPU缓一缓
		old := atomic.LoadPointer(&lfHead)
		if old == nil {
			return 0
		}

		if lfHead == old {
			new := (*LFStack)(old).Next
			if atomic.CompareAndSwapPointer(&lfHead, old, new) {
				return 1
			}
		}
	}
}
