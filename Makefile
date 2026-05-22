.PHONY: audit help sync install-hooks dash-install dash-start dash-stop dash-restart harness-build harness-start harness-stop

# Default target
all: help

## harness-build: Compiles the Go agent harness
harness-build:
	@echo "Compiling harness..."
	@cd bin/harness && go build -o ../harness-cli .
	@echo "✅ Harness compiled to bin/harness-cli"

## harness-start: Starts the interactive Go agent harness
harness-start: harness-build
	@./bin/harness-cli

## harness-stop: Stops any running harness processes
harness-stop:
	@echo "Stopping harness processes..."
	@pkill -f bin/harness-cli || echo "No running harness found."

## audit: Runs the SDD v2.3.0 Observable Governance health check
audit:
	@python3 bin/check-health.py

## sync: Synchronizes skills to .agents and .gemini governance directories
sync:
	@chmod +x bin/sync-skills.sh
	@./bin/sync-skills.sh

## install-hooks: Installs the SDD pre-push hook into the local repository
install-hooks:
	@echo "Installing pre-push hook..."
	@cp bin/pre-push.sh .git/hooks/pre-push
	@chmod +x .git/hooks/pre-push
	@echo "✅ SDD pre-push hook installed successfully!"

## dash-install: Installs Agent Skills Hub dependencies
dash-install:
	@echo "Installing dashboard dependencies..."
	@cd hub-ui-skills && npm install

## dash-start: Updates registry and starts Agent Skills Hub
dash-start:
	@echo "Updating skill registry..."
	@cd hub-ui-skills && node scripts/update-registry.cjs
	@echo "Starting dashboard..."
	@cd hub-ui-skills && npm run dev -- --host

## dash-stop: Stops the Agent Skills Hub (kills process on port 5173)
dash-stop:
	@echo "Stopping dashboard..."
	@fuser -k 5173/tcp || echo "Dashboard is not running."

## dash-restart: Stops and restarts the Agent Skills Hub
dash-restart: dash-stop dash-start

## help: Shows this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^##' Makefile | sed -e 's/## //g' | column -t -s ':'
