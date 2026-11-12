#!/bin/bash
set -e

echo "Building repo..."
go build -o repo .
echo "Build complete: ./repo"