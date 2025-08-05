# Release Notes Template

## 🚀 What's New

<!-- Describe the main features and improvements in this release -->

## 🔧 Changes

<!-- List all changes, organized by category -->

### Added
- 

### Changed
- 

### Fixed
- 

### Security
- 

## 📦 Downloads

Download the appropriate binary for your platform:

- **Linux (64-bit)**: `mf-linux-amd64`
- **Windows (64-bit)**: `mf-windows-amd64.exe`
- **macOS Intel**: `mf-macos-amd64`
- **macOS Apple Silicon**: `mf-macos-arm64`

## 🔐 Checksums

Verify your download with SHA256 checksums (see `checksums.txt`):

```bash
# Linux/macOS
sha256sum -c checksums.txt

# Windows (PowerShell)
Get-FileHash -Algorithm SHA256 mf-windows-amd64.exe
```

## 📋 Installation

### Quick Install
```bash
# Linux/macOS
curl -L https://github.com/marcelocg/mf/releases/download/vX.X.X/mf-linux-amd64 -o mf
chmod +x mf
sudo mv mf /usr/local/bin/

# Or use your package manager (if available)
```

### Verify Installation
```bash
mf --version
```

## 🔄 Upgrade Notes

<!-- Any special instructions for upgrading -->

## 🐛 Known Issues

<!-- List any known issues in this release -->

## 📚 Documentation

- [README](../README.md)
- [Contributing Guidelines](../CONTRIBUTING.md)
- [Changelog](../CHANGELOG.md)