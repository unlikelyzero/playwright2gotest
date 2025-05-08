package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestHttpbin(t *testing.T) {
	client := &http.Client{}

	// Test GET https://httpbin.org/get
	req1, err := http.NewRequest("GET", "https://httpbin.org/get", bytes.NewBufferString(``))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req1.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req1.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req1.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/136.0.7103.25 Safari/537.36")

	resp1, err := client.Do(req1)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp1.Body.Close()

	if resp1.StatusCode != 200 {
		t.Errorf("Expected status %d, got %d", 200, resp1.StatusCode)
	}

	body1, err := io.ReadAll(resp1.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// Parse response JSON
	var respData1 map[string]interface{}
	if err := json.Unmarshal(body1, &respData1); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	// Check important fields
	headers1 := respData1["headers"].(map[string]interface{})
	if headers1["Accept"] != "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7" {
		t.Errorf("Unexpected Accept header: %v", headers1["Accept"])
	}
	if respData1["url"] != "https://httpbin.org/get" {
		t.Errorf("Unexpected URL: %v", respData1["url"])
	}

	// Test POST https://httpbin.org/post
	req2, err := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBufferString(`{"message":"Hello, World!"}`))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req2.Header.Set("Accept", "application/json")
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/136.0.7103.25 Safari/537.36")

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
		t.Errorf("Unexpected Accept header: %v", headers2["Accept"])
	}
	if headers2["Content-Type"] != "application/json" {
		t.Errorf("Unexpected Content-Type header: %v", headers2["Content-Type"])
	}
	if respData2["url"] != "https://httpbin.org/post" {
		t.Errorf("Unexpected URL: %v", respData2["url"])
	}

	// Check the message in the JSON body
	jsonData := respData2["json"].(map[string]interface{})
	if jsonData["message"] != "Hello, World!" {
		t.Errorf("Unexpected message: %v", jsonData["message"])
	}
}
