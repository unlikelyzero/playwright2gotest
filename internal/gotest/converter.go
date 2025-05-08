package gotest

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

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

// safeHeaders is a list of headers that are safe to include in the test
var safeHeaders = map[string]bool{
	"Accept":          true,
	"Content-Type":    true,
	"Accept-Language": true,
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
	baseName := strings.TrimSuffix(filepath.Base(harPath), ".har")
	testCase := TestCase{
		Name:     strings.ToUpper(baseName[:1]) + baseName[1:],
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
			if safeHeaders[h.Name] {
				headers[h.Name] = h.Value
			}
		}

		// Create request
		req := Request{
			Method:  entry.Request.Method,
			URL:     template.HTMLEscapeString(entry.Request.URL),
			Headers: headers,
			Body:    entry.Request.PostData.Text,
			Expectation: Expectation{
				Status: entry.Response.Status,
				Body:   "", // Skip body comparison for now
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
	"io"
	"net/http"
	"testing"
)

func Test{{.Name}}(t *testing.T) {
	client := &http.Client{}
	var err error
	var req *http.Request
	var resp *http.Response

	{{range .Requests}}
	// Test {{.Method}} request
	req, err = http.NewRequest("{{.Method}}", "{{.URL}}", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	{{if .Body}}
	req.Body = io.NopCloser(bytes.NewBufferString(` + "`{{.Body}}`" + `))
	{{end}}

	{{range $key, $value := .Headers}}
	req.Header.Set("{{$key}}", "{{$value}}")
	{{end}}

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check if status code is in the expected range (2xx or 3xx)
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		t.Errorf("Expected successful status code (2xx or 3xx), got %d", resp.StatusCode)
	}
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
