package storage

import (
	"errors"
	"fmt"
)

type Err struct {
	Err error
}

func (e Err) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

var (
	ErrNotExists       = Err{Err: errors.New("state not found")}
	ErrAlreadyLocked   = Err{Err: errors.New("state already locked")}
	ErrAlreadyUnlocked = Err{Err: errors.New("state already unlocked")}
)
