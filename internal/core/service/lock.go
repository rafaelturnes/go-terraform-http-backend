package service

// import (
// 	"go-terraform-http-backend/internal/core/port"
// )

// type Lock struct {
// 	storage port.LockStorage
// }

// func NewLockService(storage port.LockStorage) port.LockService {
// 	return &Lock{
// 		storage: storage,
// 	}
// }

// func (s *Lock) Lock(id string) bool {
// 	return s.storage.Lock(id)
// }
// func (s *Lock) Unlock(id string) bool {
// 	return s.storage.Unlock(id)
// }
// func (s *Lock) IsLocked(id string) bool {
// 	return s.storage.IsLocked(id)
// }
