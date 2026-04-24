.PHONY: help sync validate dist verify knowledge-map verify-vers clean release changelog

# Detect python command (python3 or python)
PYTHON := $(shell command -v python3 >/dev/null 2>&1 && echo python3 || echo python)

help:
	@echo "AI Agent Skills Hub - Makefile"
	@echo ""
	@echo "Available commands:"
	@echo "  make sync          - Sync mandates and skills to agents (.gemini, .claude, .agent)"
	@echo "  make validate      - Validate all skills in repository (requires uv and pyyaml)"
	@echo "  make dist          - Generate artifacts (.zip) for distribution"
	@echo "  make dist-cursor   - Compile skills for Cursor IDE .cursorrules format"
	@echo "  make verify        - Verify integrity of generated .zip artifacts"
	@echo "  make knowledge-map - Generate relational knowledge map (Mermaid)"
	@echo "  make verify-vers   - Verify if README versions match skill versions"
	@echo "  make check-memory  - Verify if state/memory is in sync with code"
	@echo "  make auto-fix      - Automatically repair missing memory files"
	@echo "  make clean         - Remove temporary folders and artifacts (dist_staging, artifacts)"
	@echo "  make release       - Full cycle: auto-fix -> sync -> validate -> verify-vers -> knowledge-map -> changelog -> dist"
	@echo "  make release-check - Run strict audit (Bunker): Ruff, Pytest, Validations, and Knowledge Map"

sync:
	uv run --with pyyaml scripts/sync_mandates.py

validate:
	uv run --with pyyaml scripts/validate_skills.py

auto-fix:
	uv run --with pyyaml scripts/auto_fix_memory.py

changelog:
	uv run --with pyyaml scripts/consolidate_changelogs.py

dist:
	bash scripts/generate-skills-dist.sh

dist-cursor:
	uv run --with pyyaml scripts/generate_cursorrules.py --skills sdd,knowledge-architect,architecture

verify:
	bash scripts/verify-dist.sh

knowledge-map:
	uv run --with pyyaml scripts/generate_knowledge_map.py

verify-vers:
	uv run --with pyyaml scripts/verify_versions.py

check-memory:
	uv run --with pyyaml scripts/verify_memory_sync.py

list-skills:
	uv run --with pyyaml scripts/installer.py list

install-skill:
	uv run --with pyyaml scripts/installer.py install $(name) --target $(target)

clean:
	rm -rf dist_staging artifacts

release-check:
	uv run ruff check .
	uv run ruff format --check .
	uv run pytest tests/
	uv run --with pyyaml scripts/validate_skills.py
	uv run --with pyyaml scripts/generate_knowledge_map.py

release: auto-fix sync validate verify-vers knowledge-map changelog dist
	@echo "🚀 Operation completed successfully!"


dashboard:
	uv run streamlit run dashboard/app.py
