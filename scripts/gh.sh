#!/bin/bash
set -e

# Get the remote URL
REMOTE_URL=$(git remote get-url origin 2>/dev/null)
if [ -z "$REMOTE_URL" ]; then
    echo "Error: No git remote 'origin' found."
    exit 1
fi

# Parse GitHub URLs into web page URLs
# Handle HTTPS: https://github.com/owner/repo.git
if [[ "$REMOTE_URL" =~ ^https://github.com/([^/]+/[^/]+) ]]; then
    WEB_URL="https://github.com/${BASH_REMATCH[1]}"
    WEB_URL="${WEB_URL%.git}"
# Handle SSH: git@github.com:owner/repo.git
elif [[ "$REMOTE_URL" =~ ^git@github.com:([^/]+/[^/]+) ]]; then
    WEB_URL="https://github.com/${BASH_REMATCH[1]}"
    WEB_URL="${WEB_URL%.git}"
# Handle other GitHub URL formats (fallback)
else
    echo "Error: Could not parse GitHub URL: $REMOTE_URL"
    exit 1
fi

echo "Opening: $WEB_URL"

# Open in browser (macOS)
if [[ "$(uname)" == "Darwin" ]]; then
    open "$WEB_URL"
# Open in browser (Linux)
elif command -v xdg-open >/dev/null 2>&1; then
    xdg-open "$WEB_URL"
else
    echo "URL: $WEB_URL"
fi