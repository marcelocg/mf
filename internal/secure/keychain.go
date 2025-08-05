package secure

import (
	"encoding/json"
	"fmt"

	"github.com/zalando/go-keyring"
	"mf/internal/types"
)

const serviceName = "mf-totp"

type KeychainStorage struct{}

type KeychainProvider struct{}

func (p *KeychainProvider) IsAvailable() bool {
	err := keyring.Set(serviceName+"-test", "test", "test")
	if err != nil {
		return false
	}
	keyring.Delete(serviceName+"-test", "test")
	return true
}

func (p *KeychainProvider) GetStorage() (SecureStorage, error) {
	return &KeychainStorage{}, nil
}

func (k *KeychainStorage) Store(account types.Account) error {
	data, err := json.Marshal(account)
	if err != nil {
		return fmt.Errorf("failed to marshal account: %w", err)
	}

	err = keyring.Set(serviceName, account.Name, string(data))
	if err != nil {
		return fmt.Errorf("failed to store in keychain: %w", err)
	}

	return nil
}

func (k *KeychainStorage) Retrieve(name string) (*types.Account, error) {
	data, err := keyring.Get(serviceName, name)
	if err != nil {
		return nil, fmt.Errorf("account '%s' not found in keychain: %w", name, err)
	}

	var account types.Account
	if err := json.Unmarshal([]byte(data), &account); err != nil {
		return nil, fmt.Errorf("failed to unmarshal account data: %w", err)
	}

	return &account, nil
}

func (k *KeychainStorage) List() ([]string, error) {
	return nil, fmt.Errorf("keychain does not support listing all entries")
}

func (k *KeychainStorage) Delete(name string) error {
	err := keyring.Delete(serviceName, name)
	if err != nil {
		return fmt.Errorf("failed to delete account '%s' from keychain: %w", name, err)
	}
	return nil
}
