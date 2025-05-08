.PHONY: test test-debug install clean test-crowdstrike test-httpbin test-all install-ts test-ts test-ts-debug

# Default target
all: install install-ts test-all

# Install Go dependencies
install:
	go mod tidy
	go run github.com/playwright-community/playwright-go/cmd/playwright install

# Install TypeScript dependencies
install-ts:
	npm install
	npx playwright install

# Run all Go tests
test-all:
	cd examples && go test -v ./...

# Run specific Go test suites
test-crowdstrike:
	cd examples && go test -v -run TestCrowdStrikeInteractions

test-httpbin:
	cd examples && go test -v -run TestHttpBinRequests

# Run all TypeScript tests
test-ts:
	npm test

# Run specific TypeScript test suites
test-ts-crowdstrike:
	npm run test:crowdstrike

test-ts-httpbin:
	npm run test:httpbin

# Run Go tests with Playwright debug mode
test-debug:
	cd examples && PWDEBUG=1 go test -v ./...

# Run specific Go test with debug mode
test-debug-crowdstrike:
	cd examples && PWDEBUG=1 go test -v -run TestCrowdStrikeInteractions

test-debug-httpbin:
	cd examples && PWDEBUG=1 go test -v -run TestHttpBinRequests

# Run TypeScript tests with debug mode
test-ts-debug:
	npm run test:debug

# Run Go tests with Playwright debug mode and show browser
test-debug-browser:
	cd examples && PWDEBUG=1 PWDEBUG_HEADLESS=0 go test -v ./...

# Run specific Go test with debug mode and browser
test-debug-browser-crowdstrike:
	cd examples && PWDEBUG=1 PWDEBUG_HEADLESS=0 go test -v -run TestCrowdStrikeInteractions

test-debug-browser-httpbin:
	cd examples && PWDEBUG=1 PWDEBUG_HEADLESS=0 go test -v -run TestHttpBinRequests

# Clean generated files
clean:
	rm -rf har/*.har
	rm -rf tests/*_test.go
	rm -rf node_modules
	rm -rf dist

# Help command
help:
	@echo "Available commands:"
	@echo "  make install                    - Install Go dependencies and Playwright browsers"
	@echo "  make install-ts                 - Install TypeScript dependencies and Playwright browsers"
	@echo "  make test-all                   - Run all Go tests"
	@echo "  make test-ts                    - Run all TypeScript tests"
	@echo "  make test-crowdstrike           - Run only Go CrowdStrike tests"
	@echo "  make test-ts-crowdstrike        - Run only TypeScript CrowdStrike tests"
	@echo "  make test-httpbin              - Run only Go HttpBin tests"
	@echo "  make test-ts-httpbin           - Run only TypeScript HttpBin tests"
	@echo "  make test-debug                 - Run all Go tests with Playwright debug mode"
	@echo "  make test-ts-debug             - Run all TypeScript tests with debug mode"
	@echo "  make test-debug-crowdstrike     - Run Go CrowdStrike tests with debug mode"
	@echo "  make test-debug-httpbin        - Run Go HttpBin tests with debug mode"
	@echo "  make test-debug-browser         - Run all Go tests with debug mode and browser"
	@echo "  make test-debug-browser-crowdstrike - Run Go CrowdStrike tests with debug mode and browser"
	@echo "  make test-debug-browser-httpbin - Run Go HttpBin tests with debug mode and browser"
	@echo "  make clean                      - Remove generated files"
	@echo "  make help                       - Show this help message" 