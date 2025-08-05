package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"mf/internal/types"
)

type Storage struct {
	configDir string
}

func New() (*Storage, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "mf")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %w", err)
	}

	return &Storage{configDir: configDir}, nil
}

func (s *Storage) SaveAccount(account types.Account) error {
	filename := filepath.Join(s.configDir, account.Name+".json")
	data, err := json.MarshalIndent(account, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal account data: %w", err)
	}

	if err := os.WriteFile(filename, data, 0600); err != nil {
		return fmt.Errorf("failed to write account file: %w", err)
	}

	return nil
}

func (s *Storage) LoadAccount(name string) (*types.Account, error) {
	filename := filepath.Join(s.configDir, name+".json")
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("account '%s' not found", name)
		}
		return nil, fmt.Errorf("failed to read account file: %w", err)
	}

	var account types.Account
	if err := json.Unmarshal(data, &account); err != nil {
		return nil, fmt.Errorf("failed to unmarshal account data: %w", err)
	}

	return &account, nil
}

func (s *Storage) ListAccounts() ([]string, error) {
	entries, err := os.ReadDir(s.configDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read config directory: %w", err)
	}

	var accounts []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
			name := entry.Name()[:len(entry.Name())-5]
			accounts = append(accounts, name)
		}
	}

	return accounts, nil
}

func (s *Storage) DeleteAccount(name string) error {
	filename := filepath.Join(s.configDir, name+".json")
	if err := os.Remove(filename); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("account '%s' not found", name)
		}
		return fmt.Errorf("failed to delete account file: %w", err)
	}

	return nil
}
