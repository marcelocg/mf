#!/bin/bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if version is provided
if [ $# -eq 0 ]; then
    print_error "Please provide a version number"
    echo "Usage: $0 <version>"
    echo "Example: $0 1.0.0"
    exit 1
fi

VERSION=$1
TAG="v${VERSION}"

# Validate version format (basic semver check)
if [[ ! $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    print_error "Invalid version format. Use semantic versioning (e.g., 1.0.0)"
    exit 1
fi

print_status "Creating release for version ${VERSION}"

# Check if we're in a git repository
if [ ! -d .git ]; then
    print_error "Not in a git repository"
    exit 1
fi

# Check if working directory is clean
if [ -n "$(git status --porcelain)" ]; then
    print_error "Working directory is not clean. Please commit or stash changes."
    exit 1
fi

# Check if tag already exists
if git rev-parse "$TAG" >/dev/null 2>&1; then
    print_error "Tag $TAG already exists"
    exit 1
fi

# Run tests and build
print_status "Running tests and building..."
make release

# Update CHANGELOG if it exists
if [ -f CHANGELOG.md ]; then
    print_warning "Don't forget to update CHANGELOG.md with release notes"
fi

# Create and push tag
print_status "Creating tag ${TAG}"
git tag -a "$TAG" -m "Release version ${VERSION}"

print_status "Pushing tag to origin"
git push origin "$TAG"

print_success "Release ${VERSION} has been tagged and pushed!"
print_status "GitHub Actions will automatically create the release with binaries."
print_status "Check https://github.com/marcelocg/mf/actions for build status."

# Display next steps
echo ""
echo "Next steps:"
echo "1. Check GitHub Actions for build status"
echo "2. Edit the release notes on GitHub if needed"
echo "3. Update README or documentation if necessary"
echo "4. Announce the release!"