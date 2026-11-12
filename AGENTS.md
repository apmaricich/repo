# AGENTS.md

## Build & Run
- Build: `make build` (outputs `./repo` binary)
- Run: `make run` or `./repo`
- Clean: `make clean`

## Configuration
- Global: `~/.config/repo/config.json` (required)
- Project-local: `.repo.json` in working directory (optional, overrides global)
- Scripts referenced by absolute path in global config, relative path in project config

## Code Structure
- Single Go binary with two source files
- `main.go`: CLI parsing, dispatch, template substitution
- `config.go`: config loading, JSON parsing, path handling

## Testing
- Manual testing via CLI: build and run specific commands
- No automated test suite
- Verify dispatch, template substitution, config overlay with real config files