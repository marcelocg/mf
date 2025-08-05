# Scripts

This directory contains utility scripts for the MF project.

## release.sh

Automated script for creating new releases.

### Usage

```bash
./scripts/release.sh <version>
```

### Examples

```bash
# Create version 1.0.0
./scripts/release.sh 1.0.0

# Create version 2.1.3
./scripts/release.sh 2.1.3
```

### What it does

1. Validates the version format (semantic versioning)
2. Checks that the working directory is clean
3. Verifies you're in a git repository
4. Runs the full test and build process (`make release`)
5. Creates and pushes a git tag
6. Triggers GitHub Actions to build and publish the release

### Requirements

- Clean git working directory (no uncommitted changes)
- All tests must pass
- Valid semantic version number (e.g., 1.2.3)

### Troubleshooting

**"Working directory is not clean"**
- Commit or stash any pending changes before running the script

**"Tag already exists"**
- The version tag already exists. Choose a different version number or delete the existing tag

**"Invalid version format"**
- Use semantic versioning format: MAJOR.MINOR.PATCH (e.g., 1.0.0, 2.1.5)

### Manual Process

If you prefer to create releases manually:

```bash
# Run tests and build
make release

# Create and push tag
git tag -a v1.0.0 -m "Release version 1.0.0"
git push origin v1.0.0
```