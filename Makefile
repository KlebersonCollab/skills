.PHONY: help sync validate dist verify knowledge-map verify-vers clean release changelog

# Detecta o comando python (python3 ou python)
PYTHON := $(shell command -v python3 >/dev/null 2>&1 && echo python3 || echo python)

help:
	@echo "AI Agent Skills Hub - Makefile"
	@echo ""
	@echo "Comandos disponíveis:"
	@echo "  make sync          - Sincroniza mandatos e skills para os agentes (.gemini, .claude, .agent)"
	@echo "  make validate      - Valida todas as skills no repositório (exige uv e pyyaml)"
	@echo "  make dist          - Gera os artefatos (.zip) para distribuição"
	@echo "  make dist-cursor   - Compila as skills para o formato .cursorrules da IDE Cursor"
	@echo "  make verify        - Verifica a integridade dos artefatos .zip gerados"
	@echo "  make knowledge-map - Gera o mapa de conhecimento relacional (Mermaid)"
	@echo "  make verify-vers   - Verifica se as versões no README batem com as das skills"
	@echo "  make check-memory  - Verifica se o estado/memória está em sync com o código"
	@echo "  make auto-fix      - Repara automaticamente arquivos de memória ausentes"
	@echo "  make clean         - Remove pastas temporárias e artefatos (dist_staging, artifacts)"
	@echo "  make release       - Ciclo completo: auto-fix -> sync -> validate -> verify-vers -> knowledge-map -> changelog -> dist"

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

release: auto-fix sync validate verify-vers knowledge-map changelog dist
	@echo "🚀 Operação concluída com sucesso!"

