package sync

// RWMutex is a fast drop-in replacement for sync.RWMutex
type RWMutex struct {
	rcount      *int64
	wcount      *int64
	initialized bool
}
