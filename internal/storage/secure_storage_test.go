package storage

import (
	"testing"

	"mf/internal/types"
)

func TestSecureStorage(t *testing.T) {
	store, err := NewSecure()
	if err != nil {
		t.Fatalf("NewSecure failed: %v", err)
	}

	account := types.Account{
		Name:   "test-secure",
		Secret: "JBSWY3DPEHPK3PXP",
	}

	err = store.SaveAccount(account)
	if err != nil {
		t.Fatalf("SaveAccount failed: %v", err)
	}

	retrieved, err := store.LoadAccount("test-secure")
	if err != nil {
		t.Fatalf("LoadAccount failed: %v", err)
	}

	if retrieved.Name != account.Name {
		t.Errorf("Expected name %s, got %s", account.Name, retrieved.Name)
	}

	if retrieved.Secret != account.Secret {
		t.Errorf("Expected secret %s, got %s", account.Secret, retrieved.Secret)
	}

	accounts, err := store.ListAccounts()
	if err != nil {
		t.Fatalf("ListAccounts failed: %v", err)
	}

	found := false
	for _, name := range accounts {
		if name == "test-secure" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Account not found in list")
	}

	err = store.DeleteAccount("test-secure")
	if err != nil {
		t.Fatalf("DeleteAccount failed: %v", err)
	}

	_, err = store.LoadAccount("test-secure")
	if err == nil {
		t.Error("Expected error when loading deleted account")
	}
}
