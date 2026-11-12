#!/bin/bash
set -e

if [ -z "$1" ]; then
    echo "Usage: repo all \"<command>\""
    echo "Example: repo all \"git status\""
    exit 1
fi

COMMAND="$1"
FOUND=0

for dir in */; do
    [ -d "$dir/.git" ] || continue
    FOUND=1
    
    echo ""
    echo "=== ${dir%/} ==="
    cd "$dir"
    eval "$COMMAND"
    cd ..
done

if [ $FOUND -eq 0 ]; then
    echo "No git repositories found in subdirectories."
fi