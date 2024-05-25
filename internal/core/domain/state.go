package domain

import "sync"

type State struct {
	MU   *sync.Mutex
	Lock bool
	Data []byte
}

// type DefaultState struct {
// 	Version int `json:"version"`
// }

var (
	InitialState = `{"Version": 1}`
)
