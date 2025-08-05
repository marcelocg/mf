package totp

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	secret := "JBSWY3DPEHPK3PXP"

	token, err := GenerateToken(secret)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	if len(token) != 6 {
		t.Errorf("Expected token length 6, got %d", len(token))
	}

	for _, char := range token {
		if char < '0' || char > '9' {
			t.Errorf("Token should contain only digits, got: %s", token)
			break
		}
	}
}

func TestGenerateTokenWithInvalidSecret(t *testing.T) {
	invalidSecret := "invalid-secret"

	_, err := GenerateToken(invalidSecret)
	if err == nil {
		t.Error("Expected error when generating token with invalid secret")
	}
}

func TestValidateSecret(t *testing.T) {
	validSecret := "JBSWY3DPEHPK3PXP"
	err := ValidateSecret(validSecret)
	if err != nil {
		t.Errorf("ValidateSecret failed for valid secret: %v", err)
	}
}

func TestValidateSecretInvalid(t *testing.T) {
	invalidSecret := "invalid-secret"
	err := ValidateSecret(invalidSecret)
	if err == nil {
		t.Error("Expected error when validating invalid secret")
	}
}

func TestGenerateTokenConsistency(t *testing.T) {
	secret := "JBSWY3DPEHPK3PXP"

	token1, err1 := GenerateToken(secret)
	if err1 != nil {
		t.Fatalf("First GenerateToken failed: %v", err1)
	}

	token2, err2 := GenerateToken(secret)
	if err2 != nil {
		t.Fatalf("Second GenerateToken failed: %v", err2)
	}

	if token1 != token2 {
		t.Errorf("Tokens should be the same within the same time window, got %s and %s", token1, token2)
	}
}
