package examples

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/playwright-community/playwright-go"
)

func TestHttpBinRequests(t *testing.T) {
	// Create output directory if it doesn't exist
	harDir := filepath.Join("..", "har")
	if err := os.MkdirAll(harDir, 0755); err != nil {
		t.Fatalf("Failed to create HAR directory: %v", err)
	}

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("Failed to start playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch()
	if err != nil {
		t.Fatalf("Failed to launch browser: %v", err)
	}
	defer browser.Close()

	// Create a new browser context with HAR recording enabled
	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		RecordHarPath:        playwright.String(filepath.Join(harDir, "httpbin.har")),
		RecordHarOmitContent: playwright.Bool(false),
		RecordHarMode:        playwright.HarModeMinimal,
	})
	if err != nil {
		t.Fatalf("Failed to create context: %v", err)
	}
	defer context.Close()

	// Create a new page
	page, err := context.NewPage()
	if err != nil {
		t.Fatalf("Failed to create page: %v", err)
	}

	// Make GET request
	if _, err = page.Goto("https://httpbin.org/get", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		t.Fatalf("Failed to navigate: %v", err)
	}

	// Make POST request using page.Evaluate
	script := `
		fetch('https://httpbin.org/post', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				'Accept': 'application/json'
			},
			body: JSON.stringify({message: 'Hello, World!'})
		}).then(response => response.json())
	`
	_, err = page.Evaluate(script)
	if err != nil {
		t.Fatalf("Failed to make POST request: %v", err)
	}

	// Wait for network idle to ensure all requests are completed
	if err = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
		State: playwright.LoadStateNetworkidle,
	}); err != nil {
		t.Fatalf("Failed to wait for network idle: %v", err)
	}

	// Close context to ensure HAR file is saved
	if err = context.Close(); err != nil {
		t.Fatalf("Failed to close context: %v", err)
	}
}
