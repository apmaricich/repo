#!/bin/bash
set -e

branch=$(git rev-parse --abbrev-ref HEAD)

if [ "$branch" = "HEAD" ]; then
    echo "error: detached HEAD, cannot push" >&2
    exit 1
fi

echo "pushing $branch to origin"
git push origin "$branch"
