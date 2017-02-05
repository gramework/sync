package sync

import (
	"sync/atomic"
	"time"
)

// Lock the RWMutex for read and write
func (lock *RWMutex) Lock() {
	if !lock.initialized {
		lock.init()
	}
	atomic.AddInt64(lock.wcount, incrDelta)

	for *lock.rcount != emptyCount || *lock.wcount != incrDelta {
		time.Sleep(waitTime)
	}
}

// Unlock the RWMutex for read and write
func (lock *RWMutex) Unlock() {
	if !lock.initialized {
		panic(uninitUnlockPanic)
	}
	atomic.AddInt64(lock.wcount, decrDelta)
}
