package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateDirectories creates the specified directories with .gitkeep files
func CreateDirectories(paths []string) error {
	for _, path := range paths {
		// Create directory with all parent directories
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
		
		// Create .gitkeep file in the directory
		gitkeepPath := filepath.Join(path, ".gitkeep")
		file, err := os.Create(gitkeepPath)
		if err != nil {
			return fmt.Errorf("failed to create .gitkeep in %s: %w", path, err)
		}
		file.Close()
	}
	
	return nil
}

// CreateSingleDirectory creates a single directory with .gitkeep
func CreateSingleDirectory(path string) error {
	// Create directory
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}
	
	// Create .gitkeep file
	gitkeepPath := filepath.Join(path, ".gitkeep")
	file, err := os.Create(gitkeepPath)
	if err != nil {
		return fmt.Errorf("failed to create .gitkeep in %s: %w", path, err)
	}
	file.Close()
	
	return nil
}