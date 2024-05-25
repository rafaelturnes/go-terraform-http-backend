package service

import (
	"errors"
	"go-terraform-http-backend/internal/adapter/storage"
	"go-terraform-http-backend/internal/core/domain"
	"go-terraform-http-backend/internal/core/port"
	"io"
	"strings"
)

type State struct {
	storage port.StateStorage
}

func NewStateService(storage port.StateStorage) port.StateService {
	return &State{
		storage: storage,
	}
}

func (s *State) Fetch(id string) (io.Reader, error) {
	state, err := s.storage.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, storage.ErrNotExists):
			state := strings.NewReader(domain.InitialState)
			err := s.Update(id, state)
			if err != nil {
				return nil, errors.New("could not crete")
			}

			reader := strings.NewReader(domain.InitialState)

			return reader, nil
		default:
			return nil, err
		}
	}
	return state, nil
}
func (s *State) Update(id string, info io.Reader) error {
	return s.storage.Update(id, info)
}
func (s *State) Delete(id string) error {
	return nil
}

func (s *State) Lock(id string) error {
	return s.storage.Lock(id)
}
func (s *State) Unlock(id string) error {
	return s.storage.Unlock(id)
}
func (s *State) IsLocked(id string) bool {
	return s.storage.IsLocked(id)
}
