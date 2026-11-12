# repo

A CLI tool that launches configured scripts via subcommands. Perfect for wrapping repetitive development tasks across projects.

## Features

- **Command Hub**: Map subcommands to scripts via a simple JSON config
- **Two-Level Configuration**: Global settings + project-local overrides
- **Template Arguments**: Pass arguments to scripts using `{arg}` placeholders
- **Built-in Commands**: `init`, `all`, `gh` — common dev workflows out of the box
- **Task Integration**: Works with [task](https://github.com/apmaricich/task) for project management

## Build

```bash
make build
```

This creates a `repo` binary in the current directory.

## Configuration

### Global Config
`~/.config/repo/config.json` — defines commands available everywhere.

### Project Config
`.repo.json` in your project root — overrides or adds commands for that project only.

### Command Definition
```json
{
  "scripts": {
    "init": {
      "command": "bash",
      "args": ["/path/to/script.sh"],
      "description": "Initialize project"
    }
  }
}
```

## Built-in Commands

### `repo init`
Initialize a new project with git, task manager, and .gitignore:
```bash
repo init
```
Requires [task](https://github.com/apmaricich/task) installed for task manager initialization.

### `repo all`
Run a command in all git repositories within subdirectories:
```bash
repo all "git status"
```

### `repo gh`
Open the current repository's GitHub page in your browser:
```bash
repo gh
```

### `repo --help`
Show usage and available commands:
```bash
repo --help
```

## Adding Custom Commands

To add a custom command:
1. Write a shell/Node/Python script
2. Add it to your global or project config
3. Use `{arg}` in args for template substitution

Example: Add a `test` command that runs your test suite:
```json
{
  "scripts": {
    "test": {
      "command": "bash",
      "args": ["./scripts/test.sh"],
      "description": "Run test suite"
    }
  }
}
```