package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	appVersion   = "dev"
	appBuildTime = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "mf",
	Short: "MF - Multi-Factor Authentication Token Generator",
	Long: `MF is a secure, cross-platform command-line application for generating 
TOTP (Time-based One-Time Password) tokens for multi-factor authentication.

Features:
• Secure storage using system keychain with encrypted fallback
• Cross-platform support (Linux, Windows, macOS)
• Script-friendly with no password prompts
• Automatic migration from plain text to encrypted storage`,
	Version: appVersion,
}

func Execute() error {
	return rootCmd.Execute()
}

func SetVersion(version, buildTime string) {
	appVersion = version
	appBuildTime = buildTime
	rootCmd.Version = version
	rootCmd.SetVersionTemplate(fmt.Sprintf("MF version %s (built %s)\n", version, buildTime))
}

func init() {
	cobra.OnInitialize()
}
