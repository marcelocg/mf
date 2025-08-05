package storage

import (
	"fmt"
	"mf/internal/secure"
	"mf/internal/types"
)

type SecureStorage struct {
	manager *secure.Manager
}

func NewSecure() (*SecureStorage, error) {
	manager, err := secure.NewManager()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize secure storage: %w", err)
	}

	return &SecureStorage{manager: manager}, nil
}

func (s *SecureStorage) SaveAccount(account types.Account) error {
	return s.manager.Store(account)
}

func (s *SecureStorage) LoadAccount(name string) (*types.Account, error) {
	return s.manager.Retrieve(name)
}

func (s *SecureStorage) ListAccounts() ([]string, error) {
	return s.manager.List()
}

func (s *SecureStorage) DeleteAccount(name string) error {
	return s.manager.Delete(name)
}
