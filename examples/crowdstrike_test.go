package examples

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
)

func TestCrowdStrikeInteractions(t *testing.T) {
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

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false), // Set to true in CI
		Args: []string{
			"--disable-web-security",
			"--no-sandbox",
		},
	})
	if err != nil {
		t.Fatalf("Failed to launch browser: %v", err)
	}
	defer browser.Close()

	// Create a new browser context with HAR recording enabled
	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		RecordHarPath:        playwright.String(filepath.Join(harDir, "crowdstrike.har")),
		RecordHarOmitContent: playwright.Bool(false),
		RecordHarMode:        playwright.HarModeMinimal,
		Viewport: &playwright.Size{
			Width:  1280,
			Height: 720,
		},
		IgnoreHttpsErrors: playwright.Bool(true),
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

	// Set longer timeout for navigation
	page.SetDefaultTimeout(60000) // 60 seconds

	// Navigate to CrowdStrike homepage
	t.Log("Navigating to CrowdStrike homepage...")
	if _, err = page.Goto("https://www.crowdstrike.com/en-us/", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
		Timeout:   playwright.Float(60000), // 60 seconds
	}); err != nil {
		t.Fatalf("Failed to navigate to homepage: %v", err)
	}
	t.Log("Successfully navigated to homepage")

	// Handle any popup dialogs
	t.Log("Handling popup dialogs...")
	if err = page.Click("button[role='button'][name='Close']", playwright.PageClickOptions{
		Timeout: playwright.Float(5000),
	}); err != nil {
		t.Logf("First close button not found: %v", err)
	}

	if err = page.Click("button[role='button'][name='Close'][exact='true']", playwright.PageClickOptions{
		Timeout: playwright.Float(5000),
	}); err != nil {
		t.Logf("Second close button not found: %v", err)
	}

	// Click on the search icon
	t.Log("Clicking search icon...")
	if err = page.Click("button[role='button'][name='Search Icon']", playwright.PageClickOptions{
		Timeout: playwright.Float(5000),
	}); err != nil {
		t.Fatalf("Failed to click search icon: %v", err)
	}

	// Wait for search input to be visible and type search query
	t.Log("Filling search input...")
	if err = page.Fill("input[role='searchbox'][name='Search field']", "blash", playwright.PageFillOptions{
		Timeout: playwright.Float(5000),
	}); err != nil {
		t.Fatalf("Failed to fill search input: %v", err)
	}

	// Click on the search results text
	t.Log("Clicking search results...")
	if err = page.Click("text=Results for blash", playwright.PageClickOptions{
		Timeout: playwright.Float(5000),
	}); err != nil {
		t.Fatalf("Failed to click search results: %v", err)
	}

	// Add a small delay to ensure all requests are captured
	time.Sleep(2 * time.Second)

	// Close context to ensure HAR file is saved
	t.Log("Saving HAR file...")
	if err = context.Close(); err != nil {
		t.Fatalf("Failed to close context: %v", err)
	}
	t.Log("Test completed successfully")
}
