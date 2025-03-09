package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Repository struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type TAG struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	Subprojects []Repository `yaml:"subprojects"`
}

type TAGsConfig struct {
	TAGs []TAG `yaml:"tags"`
}

func main() {
	configPath := filepath.Join("..", "tags.yaml") // Adjusted for new repo structure
	outputDir := filepath.Join("..", "groups")      // Adjust as needed

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read tags.yaml: %v", err)
	}

	var config TAGsConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	for _, tag := range config.TAGs {
		generateTAGReadme(tag, outputDir)
	}
}

func generateTAGReadme(tag TAG, outputDir string) {
	filePath := filepath.Join(outputDir, fmt.Sprintf("TAG-%s.md", tag.Name))
	fileContent := fmt.Sprintf("# %s\n\n%s\n\n## Subprojects\n", tag.Name, tag.Description)

	for _, subproject := range tag.Subprojects {
		fileContent += fmt.Sprintf("- **%s**: %s\n", subproject.Name, subproject.Description)
	}

	if err := os.WriteFile(filePath, []byte(fileContent), 0644); err != nil {
		log.Fatalf("Failed to write %s: %v", filePath, err)
	}

	log.Printf("Generated: %s", filePath)
}
