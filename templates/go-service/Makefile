# Project Governance Makefile

HB_BIN := hb

.PHONY: setup audit status sync help

help:
	@echo "🛠️ Project Governance Commands"
	@echo ""
	@echo "  make setup   - Initialize governance structure"
	@echo "  make audit   - Run integrity and compliance audit"
	@echo "  make status  - Show development progress (SDD)"
	@echo "  make sync    - Sync skills from the Hub"

setup:
	@echo "📥 Installing HB-CLI..."
	@curl -sSL https://raw.githubusercontent.com/KlebersonCollab/skills/main/scripts/install-hb.sh | bash
	@hb sdd bootstrap
	@echo "✅ Governance ready!"

audit:
	hb audit

status:
	hb sdd status

sync:
	hb sync
