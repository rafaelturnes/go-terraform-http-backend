package port

import (
	"io"
)

type StateStorage interface {
	Get(id string) (io.Reader, error)
	Update(id string, state io.Reader) error
	Delete(id string) error

	HasState(id string) bool

	Lock(id string) error
	Unlock(id string) error
	IsLocked(id string) bool
}

type StateService interface {
	// lock
	Lock(id string) error
	Unlock(id string) error
	IsLocked(id string) bool

	// data
	Fetch(id string) (io.Reader, error)
	Update(id string, state io.Reader) error
	Delete(int string) error
}
