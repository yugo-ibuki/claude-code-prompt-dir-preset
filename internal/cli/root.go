package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/claude-code-prompt-dir-preset/internal/generator"
)

var (
	preset string
)

var rootCmd = &cobra.Command{
	Use:   "ccpdp [directories...]",
	Short: "Claude Code Prompt Directory Preset - Create directory structures for Claude Code",
	Long: `ccpdp is a CLI tool for creating directory structures optimized for Claude Code.
	
Examples:
  ccpdp "docs > instructions" "docs > logs"
  ccpdp --preset claude-basic`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if preset != "" {
			// Handle preset mode
			if err := generator.GenerateFromPreset(preset); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Successfully created directories from preset: %s\n", preset)
			return
		}

		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "Error: No directories specified. Use --preset flag or provide directory paths.")
			cmd.Help()
			os.Exit(1)
		}

		// Parse and generate directories
		paths, err := generator.ParsePaths(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing paths: %v\n", err)
			os.Exit(1)
		}

		if err := generator.CreateDirectories(paths); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directories: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Successfully created directories:")
		for _, path := range paths {
			fmt.Printf("  ✓ %s\n", path)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&preset, "preset", "p", "", "Use a preset configuration (e.g., claude-basic, docs, prompts, project-info)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}