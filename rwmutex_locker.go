package sync

import (
	"sync"
)

// RLocker returns a Locker interface that implements
// the Lock and Unlock methods by calling rw.RLock and rw.RUnlock.
func (lock *RWMutex) RLocker() sync.Locker {
	return &locker{
		rwmPtr: lock,
	}
}

type locker struct {
	rwmPtr *RWMutex
}

func (l *locker) Lock() {
	l.rwmPtr.RLock()
}
func (l *locker) Unlock() {
	l.rwmPtr.Unlock()
}
