.PHONY: audit help sync

# Default target
all: help

## audit: Runs the SDD v2.3.0 Observable Governance health check
audit:
	@python3 bin/check-health.py

## sync: Synchronizes skills to .agents and .gemini governance directories
sync:
	@chmod +x bin/sync-skills.sh
	@./bin/sync-skills.sh

## help: Shows this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^##' Makefile | sed -e 's/## //g' | column -t -s ':'
