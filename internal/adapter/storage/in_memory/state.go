package inmemory

import (
	"bytes"
	"go-terraform-http-backend/internal/adapter/storage"
	"go-terraform-http-backend/internal/core/domain"
	"go-terraform-http-backend/internal/core/port"
	"io"
	"sync"
)

type StateStorage struct {
	data map[string]domain.State
}

func NewStateStorage() port.StateStorage {
	return &StateStorage{
		data: map[string]domain.State{},
	}
}

func (s *StateStorage) Get(id string) (io.Reader, error) {
	if !s.HasState(id) {
		return nil, storage.ErrNotExists
	}

	reader := bytes.NewReader(s.data[id].Data)

	return reader, nil
}
func (s *StateStorage) Update(id string, state io.Reader) error {
	bs, err := io.ReadAll(state)
	if err != nil {
		return err
	}

	// state exists
	if data, ok := s.data[id]; ok {
		current := data
		current.Data = bs
		s.data[id] = current
		return nil
	}

	// new state
	s.data[id] = domain.State{
		MU:   &sync.Mutex{},
		Data: bs,
	}

	return nil
}

func (s *StateStorage) Delete(id string) error {
	if !s.HasState(id) {
		return storage.ErrNotExists
	}

	delete(s.data, id)

	return nil
}

func (s *StateStorage) Lock(id string) error {
	if !s.HasState(id) {
		return storage.ErrNotExists
	}

	if s.IsLocked(id) {
		return storage.ErrAlreadyLocked
	}
	s.data[id].MU.Lock()
	state := s.data[id]
	state.Lock = true
	s.data[id] = state

	return nil

}
func (s *StateStorage) Unlock(id string) error {
	if !s.HasState(id) {
		return storage.ErrNotExists
	}

	if !s.IsLocked(id) {
		return storage.ErrAlreadyUnlocked
	}

	state := s.data[id]
	state.Lock = false
	s.data[id] = state

	s.data[id].MU.Unlock()

	return nil
}

func (s *StateStorage) IsLocked(id string) bool {
	if state, ok := s.data[id]; ok {
		return state.Lock
	}
	return false
}

func (s *StateStorage) HasState(id string) bool {
	if _, ok := s.data[id]; ok {
		return true
	}
	return false
}
