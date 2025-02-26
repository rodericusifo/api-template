#!/bin/bash
# Copyright 2025 Rodericus Ifo Krista
# SPDX-License-Identifier: MIT

# Git Hooks Setup Script
# This script installs the project's git hooks for all developers

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

print_error() {
    echo -e "${RED}ERROR: $1${NC}" >&2
}

print_success() {
    echo -e "${GREEN}SUCCESS: $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}WARNING: $1${NC}"
}

print_info() {
    echo -e "${BLUE}INFO: $1${NC}"
}

print_info "Setting up Git hooks for API Template project..."

# Check if we're in a git repository
if [ ! -d ".git" ]; then
    print_error "This script must be run from the root of the git repository"
    exit 1
fi

# Check if hooks directory exists
if [ ! -d "script/hooks" ]; then
    print_error "Hooks directory not found at script/hooks/"
    print_error "Please ensure you're running this from the project root"
    exit 1
fi

# Create backup of existing hooks if they exist
hooks_dir=".git/hooks"
backup_dir=".git/hooks.backup.$(date +%Y%m%d_%H%M%S)"

if [ -f "$hooks_dir/pre-commit" ] && [ ! -L "$hooks_dir/pre-commit" ]; then
    print_info "Backing up existing hooks to $backup_dir"
    mkdir -p "$backup_dir"
    cp "$hooks_dir"/pre-* "$backup_dir/" 2>/dev/null || true
fi

# Install pre-commit hook
print_info "Installing pre-commit hook..."
cp "script/hooks/pre-commit" "$hooks_dir/pre-commit"
chmod +x "$hooks_dir/pre-commit"

# Install commit-msg hook
print_info "Installing commit-msg hook..."
cp "script/hooks/commit-msg" "$hooks_dir/commit-msg"
chmod +x "$hooks_dir/commit-msg"

# Install pre-push hook
print_info "Installing pre-push hook..."
cp "script/hooks/pre-push" "$hooks_dir/pre-push"
chmod +x "$hooks_dir/pre-push"

# Verify installation
if [ -x "$hooks_dir/pre-commit" ] && [ -x "$hooks_dir/commit-msg" ] && [ -x "$hooks_dir/pre-push" ]; then
    print_success "Git hooks installed successfully"
else
    print_error "Failed to install one or more hooks"
    exit 1
fi

print_success "Git hooks setup completed!"
print_info ""
print_info "The following hooks are now active:"
print_info "• pre-commit: Validates single service per commit"
print_info "• commit-msg: Validates conventional commit message format"
print_info "• pre-push: Validates version changes and creates changelog"
print_info ""
print_info "Hook features:"
print_info "• Conventional commit format validation (feat, fix, docs, etc.)"
print_info "• Ensures only one service is modified per commit"
print_info "• Validates semantic versioning rules"
print_info "• Automatically creates/updates service-specific CHANGELOG.md"
print_info "• Prevents pushes with invalid version increments"
print_info ""
print_info "Example workflow:"
print_info "  1. Edit files in application/service-name/"
print_info "  2. git add ."
print_info "  3. git commit -m \"feat(service-name): add new feature\""
print_info "  4. (Optional) Update application/service-name/VERSION"
print_info "  5. git push"
print_info "  4. git commit -m 'Your message'"
print_info ""
print_warning "Remember: Only modify one service per commit!"

exit 0
