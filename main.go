package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	config, err := loadMergedConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading config: %v\n", err)
		os.Exit(1)
	}

	args := os.Args[1:]

	if len(args) == 0 {
		printHelp(os.Stderr, config)
		os.Exit(1)
	}

	subcommand := args[0]
	remainingArgs := args[1:]

	// Handle --help flag
	if subcommand == "--help" || subcommand == "-h" {
		printHelp(os.Stdout, config)
		os.Exit(0)
	}

	script, ok := config.Scripts[subcommand]
	if !ok {
		fmt.Fprintf(os.Stderr, "error: unknown subcommand: %s\n\n", subcommand)
		printHelp(os.Stderr, config)
		os.Exit(1)
	}

	expandedArgs := expandArgs(script.Args, remainingArgs)
	cmd := exec.Command(script.Command, expandedArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
		fmt.Fprintf(os.Stderr, "error executing %s: %v\n", subcommand, err)
		os.Exit(1)
	}
}

func expandArgs(templateArgs []string, actualArgs []string) []string {
	result := make([]string, 0, len(templateArgs)+len(actualArgs))

	for _, arg := range templateArgs {
		if strings.HasPrefix(arg, "{") && strings.HasSuffix(arg, "}") {
			// This is a placeholder — replace with the corresponding actual arg
			if len(actualArgs) > 0 {
				result = append(result, actualArgs[0])
				actualArgs = actualArgs[1:]
			} else {
				// No actual arg provided, keep placeholder as-is
				result = append(result, arg)
			}
		} else {
			result = append(result, arg)
		}
	}

	// Append any remaining actual args
	result = append(result, actualArgs...)
	return result
}

func printHelp(w *os.File, config *Config) {
	fmt.Fprintf(w, "Usage: repo <subcommand> [args...]\n\n")
	fmt.Fprintf(w, "Available commands:\n")
	names := make([]string, 0, len(config.Scripts))
	for name := range config.Scripts {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		script := config.Scripts[name]
		desc := script.Description
		if desc == "" {
			desc = fmt.Sprintf("%s %s", script.Command, script.Args)
		}
		fmt.Fprintf(w, "  %-20s %s\n", name, desc)
	}
}