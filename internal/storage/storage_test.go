package storage

import (
	"os"
	"path/filepath"
	"testing"

	"mf/internal/types"
)

func TestStorage_SaveAndLoadAccount(t *testing.T) {
	tmpDir := t.TempDir()
	storage := &Storage{configDir: tmpDir}

	account := types.Account{
		Name:   "test-account",
		Secret: "JBSWY3DPEHPK3PXP",
	}

	err := storage.SaveAccount(account)
	if err != nil {
		t.Fatalf("SaveAccount failed: %v", err)
	}

	loadedAccount, err := storage.LoadAccount("test-account")
	if err != nil {
		t.Fatalf("LoadAccount failed: %v", err)
	}

	if loadedAccount.Name != account.Name {
		t.Errorf("Expected name %s, got %s", account.Name, loadedAccount.Name)
	}

	if loadedAccount.Secret != account.Secret {
		t.Errorf("Expected secret %s, got %s", account.Secret, loadedAccount.Secret)
	}
}

func TestStorage_LoadNonExistentAccount(t *testing.T) {
	tmpDir := t.TempDir()
	storage := &Storage{configDir: tmpDir}

	_, err := storage.LoadAccount("non-existent")
	if err == nil {
		t.Error("Expected error when loading non-existent account")
	}
}

func TestStorage_ListAccounts(t *testing.T) {
	tmpDir := t.TempDir()
	storage := &Storage{configDir: tmpDir}

	accounts := []types.Account{
		{Name: "account1", Secret: "SECRET1"},
		{Name: "account2", Secret: "SECRET2"},
		{Name: "account3", Secret: "SECRET3"},
	}

	for _, account := range accounts {
		err := storage.SaveAccount(account)
		if err != nil {
			t.Fatalf("SaveAccount failed: %v", err)
		}
	}

	accountNames, err := storage.ListAccounts()
	if err != nil {
		t.Fatalf("ListAccounts failed: %v", err)
	}

	if len(accountNames) != 3 {
		t.Errorf("Expected 3 accounts, got %d", len(accountNames))
	}

	expectedNames := map[string]bool{
		"account1": true,
		"account2": true,
		"account3": true,
	}

	for _, name := range accountNames {
		if !expectedNames[name] {
			t.Errorf("Unexpected account name: %s", name)
		}
	}
}

func TestStorage_DeleteAccount(t *testing.T) {
	tmpDir := t.TempDir()
	storage := &Storage{configDir: tmpDir}

	account := types.Account{
		Name:   "test-account",
		Secret: "JBSWY3DPEHPK3PXP",
	}

	err := storage.SaveAccount(account)
	if err != nil {
		t.Fatalf("SaveAccount failed: %v", err)
	}

	err = storage.DeleteAccount("test-account")
	if err != nil {
		t.Fatalf("DeleteAccount failed: %v", err)
	}

	_, err = storage.LoadAccount("test-account")
	if err == nil {
		t.Error("Expected error when loading deleted account")
	}
}

func TestStorage_DeleteNonExistentAccount(t *testing.T) {
	tmpDir := t.TempDir()
	storage := &Storage{configDir: tmpDir}

	err := storage.DeleteAccount("non-existent")
	if err == nil {
		t.Error("Expected error when deleting non-existent account")
	}
}

func TestNew(t *testing.T) {
	storage, err := New()
	if err != nil {
		t.Fatalf("New failed: %v", err)
	}

	if storage.configDir == "" {
		t.Error("Config directory should not be empty")
	}

	homeDir, _ := os.UserHomeDir()
	expectedDir := filepath.Join(homeDir, ".config", "mf")
	if storage.configDir != expectedDir {
		t.Errorf("Expected config dir %s, got %s", expectedDir, storage.configDir)
	}
}
