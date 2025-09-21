// Script to add version metadata to all lesson files
// This ensures each lesson file has proper version information for accurate execution
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// VersionMetadata represents version information for a lesson
type VersionMetadata struct {
	GoVersion    string
	MinVersion   string
	IntroducedIn string
	Features     []string
}

// Version mapping for lessons
var versionMapping = map[string]VersionMetadata{
	"1.18": {
		GoVersion:    "1.18",
		MinVersion:   "1.18.0",
		IntroducedIn: "1.18.0",
		Features:     []string{"generics", "type-parameters", "workspace"},
	},
	"1.19": {
		GoVersion:    "1.19",
		MinVersion:   "1.19.0",
		IntroducedIn: "1.19.0",
		Features:     []string{"atomic-types", "memory-arenas"},
	},
	"1.20": {
		GoVersion:    "1.20",
		MinVersion:   "1.20.0",
		IntroducedIn: "1.20.0",
		Features:     []string{"comparable-types", "slice-to-array", "errors-join"},
	},
	"1.21": {
		GoVersion:    "1.21",
		MinVersion:   "1.21.0",
		IntroducedIn: "1.21.0",
		Features:     []string{"builtin-functions", "slices-package", "maps-package"},
	},
	"1.22": {
		GoVersion:    "1.22",
		MinVersion:   "1.22.0",
		IntroducedIn: "1.22.0",
		Features:     []string{"for-range-int", "enhanced-routing", "loop-variables"},
	},
	"1.23": {
		GoVersion:    "1.23",
		MinVersion:   "1.23.0",
		IntroducedIn: "1.23.0",
		Features:     []string{"structured-logging", "iterators", "timer-reset"},
	},
	"1.24": {
		GoVersion:    "1.24",
		MinVersion:   "1.24.0",
		IntroducedIn: "1.24.0",
		Features:     []string{"generic-aliases", "swiss-tables", "weak-pointers"},
	},
	"1.25": {
		GoVersion:    "1.25",
		MinVersion:   "1.25.0",
		IntroducedIn: "1.25.0",
		Features:     []string{"container-gomaxprocs", "synctest", "json-v2"},
	},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run add-version-metadata.go <releases-directory>")
		fmt.Println("Example: go run add-version-metadata.go releases/v")
		os.Exit(1)
	}

	releasesDir := os.Args[1]

	fmt.Printf("Adding version metadata to lesson files in %s\n", releasesDir)

	err := filepath.Walk(releasesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Process only .go files
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		// Extract version from path (e.g., releases/v/1.18/01_generics.go -> 1.18)
		versionRegex := regexp.MustCompile(`releases/v/(\d+\.\d+)/`)
		matches := versionRegex.FindStringSubmatch(path)
		if len(matches) < 2 {
			fmt.Printf("Warning: Could not extract version from path: %s\n", path)
			return nil
		}

		version := matches[1]
		metadata, exists := versionMapping[version]
		if !exists {
			fmt.Printf("Warning: No metadata defined for version %s\n", version)
			return nil
		}

		fmt.Printf("Processing: %s (Go %s)\n", path, version)
		return addMetadataToFile(path, metadata)
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Version metadata added to all lesson files")
}

func addMetadataToFile(filePath string, metadata VersionMetadata) error {
	// Read the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %v", filePath, err)
	}

	contentStr := string(content)

	// Check if metadata already exists
	if strings.Contains(contentStr, "// GO_VERSION:") {
		fmt.Printf("  Metadata already exists in %s, skipping\n", filePath)
		return nil
	}

	// Find the position to insert metadata (after the description, before go:build)
	buildDirectiveRegex := regexp.MustCompile(`//go:build ignore`)
	buildMatch := buildDirectiveRegex.FindStringIndex(contentStr)
	if buildMatch == nil {
		return fmt.Errorf("could not find //go:build directive in %s", filePath)
	}

	// Create metadata block
	metadataBlock := fmt.Sprintf(`//
// GO_VERSION: %s
// MIN_VERSION: %s
// INTRODUCED_IN: %s
// FEATURES: %s
`,
		metadata.GoVersion,
		metadata.MinVersion,
		metadata.IntroducedIn,
		strings.Join(metadata.Features, ","),
	)

	// Insert metadata before //go:build directive
	insertPos := buildMatch[0]

	// Make sure there's a newline before the build directive
	newContent := contentStr[:insertPos] + metadataBlock + "\n" + contentStr[insertPos:]

	// Write the updated content back
	err = ioutil.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %v", filePath, err)
	}

	fmt.Printf("  ✅ Added metadata to %s\n", filePath)
	return nil
}