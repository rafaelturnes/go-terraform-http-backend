package domain

import "sync"

type State struct {
	MU   *sync.Mutex
	Lock bool
	Data []byte
}

var (
	InitialState = `{"Version": 1}`
)

type StateInfo struct {
	Lock bool
	ID   string
}
