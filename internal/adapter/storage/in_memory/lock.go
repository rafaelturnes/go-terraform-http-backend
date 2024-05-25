package inmemory

// import (
// 	"bytes"
// 	"go-terraform-http-backend/internal/core/domain"
// 	"go-terraform-http-backend/internal/core/port"
// 	"io"
// )

// type LockStorage struct {
// 	state map[string]domain.Lock
// }

// func NewLockStorage() port.LockStorage {
// 	return &LockStorage{
// 		state: map[string]domain.Lock{},
// 	}
// }

// func (s *LockStorage) Lock(id string) bool {
// 	if s.IsLocked(id) {
// 		return false
// 	}

// 	s.state[id].MU.
// 		s.state[id].Lock = true

// 	bs, err := io.ReadAll(info)
// 	if err != nil {
// 		s.state.Lock = false
// 		s.mu.Unlock()
// 		return false
// 	}

// 	s.state.Data = bs

// 	return true

// }
// func (s *LockStorage) Unlock(id string) bool {
// 	if s.IsLocked() {
// 		s.state.Lock = false
// 		state := s.state.Data
// 		s.Unlock()
// 		return bytes.NewBuffer(state), true
// 	}

// 	return nil, false

// }
// func (s *LockStorage) IsLocked(id string) bool {
// 	if state, ok := s.state[id]; ok {
// 		return state.Lock
// 	}
// 	return false
// }
