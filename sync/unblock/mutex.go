package unblock

import "errors"

var (
	ErrAlreadyLocked   = errors.New("already locked")
	ErrAlreadyUnlocked = errors.New("already unlocked")
	ErrNotOwner        = errors.New("not owner")
)

type Mutex interface {
	TryLock() error
	TryUnlock() error
}