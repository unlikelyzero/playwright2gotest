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
	inputPath := flag.String("input", "", "HAR file or directory containing HAR files")
	outputPath := flag.String("output", "tests", "Output file or directory for generated Go tests")
	flag.Parse()

	// Validate input path
	if *inputPath == "" {
		log.Fatal("Input path is required")
	}

	// Check if input is a directory
	inputInfo, err := os.Stat(*inputPath)
	if err != nil {
		log.Fatalf("Failed to stat input path: %v", err)
	}

	if inputInfo.IsDir() {
		// Create output directory if it doesn't exist
		if err := os.MkdirAll(*outputPath, 0755); err != nil {
			log.Fatalf("Failed to create output directory: %v", err)
		}

		// Find all HAR files in the input directory
		harFiles, err := filepath.Glob(filepath.Join(*inputPath, "*.har"))
		if err != nil {
			log.Fatalf("Failed to find HAR files: %v", err)
		}

		// Convert each HAR file to Go test
		for _, harFile := range harFiles {
			outputFile := filepath.Join(*outputPath, filepath.Base(harFile[:len(harFile)-4])+"_test.go")

			if err := gotest.Convert(harFile, outputFile); err != nil {
				log.Printf("Failed to convert %s: %v", harFile, err)
				continue
			}

			fmt.Printf("Converted %s to %s\n", harFile, outputFile)
		}
	} else {
		// Single file conversion
		if err := gotest.Convert(*inputPath, *outputPath); err != nil {
			log.Fatalf("Failed to convert %s: %v", *inputPath, err)
		}
		fmt.Printf("Converted %s to %s\n", *inputPath, *outputPath)
	}
}
