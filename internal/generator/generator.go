package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
		_ = file.Close()
	}
	
	return nil
}

// CreateClaudeMD creates a CLAUDE.md file in the root directory
func CreateClaudeMD(preset string, directories []string) error {
	var content string
	
	switch preset {
	case "claude-basic":
		content = generateClaudeBasicContent(directories)
	case "web-app":
		content = generateWebAppContent(directories)
	case "api-server":
		content = generateAPIServerContent(directories)
	default:
		content = generateDefaultContent(directories)
	}
	
	// Create CLAUDE.md file
	file, err := os.Create("CLAUDE.md")
	if err != nil {
		return fmt.Errorf("failed to create CLAUDE.md: %w", err)
	}
	defer func() { _ = file.Close() }()
	
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write CLAUDE.md: %w", err)
	}
	
	return nil
}

// generateClaudeBasicContent generates content for claude-basic preset
func generateClaudeBasicContent(directories []string) string {
	return `# Claude Code Project

## Project Overview
This is a Claude Code project with basic structure for prompts and memory.

## Directory Structure
` + generateDirectoryList(directories) + `

## Development Guidelines
- Use the prompts directory for storing conversation prompts
- Use the memory directory for storing context and instructions
- Keep prompts focused and specific to their purpose

## Usage
1. Add your prompts to the prompts/ directory
2. Store important context in memory/instructions/
3. Log interactions in memory/logs/
`
}

// generateWebAppContent generates content for web-app preset
func generateWebAppContent(directories []string) string {
	return `# Web Application Project

## Project Overview
This is a web application project structured for Claude Code development.

## Directory Structure
` + generateDirectoryList(directories) + `

## Tech Stack
- Frontend: [Specify framework]
- Backend: [Specify if applicable]
- Database: [Specify if applicable]

## Development Guidelines
- Component prompts go in prompts/components/
- Feature-specific prompts in prompts/features/
- API-related prompts in prompts/api/
- Keep architectural decisions in memory/architecture/

## Getting Started
1. Install dependencies: [Add command]
2. Run development server: [Add command]
3. Build for production: [Add command]
`
}

// generateAPIServerContent generates content for api-server preset
func generateAPIServerContent(directories []string) string {
	return `# API Server Project

## Project Overview
This is an API server project structured for Claude Code development.

## Directory Structure
` + generateDirectoryList(directories) + `

## Tech Stack
- Language: [Specify language]
- Framework: [Specify framework]
- Database: [Specify database]

## API Documentation
- Endpoint prompts in prompts/endpoints/
- Model definitions in prompts/models/
- Database queries in prompts/database/
- API design in memory/api-design/

## Development Guidelines
- Follow RESTful conventions
- Document all endpoints
- Include error handling patterns

## Getting Started
1. Set up database: [Add instructions]
2. Install dependencies: [Add command]
3. Run server: [Add command]
`
}

// generateDefaultContent generates default content when no preset is used
func generateDefaultContent(directories []string) string {
	return `# Claude Code Project

## Project Overview
[Add your project description here]

## Directory Structure
` + generateDirectoryList(directories) + `

## Development Guidelines
[Add your development guidelines here]

## Usage
[Add usage instructions here]
`
}

// generateDirectoryList formats directories into a markdown list
func generateDirectoryList(directories []string) string {
	var result strings.Builder
	result.WriteString("```\n")
	
	// Sort and format directories
	for _, dir := range directories {
		result.WriteString(dir + "/\n")
	}
	
	result.WriteString("```")
	return result.String()
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
	_ = file.Close()
	
	return nil
}