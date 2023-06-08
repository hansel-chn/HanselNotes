# atomic in golang

## atomic.Value

## golang的实现

[https://levelup.gitconnected.com/deep-understanding-of-golang-mutex-9964b02c56e9](https://levelup.gitconnected.com/deep-understanding-of-golang-mutex-9964b02c56e9)

[https://zhuanlan.zhihu.com/p/604656431](https://zhuanlan.zhihu.com/p/604656431)

采用信号量进行阻塞会进行上下文的切换，进入内核态

golang的`mutex`经历过四个版本

* v1版本：

```go
//@formatter:off
func cas(val *int32, old, new int32) bool
func semacquire(*int32)
func semrelease(*int32)

// The structure of the mutex, containing two fields
type Mutex struct {
	key  int32 // Indication of whether the lock is held
	sema int32 // Semaphore dedicated to block/wake up goroutine
}

// Guaranteed to successfully increment the value of delta on val
func xadd(val *int32, delta int32) (new int32) {
	for {
		v := *val
		if cas(val, v, v+delta) {
			return v + delta
		}
	}
	panic("unreached")
}

// request lock
func (m *Mutex) Lock() {
	if xadd(&m.key, 1) == 1 { // Add 1 to the ID, if it is equal to 1, the lock is successfully acquired
		return
	}
	semacquire(&m.sema) // Otherwise block waiting
}

func (m *Mutex) Unlock() {
	if xadd(&m.key, -1) == 0 { // Subtract 1 from the flag, if equal to 0, there are no other waiters
		return
	}
	semrelease(&m.sema) // Wake up other blocked goroutines
}
//@formatter:on
```

在当前的V1版本中，锁和解锁完全基于FIFO方法。虽然这种方法很公平，但从效率的角度来看，它并不是最优的。

* v2版本

```go
//@formatter:off
// A Mutex is a mutual exclusion lock.
// Mutexes can be created as part of other structures;
// the zero value for a Mutex is an unlocked mutex.
type Mutex struct {
    state int32
    sema  uint32
}

const (
    mutexLocked = 1 << iota // mutex is locked
    mutexWoken
    mutexWaiterShift = iota
)
//@formatter:on
```

```go
//@formatter:off
// Lock locks m.
// If the lock is already in use, the calling goroutine
// blocks until the mutex is available.
func (m *Mutex) Lock() {
	// Fast path: grab unlocked mutex.
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}

	awoke := false
	for {
		old := m.state
		new := old | mutexLocked
		if old&mutexLocked != 0 {
			new = old + 1<<mutexWaiterShift
		}
		if awoke {
			// The goroutine has been woken from sleep,
			// so we need to reset the flag in either case.
			new &^= mutexWoken
		}
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			if old&mutexLocked == 0 {
				break
			}
			runtime.Semacquire(&m.sema)
			awoke = true
		}
	}
}
//@formatter:on
```

```go
//@formatter:off
// Unlock unlocks m.
// It is a run-time error if m is not locked on entry to Unlock.
//
// A locked Mutex is not associated with a particular goroutine.
// It is allowed for one goroutine to lock a Mutex and then
// arrange for another goroutine to unlock it.
func (m *Mutex) Unlock() {
	// Fast path: drop lock bit.
	new := atomic.AddInt32(&m.state, -mutexLocked)
	if (new+mutexLocked)&mutexLocked == 0 {
		panic("sync: unlock of unlocked mutex")
	}

	old := new
	for {
		// If there are no waiters or a goroutine has already
		// been woken or grabbed the lock, no need to wake anyone.
		if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken) != 0 {
			return
		}
		// Grab the right to wake someone.
		new = (old - 1<<mutexWaiterShift) | mutexWoken
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			runtime.Semrelease(&m.sema)
			return
		}
		old = m.state
	}
}
//@formatter:on
```

In v2 version, New Goroutine participates in lock competition.

* v3版本

```go
//@formatter:off
// Lock locks m.
// If the lock is already in use, the calling goroutine
// blocks until the mutex is available.
func (m *Mutex) Lock() {
    // Fast path: grab unlocked mutex.
    if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
        if raceenabled {
            raceAcquire(unsafe.Pointer(m))
        }
        return
    }

    awoke := false
    iter := 0
    for {
        old := m.state
        new := old | mutexLocked
        if old&mutexLocked != 0 {
            if runtime_canSpin(iter) {
                // Active spinning makes sense.
                // Try to set mutexWoken flag to inform Unlock
                // to not wake other blocked goroutines.
                if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 &&
                    atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
                    awoke = true
                }
                runtime_doSpin()
                iter++
                continue
            }
            new = old + 1<<mutexWaiterShift
        }
        if awoke {
            // The goroutine has been woken from sleep,
            // so we need to reset the flag in either case.
            if new&mutexWoken == 0 {
                panic("sync: inconsistent mutex state")
            }
            new &^= mutexWoken
        }
        if atomic.CompareAndSwapInt32(&m.state, old, new) {
            if old&mutexLocked == 0 {
                break
            }
            runtime_Semacquire(&m.sema)
            awoke = true
            iter = 0
        }
    }

    if raceenabled {
        raceAcquire(unsafe.Pointer(m))
    }
}
//@formatter:on
```

在v2版本中，若goroutineA占有CPU资源，goroutineB在阻塞队列的头部，goroutineA试图获取锁，但是当前锁被占据，根据v2的策略，goroutineA被阻塞，当锁释放是goroutineB被唤醒获取锁。
v3版本呢允许goroutineA通过自旋等待锁被释放，进一步减少了上下文切换

* v4版本

```go
//@formatter:off
const (
    mutexLocked = 1 << iota // mutex is locked
    mutexWoken
    mutexStarving // separate out a starvation token from the state field
    mutexWaiterShift = iota
    starvationThresholdNs = 1e6    
)
//@formatter:on
```

```go
//@formatter:off

// Lock locks m.
// If the lock is already in use, the calling goroutine
// blocks until the mutex is available.
func (m *Mutex) Lock() {
	// Fast path: grab unlocked mutex.
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		if race.Enabled {
			race.Acquire(unsafe.Pointer(m))
		}
		return
	}
	// Slow path (outlined so that the fast path can be inlined)
	m.lockSlow()
}

func (m *Mutex) lockSlow() {
	var waitStartTime int64
	starving := false
	awoke := false
	iter := 0
	old := m.state
	for {
		// Don't spin in starvation mode, ownership is handed off to waiters
		// so we won't be able to acquire the mutex anyway.
		if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {
			// Active spinning makes sense.
			// Try to set mutexWoken flag to inform Unlock
			// to not wake other blocked goroutines.
			if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 &&
				atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
				awoke = true
			}
			runtime_doSpin()
			iter++
			old = m.state
			continue
		}
		new := old
		// Don't try to acquire starving mutex, new arriving goroutines must queue.
		if old&mutexStarving == 0 {
			new |= mutexLocked
		}
		if old&(mutexLocked|mutexStarving) != 0 {
			new += 1 << mutexWaiterShift
		}
		// The current goroutine switches mutex to starvation mode.
		// But if the mutex is currently unlocked, don't do the switch.
		// Unlock expects that starving mutex has waiters, which will not
		// be true in this case.
		if starving && old&mutexLocked != 0 {
			new |= mutexStarving
		}
		if awoke {
			// The goroutine has been woken from sleep,
			// so we need to reset the flag in either case.
			if new&mutexWoken == 0 {
				throw("sync: inconsistent mutex state")
			}
			new &^= mutexWoken
		}
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			if old&(mutexLocked|mutexStarving) == 0 {
				break // locked the mutex with CAS
			}
			// If we were already waiting before, queue at the front of the queue.
			queueLifo := waitStartTime != 0
			if waitStartTime == 0 {
				waitStartTime = runtime_nanotime()
			}
			runtime_SemacquireMutex(&m.sema, queueLifo, 1)
			starving = starving || runtime_nanotime()-waitStartTime > starvationThresholdNs
			old = m.state
			if old&mutexStarving != 0 {
				// If this goroutine was woken and mutex is in starvation mode,
				// ownership was handed off to us but mutex is in somewhat
				// inconsistent state: mutexLocked is not set and we are still
				// accounted as waiter. Fix that.
				if old&(mutexLocked|mutexWoken) != 0 || old>>mutexWaiterShift == 0 {
					throw("sync: inconsistent mutex state")
				}
				delta := int32(mutexLocked - 1<<mutexWaiterShift)
				if !starving || old>>mutexWaiterShift == 1 {
					// Exit starvation mode.
					// Critical to do it here and consider wait time.
					// Starvation mode is so inefficient, that two goroutines
					// can go lock-step infinitely once they switch mutex
					// to starvation mode.
					delta -= mutexStarving
				}
				atomic.AddInt32(&m.state, delta)
				break
			}
			awoke = true
			iter = 0
		} else {
			old = m.state
		}
	}

	if race.Enabled {
		race.Acquire(unsafe.Pointer(m))
	}
}
//@formatter:on
```

```go
//@formatter:off
func (m *Mutex) Unlock() {
	if race.Enabled {
		_ = m.state
		race.Release(unsafe.Pointer(m))
	}

	// Fast path: drop lock bit.
	new := atomic.AddInt32(&m.state, -mutexLocked)
	if new != 0 {
		// Outlined slow path to allow inlining the fast path.
		// To hide unlockSlow during tracing we skip one extra frame when tracing GoUnblock.
		m.unlockSlow(new)
	}
}

func (m *Mutex) unlockSlow(new int32) {
	if (new+mutexLocked)&mutexLocked == 0 {
		throw("sync: unlock of unlocked mutex")
	}
	if new&mutexStarving == 0 {
		old := new
		for {
			// If there are no waiters or a goroutine has already
			// been woken or grabbed the lock, no need to wake anyone.
			// In starvation mode ownership is directly handed off from unlocking
			// goroutine to the next waiter. We are not part of this chain,
			// since we did not observe mutexStarving when we unlocked the mutex above.
			// So get off the way.
			if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
				return
			}
			// Grab the right to wake someone.
			new = (old - 1<<mutexWaiterShift) | mutexWoken
			if atomic.CompareAndSwapInt32(&m.state, old, new) {
				runtime_Semrelease(&m.sema, false, 1)
				return
			}
			old = m.state
		}
	} else {
		// Starving mode: handoff mutex ownership to the next waiter, and yield
		// our time slice so that the next waiter can start to run immediately.
		// Note: mutexLocked is not set, the waiter will set it after wakeup.
		// But mutex is still considered locked if mutexStarving is set,
		// so new coming goroutines won't acquire it.
		runtime_Semrelease(&m.sema, true, 1)
	}
}
//@formatter:on
```
v4版本增加了对锁饥饿模式的处理