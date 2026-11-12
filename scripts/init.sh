#!/bin/bash
set -e

# Initialize git repository if not already initialized
if [ ! -d ".git" ]; then
    echo "Initializing git repository..."
    git init
else
    echo "Git repository already exists."
fi

# Create .gitignore if not already present
if [ ! -f ".gitignore" ]; then
    echo "Creating .gitignore..."
    cp /Users/apmaricich/Code/repo/templates/.gitignore .gitignore
else
    echo ".gitignore already exists."
fi

# Initialize task manager if not already initialized
if [ ! -f "tasks.json" ]; then
    echo "Initializing task manager..."
    task init
else
    echo "Task manager already initialized."
fi