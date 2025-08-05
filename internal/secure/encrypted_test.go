package secure

import (
	"mf/internal/types"
	"testing"
)

func TestEncryptedStorage(t *testing.T) {
	provider := &EncryptedProvider{}

	if !provider.IsAvailable() {
		t.Fatal("EncryptedProvider should always be available")
	}

	store, err := provider.GetStorage()
	if err != nil {
		t.Fatalf("Failed to get encrypted storage: %v", err)
	}

	account := types.Account{
		Name:   "test-encrypted",
		Secret: "JBSWY3DPEHPK3PXP",
	}

	err = store.Store(account)
	if err != nil {
		t.Fatalf("Store failed: %v", err)
	}

	retrieved, err := store.Retrieve("test-encrypted")
	if err != nil {
		t.Fatalf("Retrieve failed: %v", err)
	}

	if retrieved.Name != account.Name {
		t.Errorf("Expected name %s, got %s", account.Name, retrieved.Name)
	}

	if retrieved.Secret != account.Secret {
		t.Errorf("Expected secret %s, got %s", account.Secret, retrieved.Secret)
	}

	accounts, err := store.List()
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}

	found := false
	for _, name := range accounts {
		if name == "test-encrypted" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Account not found in list")
	}

	err = store.Delete("test-encrypted")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	_, err = store.Retrieve("test-encrypted")
	if err == nil {
		t.Error("Expected error when retrieving deleted account")
	}
}

func TestEncryptDecryptRoundTrip(t *testing.T) {
	provider := &EncryptedProvider{}
	store, err := provider.GetStorage()
	if err != nil {
		t.Fatalf("Failed to get encrypted storage: %v", err)
	}

	encStorage, ok := store.(*EncryptedStorage)
	if !ok {
		t.Fatal("Expected EncryptedStorage type")
	}

	originalData := []byte("test data for encryption")

	encrypted, err := encStorage.encrypt(originalData)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	if string(encrypted) == string(originalData) {
		t.Error("Encrypted data should be different from original")
	}

	decrypted, err := encStorage.decrypt(encrypted)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if string(decrypted) != string(originalData) {
		t.Errorf("Decrypted data doesn't match original. Expected: %s, Got: %s", originalData, decrypted)
	}
}
