# Contributing to MF

Thank you for your interest in contributing to MF! This document provides guidelines for contributing to this project.

## Code of Conduct

This project adheres to a code of conduct. By participating, you are expected to uphold this code. Please report unacceptable behavior.

## How to Contribute

### Reporting Bugs

1. **Check existing issues** - Search through existing issues to see if the bug has already been reported
2. **Create a detailed issue** - Include:
   - Clear title and description
   - Steps to reproduce
   - Expected vs actual behavior
   - OS and version information
   - MF version (`mf --version`)

### Suggesting Features

1. **Check existing issues** - See if the feature has been suggested before
2. **Create a feature request** with:
   - Clear description of the feature
   - Use cases and benefits
   - Possible implementation approach

### Development Setup

#### Prerequisites

- Go 1.19 or later
- Git

#### Getting Started

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/mf.git
   cd mf
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Run tests to ensure everything works:
   ```bash
   make test
   ```

#### Development Workflow

1. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes following the coding standards below

3. Run the development workflow:
   ```bash
   make dev  # formats, lints, tests, and builds
   ```

4. Commit your changes:
   ```bash
   git add .
   git commit -m "descriptive commit message"
   ```

5. Push to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```

6. Create a Pull Request

## Coding Standards

### Go Style

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting (included in `make dev`)
- Use `go vet` for linting (included in `make dev`)

### Code Organization

- Keep functions small and focused
- Use meaningful variable and function names
- Add comments for complex logic
- Follow the existing package structure:
  - `cmd/` - CLI commands
  - `internal/types/` - Shared types
  - `internal/storage/` - Storage implementations
  - `internal/secure/` - Security-related code
  - `internal/totp/` - TOTP generation

### Testing

- Write tests for new functionality
- Maintain or improve test coverage
- Use table-driven tests where appropriate
- Test both success and error cases

### Security

- Never expose secrets in logs or error messages
- Follow secure coding practices
- Be cautious with file permissions
- Consider cross-platform security implications

## Pull Request Guidelines

### Before Submitting

- [ ] Run `make dev` and ensure all checks pass
- [ ] Add tests for new functionality
- [ ] Update documentation if needed
- [ ] Update CHANGELOG.md if applicable

### PR Description

Include:
- **What** - Brief description of changes
- **Why** - Reason for the changes
- **How** - Implementation approach
- **Testing** - How you tested the changes

### Review Process

1. Automated checks must pass
2. Code review by maintainer(s)
3. Address any feedback
4. Approval and merge

## Security Issues

For security-related issues, please do **NOT** create a public issue. Instead:
1. Email the maintainers directly
2. Provide detailed information about the vulnerability
3. Allow time for a fix before public disclosure

## Questions?

- Check existing issues and documentation
- Create a discussion for general questions
- Ensure your question hasn't been answered before

## Development Commands

```bash
# Build for current platform
make build

# Run tests
make test

# Run tests with coverage
make test-coverage

# Format and lint code
make fmt
make lint

# Development workflow (fmt + lint + test + build)
make dev

# Build for all platforms
make build-all

# Clean build artifacts
make clean
```

## Thank You!

Your contributions make this project better for everyone. Thank you for taking the time to contribute!