package sync

import "time"

const (
	waitTime           = 2 * time.Microsecond
	uninitUnlockPanic  = "Unlock of uninitialized mutex"
	uninitRUnlockPanic = "RUnlock of uninitialized mutex"

	decrDelta  = -1
	incrDelta  = 1
	emptyCount = 0
)
