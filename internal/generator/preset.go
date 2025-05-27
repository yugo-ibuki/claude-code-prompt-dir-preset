package generator

import (
	"fmt"
)

// Preset represents a directory structure preset
type Preset struct {
	Name        string
	Description string
	Directories []string
}

// GetPresets returns all available presets
func GetPresets() map[string]Preset {
	return map[string]Preset{
		"claude-basic": {
			Name:        "claude-basic",
			Description: "Claude Code basic configuration",
			Directories: []string{
				".claude/instructions",
				".claude/logs",
				"docs/specifications",
				"docs/references",
			},
		},
		"docs": {
			Name:        "docs",
			Description: "Documentation management",
			Directories: []string{
				"docs/instructions",
				"docs/logs",
				"docs/specifications",
				"docs/references",
				"docs/decisions",
			},
		},
		"prompts": {
			Name:        "prompts",
			Description: "Prompt management",
			Directories: []string{
				"prompts/system",
				"prompts/user",
				"prompts/examples",
				"prompts/templates",
			},
		},
		"project-info": {
			Name:        "project-info",
			Description: "Project information management",
			Directories: []string{
				"project/requirements",
				"project/architecture",
				"project/decisions",
				"project/meetings",
			},
		},
	}
}

// GenerateFromPreset creates directories based on a preset
func GenerateFromPreset(presetName string) error {
	presets := GetPresets()
	
	preset, exists := presets[presetName]
	if !exists {
		availablePresets := []string{}
		for name := range presets {
			availablePresets = append(availablePresets, name)
		}
		return fmt.Errorf("unknown preset: %s. Available presets: %v", presetName, availablePresets)
	}
	
	// Create all directories in the preset
	for _, dir := range preset.Directories {
		if err := CreateSingleDirectory(dir); err != nil {
			return err
		}
	}
	
	return nil
}