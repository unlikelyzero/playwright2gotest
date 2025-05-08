.PHONY: test test-debug install clean test-crowdstrike test-httpbin test-all

# Default target
all: install test-all

# Install dependencies
install:
	go mod tidy
	go run github.com/playwright-community/playwright-go/cmd/playwright install

# Run all tests
test-all:
	cd examples && go test -v ./...

# Run specific test suites
test-crowdstrike:
	cd examples && go test -v -run TestCrowdStrikeInteractions

test-httpbin:
	cd examples && go test -v -run TestHttpBinRequests

# Run tests with Playwright debug mode
test-debug:
	cd examples && PWDEBUG=1 go test -v ./...

# Run specific test with debug mode
test-debug-crowdstrike:
	cd examples && PWDEBUG=1 go test -v -run TestCrowdStrikeInteractions

test-debug-httpbin:
	cd examples && PWDEBUG=1 go test -v -run TestHttpBinRequests

# Run tests with Playwright debug mode and show browser
test-debug-browser:
	cd examples && PWDEBUG=1 PWDEBUG_HEADLESS=0 go test -v ./...

# Run specific test with debug mode and browser
test-debug-browser-crowdstrike:
	cd examples && PWDEBUG=1 PWDEBUG_HEADLESS=0 go test -v -run TestCrowdStrikeInteractions

test-debug-browser-httpbin:
	cd examples && PWDEBUG=1 PWDEBUG_HEADLESS=0 go test -v -run TestHttpBinRequests

# Clean generated files
clean:
	rm -rf har/*.har
	rm -rf tests/*_test.go

# Help command
help:
	@echo "Available commands:"
	@echo "  make install                    - Install dependencies and Playwright browsers"
	@echo "  make test-all                   - Run all tests"
	@echo "  make test-crowdstrike           - Run only CrowdStrike tests"
	@echo "  make test-httpbin              - Run only HttpBin tests"
	@echo "  make test-debug                 - Run all tests with Playwright debug mode"
	@echo "  make test-debug-crowdstrike     - Run CrowdStrike tests with debug mode"
	@echo "  make test-debug-httpbin        - Run HttpBin tests with debug mode"
	@echo "  make test-debug-browser         - Run all tests with debug mode and browser"
	@echo "  make test-debug-browser-crowdstrike - Run CrowdStrike tests with debug mode and browser"
	@echo "  make test-debug-browser-httpbin - Run HttpBin tests with debug mode and browser"
	@echo "  make clean                      - Remove generated files"
	@echo "  make help                       - Show this help message" 