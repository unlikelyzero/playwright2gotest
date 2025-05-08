package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"playwright2gotest/internal/gotest"
)

func main() {
	// Define command line flags
	inputDir := flag.String("input", "", "Directory containing HAR files")
	outputDir := flag.String("output", "tests", "Directory to output generated Go tests")
	flag.Parse()

	// Validate input directory
	if *inputDir == "" {
		log.Fatal("Input directory is required")
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Find all HAR files in the input directory
	harFiles, err := filepath.Glob(filepath.Join(*inputDir, "*.har"))
	if err != nil {
		log.Fatalf("Failed to find HAR files: %v", err)
	}

	// Convert each HAR file to Go test
	for _, harFile := range harFiles {
		outputFile := filepath.Join(*outputDir, filepath.Base(harFile[:len(harFile)-4])+".go")
		
		if err := gotest.Convert(harFile, outputFile); err != nil {
			log.Printf("Failed to convert %s: %v", harFile, err)
			continue
		}
		
		fmt.Printf("Converted %s to %s\n", harFile, outputFile)
	}
} 