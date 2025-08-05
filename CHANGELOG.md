# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.0] - 2025-08-04

### Added
- 🔐 **Secure Storage System**: Hybrid approach using system keychain with encrypted fallback
- 🛡️ **AES-256-GCM Encryption**: Local data encryption for enhanced security
- 🔄 **Auto-migration**: Automatic upgrade from plain text to encrypted storage
- 🖥️ **Cross-platform Keychain Support**:
  - Linux: Secret Service API (gnome-keyring, KWallet)
  - Windows: Windows Credential Manager
  - macOS: Keychain Services
- 🔑 **Machine-specific Key Derivation**: Uses PBKDF2 with machine identifiers
- 📁 **Encrypted File Format**: New `.enc` format for secure local storage
- 🧪 **Comprehensive Test Suite**: Full test coverage for security modules

### Changed
- **BREAKING**: Storage format changed from `.json` to `.enc` (automatic migration provided)
- **Enhanced Security**: All secrets now encrypted at rest
- **Improved File Permissions**: Stricter access controls on configuration files

### Security
- All stored secrets are now encrypted using AES-256-GCM
- Machine-specific encryption keys prevent cross-machine access
- Automatic cleanup of old plain text files after migration
- No longer stores secrets in plain text format

## [1.0.0] - 2025-08-04

### Added
- 🚀 **Initial Release**: Complete TOTP token generator
- 📱 **TOTP Support**: Compatible with Google Authenticator, Authy, and other TOTP apps
- 🖥️ **Cross-platform**: Support for Linux, Windows, and macOS
- 📋 **CLI Commands**:
  - `add`: Add new TOTP accounts
  - `get`: Generate TOTP tokens
  - `list`: List all configured accounts
- 🛠️ **Cobra Framework**: Professional CLI with help system and auto-completion
- 📁 **Configuration Management**: Stores accounts in `~/.config/mf/`
- ✅ **Input Validation**: Validates TOTP secrets before storage
- 🧪 **Test Coverage**: Comprehensive test suite for core functionality
- 📦 **Easy Distribution**: Single binary with no external dependencies

### Technical Details
- Built with Go 1.24+
- Uses `github.com/pquerna/otp` for TOTP generation
- Uses `github.com/spf13/cobra` for CLI framework
- Stores configuration in JSON format
- File permissions set to 0600 for account files
- Directory permissions set to 0700 for config directory