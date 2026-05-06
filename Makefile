.PHONY: audit help

# Default target
all: help

## audit: Runs the SDD v2.3.0 Observable Governance health check
audit:
	@python3 bin/check-health.py

## help: Shows this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^##' Makefile | sed -e 's/## //g' | column -t -s ':'
