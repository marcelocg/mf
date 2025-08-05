# MF - Multi-Factor Authentication Token Generator

A secure, cross-platform command-line application for generating TOTP (Time-based One-Time Password) tokens for multi-factor authentication.

## Features

- 🔐 **Secure Storage**: Uses system keychain when available, with encrypted fallback
- 🚀 **Cross-platform**: Works on Linux, macOS, and Windows
- 🔄 **Script-friendly**: Perfect for automation and scripts (no password prompts)
- 📱 **TOTP Compatible**: Works with any service that uses TOTP (Google, AWS, GitHub, etc.)
- 🛡️ **AES-256 Encryption**: Local data is encrypted using AES-256-GCM
- 🔄 **Auto-migration**: Automatically upgrades from plain text storage to encrypted

## Installation

### Download Binary

Download the latest release from the [releases page](../../releases) for your platform:

- **Linux (64-bit)**: `mf-linux-amd64`
- **Windows (64-bit)**: `mf-windows-amd64.exe`
- **macOS Intel**: `mf-macos-amd64`
- **macOS Apple Silicon**: `mf-macos-arm64`

#### Quick Install (Linux/macOS)

```bash
# Download and install latest version
curl -L https://github.com/marcelocg/mf/releases/latest/download/mf-linux-amd64 -o mf
chmod +x mf
sudo mv mf /usr/local/bin/

# Verify installation
mf --version
```

#### Verify Downloads

All releases include SHA256 checksums in `checksums.txt`:

```bash
# Linux/macOS
sha256sum -c checksums.txt

# Windows (PowerShell)
Get-FileHash -Algorithm SHA256 mf-windows-amd64.exe
```

### Build from Source

```bash
git clone <repository-url>
cd mf
make build
```

#### Build Commands

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Build for specific platforms
make build-linux     # Linux 64-bit
make build-windows   # Windows 64-bit  
make build-macos     # macOS (Intel + Apple Silicon)

# Development workflow (format, lint, test, build)
make dev

# Run tests
make test

# Clean build artifacts
make clean
```

## Usage

### Add an Account

```bash
mf add ACCOUNT_NAME SECRET_KEY
```

Example:

```bash
mf add AWS-DEV 7C2FFYEHYDUKFDYYNMALARRODZ5CXTD2LWOAID2F4KZD63MMH3XWVWNTZLTR7T3X
```

### Generate Token

```bash
mf get ACCOUNT_NAME
```

Example:

```bash
mf get AWS-DEV
# Output: 756815
```

### List All Accounts

```bash
mf list
```

### Help

```bash
mf --help
mf [command] --help
```

## Security

### Storage Locations

- **Linux**: `~/.config/mf/`
- **Windows**: `%USERPROFILE%\.config\mf\`
- **macOS**: `~/.config/mf/`

### Security Features

1. **System Keychain Integration**:
   - Linux: Secret Service API (gnome-keyring, KWallet)
   - Windows: Windows Credential Manager
   - macOS: Keychain Services

2. **Encrypted Fallback**:
   - AES-256-GCM encryption
   - Machine-specific key derivation
   - PBKDF2 key stretching

3. **File Permissions**:
   - Configuration directory: `0700` (owner only)
   - Account files: `0600` (owner read/write only)

## Script Integration

MF is designed to work seamlessly in scripts without user interaction:

```bash
#!/bin/bash
TOKEN=$(mf get AWS-DEV)
aws sts get-caller-identity --token-code $TOKEN
```

```powershell
# PowerShell
$token = mf get AWS-DEV
aws sts get-caller-identity --token-code $token
```

## Examples

### AWS CLI with MFA

```bash
# Add your AWS MFA device
mf add AWS-MFA JBSWY3DPEHPK3PXP...

# Use in AWS CLI
aws sts assume-role \
  --role-arn arn:aws:iam::123456789012:role/MyRole \
  --role-session-name MySession \
  --serial-number arn:aws:iam::123456789012:mfa/user \
  --token-code $(mf get AWS-MFA)
```

### GitHub with MFA

```bash
# Add GitHub TOTP
mf add GITHUB-MFA ABCDEFGH...

# Get token for GitHub
mf get GITHUB-MFA
```

## Configuration

No configuration files are needed. The application works out of the box with secure defaults.

## Building

### Prerequisites

- Go 1.19 or later
- Make (optional, but recommended)

### Using Makefile (Recommended)

```bash
# Development workflow
make dev             # Format, lint, test, and build

# Production builds
make build-all       # Build for all platforms
make build-linux     # Linux only
make build-windows   # Windows only
make build-macos     # macOS only

# Testing and maintenance
make test           # Run tests
make test-coverage  # Run tests with coverage report
make clean          # Clean build artifacts
```

### Manual Build (Alternative)

```bash
# Local build
go build -o mf

# Cross-compilation examples
GOOS=windows GOARCH=amd64 go build -o mf.exe
GOOS=darwin GOARCH=amd64 go build -o mf-macos
GOOS=darwin GOARCH=arm64 go build -o mf-macos-arm64

# Testing
go test ./...
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [go-keyring](https://github.com/zalando/go-keyring) - Cross-platform keychain access
- [otp](https://github.com/pquerna/otp) - TOTP generation
- [crypto](https://pkg.go.dev/golang.org/x/crypto) - Encryption utilities

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Run `make dev` to format, lint, test, and build
6. Submit a pull request

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed contribution guidelines.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Security Considerations

- **Secret Keys**: Never share your secret keys or commit them to version control
- **Backups**: Consider backing up your `~/.config/mf/` directory securely
- **Machine Access**: Anyone with access to your user account can potentially access stored secrets
- **Network**: This application works offline and never transmits your secrets over the network

## Troubleshooting

### Common Issues

1. **"Account not found"**: Make sure you've added the account using `mf add`
2. **Permission errors**: Ensure you have write access to `~/.config/mf/`
3. **Invalid secret**: Verify the secret key is a valid base32-encoded string

### Getting Help

- Check the help: `mf --help`
- Enable verbose output (if implemented): `mf --verbose get ACCOUNT`
- Check file permissions: `ls -la ~/.config/mf/`

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for version history and changes.
