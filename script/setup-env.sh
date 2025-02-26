#!/bin/bash
# Copyright 2025 Rodericus Ifo Krista
# SPDX-License-Identifier: MIT

# Script to copy all .env.example files to corresponding <ENV>.env files
# Usage: ./script/setup-env.sh <ENV>
# Example: ./script/setup-env.sh dev

set -e

# Check if environment argument is provided
if [ $# -eq 0 ]; then
    echo "❌ Error: Environment name is required"
    echo "Usage: $0 <ENV>"
    echo "Example: $0 dev"
    exit 1
fi

ENV=$1

# Validate environment name (only allow alphanumeric and dash/underscore)
if [[ ! "$ENV" =~ ^[a-zA-Z0-9_-]+$ ]]; then
    echo "❌ Error: Invalid environment name. Only alphanumeric characters, dashes, and underscores are allowed."
    exit 1
fi

echo "🚀 Setting up environment files for: $ENV"

# Find all .env.example files and copy them to <ENV>.env
find . -name ".env.example" -type f | while read -r example_file; do
    # Get the directory containing the .env.example file
    dir=$(dirname "$example_file")
    
    # Create the target filename
    target_file="$dir/$ENV.env"
    
    # Copy the file
    cp "$example_file" "$target_file"
    
    echo "✅ Copied: $example_file → $target_file"
done

echo "🎉 Environment setup completed for: $ENV"
echo ""
echo "📝 Next steps:"
echo "1. Update the generated $ENV.env files with your actual configuration values"
echo "2. Make sure to never commit these files to version control"
echo "3. Add $ENV.env to .gitignore if not already present"
