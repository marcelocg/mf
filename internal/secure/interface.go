package secure

import "mf/internal/types"

type SecureStorage interface {
	Store(account types.Account) error
	Retrieve(name string) (*types.Account, error)
	List() ([]string, error)
	Delete(name string) error
}

type SecureStorageProvider interface {
	IsAvailable() bool
	GetStorage() (SecureStorage, error)
}
