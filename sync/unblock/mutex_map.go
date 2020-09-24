package unblock

import (
	"sync"
)

var (
	mapMutexesMu = new(sync.Mutex)
	mapMutexes   = make(map[string]struct{})
)

func NewMapMutex(name string) Mutex {
	return &MapMutex{
		name: name,
	}
}

type MapMutex struct {
	name    string
	isOwner bool
}

func (m *MapMutex) TryLock() error {
	mapMutexesMu.Lock()
	defer mapMutexesMu.Unlock()

	if _, locked := mapMutexes[m.name]; locked {
		return ErrAlreadyLocked
	}

	mapMutexes[m.name] = struct{}{}
	m.isOwner = true
	return nil
}

func (m *MapMutex) TryUnlock() error {
	mapMutexesMu.Lock()
	defer mapMutexesMu.Unlock()

	if !m.isOwner {
		return ErrNotOwner
	}

	_, locked := mapMutexes[m.name]
	if !locked {
		return ErrAlreadyUnlocked
	}

	delete(mapMutexes, m.name)
	return nil
}
