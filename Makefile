.PHONY: audit help sync install-hooks

# Default target
all: help

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
	@cp git-workflow/resources/pre-push.sh .git/hooks/pre-push
	@chmod +x .git/hooks/pre-push
	@echo "✅ SDD pre-push hook installed successfully!"

## help: Shows this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^##' Makefile | sed -e 's/## //g' | column -t -s ':'
