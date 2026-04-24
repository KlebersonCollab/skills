.PHONY: help build release audit sync rules check map list status clean dashboard

HB_BIN := ./bin/hb

help:
	@echo "🚀 AI Agent Skills Hub - Governance CLI"
	@echo ""
	@echo "Available commands:"
	@echo "  make build         - Build the local hb binary"
	@echo "  make release       - Cross-compile hb for all supported OS"
	@echo "  make audit         - Run full project integrity audit"
	@echo "  make sync          - Sync skills to local agent folders"
	@echo "  make rules         - Distribute agent rules (.gemini, .claude, .agent)"
	@echo "  make check         - Verify links and version consistency"
	@echo "  make map           - Generate knowledge map (Mermaid)"
	@echo "  make list          - List all available skills"
	@echo "  make status        - Show SDD development status"
	@echo "  make clean         - Remove artifacts and temporary folders"
	@echo "  make dashboard     - Run the visual dashboard (Python/Streamlit)"

build:
	@mkdir -p bin
	cd hb && go build -o ../bin/hb .
	@echo "✅ Binary built at $(HB_BIN)"

release:
	@mkdir -p bin
	cd hb && make build-all
	@mv hb/bin/* bin/
	@echo "✅ Multi-OS releases ready in bin/"

audit: build
	$(HB_BIN) audit

sync: build
	$(HB_BIN) sync

rules: build
	$(HB_BIN) rules

check: build
	$(HB_BIN) check

map: build
	$(HB_BIN) map

list: build
	$(HB_BIN) list

status: build
	$(HB_BIN) sdd status

clean:
	rm -rf bin/
	rm -rf artifacts/
	find . -name "__pycache__" -type d -exec rm -rf {} +

dashboard:
	@echo "📊 Starting Dashboard..."
	streamlit run dashboard/app.py
