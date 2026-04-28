# Tasks: HB-CLI Development

## 🛠️ Fase 3: IMPLEMENT

### 1. Project Initialization
- [x] Criar diretório `hb/` na raiz do projeto.
- [x] Inicializar módulo Go: `go mod init github.com/klebersonromero/hb`.
- [x] Instalar dependências base (`cobra`, `color`, `yaml`).
- [x] Criar estrutura de pastas: `internal/scanner`, `internal/validator`, `internal/report`, `internal/domain`.

### 2. Scanner Implementation
- [x] Implementar `Scanner` recursivo que ignora pastas da blacklist.
- [x] Implementar filtro para detectar diretórios de skills.
- [x] Adicionar suporte a concorrência (Worker Pool) para processamento paralelo.

### 3. Core Validators
- [x] Implementar validador de estrutura (check de arquivos obrigatórios).
- [x] Implementar validador de YAML (parsing do Frontmatter).
- [x] Implementar check de SemVer para versões.
- [x] Implementar validador de estado SDD (integridade de `tasks.md`).

### 4. CLI & Reporting
- [x] **Comando `hb audit`**: Validação recursiva de skills (Substituiu `validate_skills.py`).
- [x] **Comando `hb sync`**: Sincronização paralela de mandatos (Substituiu `sync_mandates.py`).
- [x] **Comando `hb fix`**: Reparo automático de estrutura (Substituiu `auto_fix_memory.py`).
- [x] **Comando `hb list`**: Listagem com suporte a JSON (Substituiu lógica no `utils.py`).
- [x] **Comando `hb install`**: Instalação híbrida (Substituiu `installer.py`).
- [x] **Comando `hb check`**: Paridade de versão e Context Drift (Substituiu `verify_versions.py` e `verify_memory_sync.py`).
- [x] **Comando `hb changelog`**: Consolidação global (Substituiu `consolidate_changelogs.py`).
- [x] **Comando `hb map`**: Geração de mapa Mermaid (Substituiu `generate_knowledge_map.py`).
- [x] **Comando `hb rules`**: Geração de `.cursorrules` (Substituiu `generate_cursorrules.py`).
- [x] **Comando `hb mode/compress`**: Gestão de densidade de tokens (Substituiu `hub.py` e `compressor_utility.py`).
- [x] **Eliminação Total de Débito Python**: 13 scripts legados removidos.
- [x] **Estratégia de CI/CD**: Implementar GitHub Action para build/release automático dos binários.

### 5. Multi-OS Distribution & CI/CD
- [x] Adicionar targets de cross-compilation no `Makefile`.
- [x] Criar workflow do GitHub Actions (`.github/workflows/ci.yml`).
- [x] Implementar comando `hb version` para rastreabilidade de build.
- [x] Criar script de instalação rápida `install.sh`.

## 🏁 Fase 4: REVIEW
- [x] Criar `test_fixtures` para validação de erros.
- [x] Rodar `hb audit` contra o próprio repositório.
- [x] Executar benchmark de tempo de execução.
- [x] Atualizar `STATE.md` e `LEARNINGS.md`.
