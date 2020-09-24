package unblock

import (
	"errors"
	"time"
)

var ErrMaxTryExceeded = errors.New("max try exceeded")

type retryMutex struct {
	mutex    Mutex
	maxTries int
	interval time.Duration
}

func WithRetry(mutex Mutex, maxTries int, interval time.Duration) Mutex {
	return &retryMutex{
		mutex:    mutex,
		maxTries: maxTries,
		interval: interval,
	}
}

func (m *retryMutex) TryLock() error {
	for i := 0; i < m.maxTries; i++ {
		if err := m.mutex.TryLock(); err == nil {
			return nil
		}
		time.Sleep(m.interval)
	}

	return ErrMaxTryExceeded
}

func (m *retryMutex) TryUnlock() error {
	return m.mutex.TryUnlock()
}
