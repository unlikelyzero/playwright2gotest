package gotest

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/mrichman/hargo"
)

// TestCase represents a single test case generated from HAR
type TestCase struct {
	Name     string
	Requests []Request
}

// Request represents an HTTP request in the test case
type Request struct {
	Method      string
	URL         string
	Headers     map[string]string
	Body        string
	Expectation Expectation
}

// Expectation represents the expected response
type Expectation struct {
	Status int
	Body   string
}

// Convert converts a HAR file to a Go test file
func Convert(harPath, outputPath string) error {
	// Open and read HAR file
	file, err := os.Open(harPath)
	if err != nil {
		return fmt.Errorf("failed to open HAR file: %w", err)
	}
	defer file.Close()

	// Create buffered reader
	reader := bufio.NewReader(file)

	// Parse HAR file
	har, err := hargo.Decode(reader)
	if err != nil {
		return fmt.Errorf("failed to decode HAR file: %w", err)
	}

	// Create test case
	testCase := TestCase{
		Name:     strings.TrimSuffix(filepath.Base(harPath), ".har"),
		Requests: make([]Request, 0),
	}

	// Convert entries to requests
	for _, entry := range har.Log.Entries {
		// Skip non-HTTP requests
		if !strings.HasPrefix(entry.Request.URL, "http") {
			continue
		}

		// Convert headers
		headers := make(map[string]string)
		for _, h := range entry.Request.Headers {
			headers[h.Name] = h.Value
		}

		// Create request
		req := Request{
			Method:  entry.Request.Method,
			URL:     entry.Request.URL,
			Headers: headers,
			Body:    entry.Request.PostData.Text,
			Expectation: Expectation{
				Status: entry.Response.Status,
				Body:   entry.Response.Content.Text,
			},
		}

		testCase.Requests = append(testCase.Requests, req)
	}

	// Generate Go test file
	return generateGoTest(testCase, outputPath)
}

// generateGoTest generates a Go test file from a test case
func generateGoTest(testCase TestCase, outputPath string) error {
	tmpl := `package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func Test{{.Name}}(t *testing.T) {
	client := &http.Client{}

	{{range .Requests}}
	// Test {{.Method}} {{.URL}}
	req, err := http.NewRequest("{{.Method}}", "{{.URL}}", bytes.NewBufferString(` + "`{{.Body}}`" + `))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	{{range $key, $value := .Headers}}
	req.Header.Set("{{$key}}", "{{$value}}")
	{{end}}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != {{.Expectation.Status}} {
		t.Errorf("Expected status %d, got %d", {{.Expectation.Status}}, resp.StatusCode)
	}

	{{if .Expectation.Body}}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	expectedBody := ` + "`{{.Expectation.Body}}`" + `
	if string(body) != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, string(body))
	}
	{{end}}
	{{end}}
}
`

	t, err := template.New("test").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	if err := t.Execute(f, testCase); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
} 