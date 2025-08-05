# Release Process

This document describes how to create and publish releases for the MF project.

## Automated Release Process

Releases are automated using GitHub Actions. When a version tag is pushed, the system automatically:

1. Builds binaries for all supported platforms
2. Runs tests to ensure quality
3. Creates SHA256 checksums
4. Publishes a GitHub Release with all artifacts

## Creating a Release

### Method 1: Using the Release Script (Recommended)

```bash
# Make sure you're on the main branch with latest changes
git checkout main
git pull origin main

# Run the release script
./scripts/release.sh 1.0.0
```

The script will:

- Validate the version format
- Check that the working directory is clean
- Run tests and build locally
- Create and push the version tag
- Trigger the automated GitHub Actions workflow

### Method 2: Manual Process

```bash
# 1. Update version information
# Edit CHANGELOG.md with release notes

# 2. Create and push tag
git tag -a v1.0.0 -m "Release version 1.0.0"
git push origin v1.0.0

# 3. GitHub Actions will handle the rest automatically
```

## Pre-Release Checklist

Before creating a release, ensure:

- [ ] All tests pass: `make test`
- [ ] Code is properly formatted: `make fmt`
- [ ] No linting issues: `make lint`
- [ ] CHANGELOG.md is updated with release notes
- [ ] Version number follows [Semantic Versioning](https://semver.org/)
- [ ] README.md is up to date
- [ ] All documentation is current

## Version Numbering

We follow [Semantic Versioning](https://semver.org/):

- **MAJOR** version for incompatible API changes
- **MINOR** version for backwards-compatible functionality additions
- **PATCH** version for backwards-compatible bug fixes

Examples:

- `1.0.0` - Major release
- `1.1.0` - Minor release (new features)
- `1.1.1` - Patch release (bug fixes)

## Release Notes

### Automatic Generation

GitHub Actions automatically generates release notes based on:

- Commit messages since the last release
- Pull request titles and descriptions
- Issues that were closed

### Manual Enhancement

After the automated release is created:

1. Go to the [Releases page](../../releases)
2. Click "Edit" on the new release
3. Enhance the automatically generated notes
4. Use the template from `.github/RELEASE_TEMPLATE.md`
5. Categorize changes (Added, Changed, Fixed, Security)
6. Add upgrade instructions if needed

## Supported Platforms

Current release targets:

- **Linux**: `mf-linux-amd64` (x86_64)
- **Windows**: `mf-windows-amd64.exe` (x86_64)
- **macOS**:
  - `mf-macos-amd64` (Intel)
  - `mf-macos-arm64` (Apple Silicon)

## Build Process

The automated build process:

1. **Environment**: Ubuntu latest with Go 1.21
2. **Testing**: Full test suite must pass
3. **Building**: Cross-compilation for all platforms
4. **Verification**: Checksums are generated for all binaries
5. **Publishing**: Artifacts are attached to the GitHub Release

## Post-Release Tasks

After a successful release:

- [ ] Verify all binaries download correctly
- [ ] Test installation on different platforms
- [ ] Update any external documentation
- [ ] Announce the release (if significant)
- [ ] Monitor for issues or feedback

## Troubleshooting

### Build Failures

If the GitHub Actions build fails:

1. Check the [Actions page](../../actions) for detailed logs
2. Fix any issues in the code
3. Push fixes to the main branch
4. Delete and recreate the tag if necessary:

   ```bash
   git tag -d v1.0.0
   git push --delete origin v1.0.0
   git tag -a v1.0.0 -m "Release version 1.0.0"
   git push origin v1.0.0
   ```

### Missing Binaries

If some platform binaries are missing:

1. Check the build matrix in `.github/workflows/release.yml`
2. Ensure all target platforms are included
3. Verify cross-compilation works locally: `make build-all`

## Security Considerations

- Release binaries are built in a controlled GitHub Actions environment
- All releases include SHA256 checksums for verification
- Tags are signed when possible (configure GPG keys)
- Never include secrets or sensitive data in releases

## Emergency Releases

For critical security fixes or major bugs:

1. Create a hotfix branch from the affected release tag
2. Apply the minimal necessary fix
3. Follow the normal release process with a patch version
4. Consider marking previous versions as deprecated

## Questions?

If you have questions about the release process:

- Check existing [Issues](../../issues) and [Discussions](../../discussions)
- Review the GitHub Actions logs for any recent releases
- Refer to the [Contributing Guidelines](CONTRIBUTING.md)
