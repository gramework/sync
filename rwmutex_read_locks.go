package sync

import (
	"sync/atomic"
	"time"
)

// RLock the RWMutex for read only
func (lock *RWMutex) RLock() {
	if !lock.initialized {
		lock.init()
	}
	atomic.AddInt64(lock.rcount, incrDelta)
	if *lock.wcount == emptyCount {
		return
	}
	for *lock.wcount != emptyCount {
		time.Sleep(waitTime)
	}
}

// RUnlock the RWMutex for read only
func (lock *RWMutex) RUnlock() {
	if !lock.initialized {
		panic(uninitRUnlockPanic)
	}
	atomic.AddInt64(lock.rcount, decrDelta)
}
