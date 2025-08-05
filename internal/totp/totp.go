package totp

import (
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
)

func GenerateToken(secret string) (string, error) {
	token, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", fmt.Errorf("failed to generate TOTP token: %w", err)
	}
	return token, nil
}

func ValidateSecret(secret string) error {
	_, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return fmt.Errorf("invalid secret: %w", err)
	}
	return nil
}
