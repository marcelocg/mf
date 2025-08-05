package secure

import (
	"testing"
)

func TestGetMachineKey(t *testing.T) {
	key, err := GetMachineKey()
	if err != nil {
		t.Fatalf("GetMachineKey failed: %v", err)
	}

	if len(key) != 16 {
		t.Errorf("Expected key length 16, got %d", len(key))
	}

	key2, err := GetMachineKey()
	if err != nil {
		t.Fatalf("Second GetMachineKey failed: %v", err)
	}

	if string(key) != string(key2) {
		t.Error("Machine key should be consistent between calls")
	}
}
