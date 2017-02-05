package sync

func (lock *RWMutex) init() {
	var (
		initRCount int64
		initWCount int64
	)
	lock.rcount = &initRCount
	lock.wcount = &initWCount
	lock.initialized = true
}
