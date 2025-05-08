package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestSample(t *testing.T) {
	client := &http.Client{}

	// Test GET https://httpbin.org/get
	req, err := http.NewRequest("GET", "https://httpbin.org/get", bytes.NewBufferString(``))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Sample HAR Test")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Parse response JSON
	var respData map[string]interface{}
	if err := json.Unmarshal(body, &respData); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	// Check important fields
	headers := respData["headers"].(map[string]interface{})
	if headers["Accept"] != "application/json" {
		t.Errorf("Expected Accept header to be application/json, got %v", headers["Accept"])
	}
	if headers["User-Agent"] != "Sample HAR Test" {
		t.Errorf("Expected User-Agent header to be Sample HAR Test, got %v", headers["User-Agent"])
	}
	if respData["url"] != "https://httpbin.org/get" {
		t.Errorf("Expected URL to be https://httpbin.org/get, got %v", respData["url"])
	}

	// Test POST https://httpbin.org/post
	req2, err := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBufferString(`{"message":"Hello, World!"}`))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req2.Header.Set("Accept", "application/json")
	req2.Header.Set("Content-Type", "application/json")

	resp2, err := client.Do(req2)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp2.Body.Close()

	if resp2.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp2.StatusCode)
	}

	body2, err := io.ReadAll(resp2.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Parse response JSON
	var respData2 map[string]interface{}
	if err := json.Unmarshal(body2, &respData2); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	// Check important fields
	headers2 := respData2["headers"].(map[string]interface{})
	if headers2["Accept"] != "application/json" {
		t.Errorf("Expected Accept header to be application/json, got %v", headers2["Accept"])
	}
	if headers2["Content-Type"] != "application/json" {
		t.Errorf("Expected Content-Type header to be application/json, got %v", headers2["Content-Type"])
	}
	if respData2["url"] != "https://httpbin.org/post" {
		t.Errorf("Expected URL to be https://httpbin.org/post, got %v", respData2["url"])
	}

	// Check the message in the JSON body
	jsonData := respData2["json"].(map[string]interface{})
	if jsonData["message"] != "Hello, World!" {
		t.Errorf("Expected message to be Hello, World!, got %v", jsonData["message"])
	}
}
