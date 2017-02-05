package sync

// A Locker represents an object that can be locked and unlocked.
type Locker interface {
	Lock()
	Unlock()
}
