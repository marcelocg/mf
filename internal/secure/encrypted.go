package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/crypto/pbkdf2"
	"mf/internal/types"
)

type EncryptedStorage struct {
	configDir string
	key       []byte
}

type EncryptedProvider struct{}

func (p *EncryptedProvider) IsAvailable() bool {
	return true
}

func (p *EncryptedProvider) GetStorage() (SecureStorage, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "mf")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %w", err)
	}

	machineKey, err := GetMachineKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get machine key: %w", err)
	}

	key := pbkdf2.Key(machineKey, []byte("mf-salt"), 10000, 32, sha256.New)

	return &EncryptedStorage{
		configDir: configDir,
		key:       key,
	}, nil
}

func (e *EncryptedStorage) Store(account types.Account) error {
	data, err := json.Marshal(account)
	if err != nil {
		return fmt.Errorf("failed to marshal account: %w", err)
	}

	encryptedData, err := e.encrypt(data)
	if err != nil {
		return fmt.Errorf("failed to encrypt account data: %w", err)
	}

	filename := filepath.Join(e.configDir, account.Name+".enc")
	if err := os.WriteFile(filename, encryptedData, 0600); err != nil {
		return fmt.Errorf("failed to write encrypted account file: %w", err)
	}

	e.removeOldFormat(account.Name)
	return nil
}

func (e *EncryptedStorage) Retrieve(name string) (*types.Account, error) {
	filename := filepath.Join(e.configDir, name+".enc")
	encryptedData, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			if account, legacyErr := e.tryLoadLegacy(name); legacyErr == nil {
				e.Store(*account)
				return account, nil
			}
			return nil, fmt.Errorf("account '%s' not found", name)
		}
		return nil, fmt.Errorf("failed to read encrypted account file: %w", err)
	}

	data, err := e.decrypt(encryptedData)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt account data: %w", err)
	}

	var account types.Account
	if err := json.Unmarshal(data, &account); err != nil {
		return nil, fmt.Errorf("failed to unmarshal account data: %w", err)
	}

	return &account, nil
}

func (e *EncryptedStorage) List() ([]string, error) {
	entries, err := os.ReadDir(e.configDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read config directory: %w", err)
	}

	var accounts []string
	for _, entry := range entries {
		if !entry.IsDir() {
			name := entry.Name()
			if filepath.Ext(name) == ".enc" {
				accountName := name[:len(name)-4]
				accounts = append(accounts, accountName)
			} else if filepath.Ext(name) == ".json" {
				accountName := name[:len(name)-5]
				accounts = append(accounts, accountName)
			}
		}
	}

	return accounts, nil
}

func (e *EncryptedStorage) Delete(name string) error {
	filename := filepath.Join(e.configDir, name+".enc")
	if err := os.Remove(filename); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("account '%s' not found", name)
		}
		return fmt.Errorf("failed to delete encrypted account file: %w", err)
	}

	e.removeOldFormat(name)
	return nil
}

func (e *EncryptedStorage) encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func (e *EncryptedStorage) decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(data) < gcm.NonceSize() {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func (e *EncryptedStorage) tryLoadLegacy(name string) (*types.Account, error) {
	filename := filepath.Join(e.configDir, name+".json")
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var account types.Account
	if err := json.Unmarshal(data, &account); err != nil {
		return nil, err
	}

	return &account, nil
}

func (e *EncryptedStorage) removeOldFormat(name string) {
	oldFile := filepath.Join(e.configDir, name+".json")
	os.Remove(oldFile)
}
