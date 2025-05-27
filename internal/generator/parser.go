package generator

import (
	"fmt"
	"path/filepath"
	"strings"
)

// ParsePaths converts path strings with ">" separator into actual file paths
// Example: "docs > instructions" -> "docs/instructions"
func ParsePaths(args []string) ([]string, error) {
	var paths []string
	
	for _, arg := range args {
		// Split by ">" and trim spaces
		parts := strings.Split(arg, ">")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		
		// Validate parts
		for _, part := range parts {
			if part == "" {
				return nil, fmt.Errorf("invalid path format: %s", arg)
			}
			// Check for invalid characters
			if strings.ContainsAny(part, `\/:*?"<>|`) {
				return nil, fmt.Errorf("invalid directory name: %s", part)
			}
		}
		
		// Join parts with proper path separator
		path := filepath.Join(parts...)
		paths = append(paths, path)
	}
	
	return paths, nil
}