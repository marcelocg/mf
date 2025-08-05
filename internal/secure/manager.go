package secure

import (
	"fmt"
	"mf/internal/types"
)

type Manager struct {
	primary   SecureStorage
	secondary SecureStorage
}

func NewManager() (*Manager, error) {
	var primary, secondary SecureStorage

	keychainProvider := &KeychainProvider{}
	if keychainProvider.IsAvailable() {
		var err error
		primary, err = keychainProvider.GetStorage()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize keychain storage: %w", err)
		}
	}

	encryptedProvider := &EncryptedProvider{}
	var err error
	secondary, err = encryptedProvider.GetStorage()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize encrypted storage: %w", err)
	}

	if primary == nil {
		primary = secondary
		secondary = nil
	}

	return &Manager{
		primary:   primary,
		secondary: secondary,
	}, nil
}

func (m *Manager) Store(account types.Account) error {
	err := m.primary.Store(account)
	if err != nil && m.secondary != nil {
		return m.secondary.Store(account)
	}
	return err
}

func (m *Manager) Retrieve(name string) (*types.Account, error) {
	account, err := m.primary.Retrieve(name)
	if err != nil && m.secondary != nil {
		return m.secondary.Retrieve(name)
	}
	return account, err
}

func (m *Manager) List() ([]string, error) {
	accounts, err := m.primary.List()
	if err != nil && m.secondary != nil {
		return m.secondary.List()
	}
	return accounts, err
}

func (m *Manager) Delete(name string) error {
	err := m.primary.Delete(name)
	if err != nil && m.secondary != nil {
		return m.secondary.Delete(name)
	}
	return err
}
