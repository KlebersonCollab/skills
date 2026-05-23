.PHONY: audit skill-validate help sync install-hooks dash-install dash-start dash-stop dash-restart
.PHONY: harness-build harness-start harness-stop harness-clean harness-test harness-vet
.PHONY: harness-print harness-continue harness-resume harness-rpc harness-fork harness-models harness-with
.PHONY: harness-gemini harness-deepseek harness-vl2 harness-thinking harness-info
.PHONY: harness-voice harness-voice-listen harness-voice-speak harness-voice-full

# Default target
all: help

## harness-build: Compila o harness e todos os pacotes internos
harness-build:
	@echo "🔨 Compilando harness (todos os pacotes)..."
	@cd bin/harness && go build -o ../harness-cli . && go build ./...
	@echo "✅ Harness compilado: bin/harness-cli"
	@echo "✅ Todos os pacotes internos: OK"

## harness-test: Roda todos os testes
harness-test:
	@echo "🧪 Testando harness..."
	@cd bin/harness && go test ./... -count=1 -timeout 30s
	@echo "✅ Todos os testes passaram"

## harness-vet: Roda go vet
harness-vet:
	@echo "🔍 Verificando harness (go vet)..."
	@cd bin/harness && go vet ./...
	@echo "✅ go vet: OK"

## harness-clean: Limpa artefatos
harness-clean:
	@rm -f bin/harness-cli bin/harness/harness bin/harness/harness_new
	@echo "✅ Artefatos removidos"

## harness-start: Inicia o modo interativo
harness-start: harness-build
	@./bin/harness-cli

## harness-stop: Para processos do harness
harness-stop:
	@echo "Parando harness..."
	@pkill -f bin/harness-cli 2>/dev/null || echo "Nenhum em execucao."

## harness-print: Modo print (non-interactive) — uso: make harness-print MSG="prompt"
harness-print: harness-build
	@./bin/harness-cli -p "$(MSG)"

## harness-continue: Continua a ultima sessao
harness-continue: harness-build
	@./bin/harness-cli -c

## harness-resume: Navega e retoma sessao
harness-resume: harness-build
	@./bin/harness-cli -r

## harness-rpc: Inicia modo RPC (stdin/stdout JSONL)
harness-rpc: harness-build
	@./bin/harness-cli --mode rpc

## harness-fork: Cria fork de uma sessao — uso: make harness-fork ID=<session-id>
harness-fork: harness-build
	@./bin/harness-cli --fork "$(ID)"

## harness-models: Lista providers e modelos disponiveis
harness-models:
	@cd bin/harness && go run . --list-models 2>/dev/null || ./bin/harness-cli --list-models

## harness-with: Inicia com provider/modelo especifico — uso: make harness-with PROVIDER=deepseek MODEL=deepseek-vl2
harness-with: harness-build
	@HARNESS_PROVIDER="$(PROVIDER)" HARNESS_MODEL="$(MODEL)" ./bin/harness-cli

## harness-vl2: Inicia modo interativo com DeepSeek VL2 (visao)
harness-vl2: harness-build
	@HARNESS_PROVIDER="deepseek" HARNESS_MODEL="deepseek-vl2" ./bin/harness-cli

## harness-gemini: Inicia modo interativo com Gemini (se configurado)
harness-gemini: harness-build
	@HARNESS_PROVIDER="gemini" ./bin/harness-cli

## harness-thinking: Modo thinking alto — uso: make harness-thinking MSG="prompt"
harness-thinking: harness-build
	@./bin/harness-cli --thinking high -p "$(MSG)"

## harness-info: Mostra info do harness (versao + models)
harness-info:
	@echo "=== Harness Info ==="
	@cd bin/harness && go run . --version 2>/dev/null || ./bin/harness-cli --version
	@echo ""
	@cd bin/harness && go run . --list-models 2>/dev/null || ./bin/harness-cli --list-models

# ── Voice Mode ──────────────────────────────────────────────────────────

## harness-voice: Inicia modo de voz completo (STT + TTS) — uso: make harness-voice
harness-voice: harness-build
	@./bin/harness-cli --voice

## harness-voice-listen: Modo voz apenas escuta (STT, sem TTS) — uso: make harness-voice-listen
harness-voice-listen: harness-build
	@./bin/harness-cli --voice=listen

## harness-voice-speak: Modo voz apenas fala (TTS, texto como input) — uso: make harness-voice-speak
harness-voice-speak: harness-build
	@./bin/harness-cli --voice=speak

## harness-voice-full: Modo voz completo (STT + TTS, mesmo que harness-voice) — uso: make harness-voice-full
harness-voice-full: harness-build
	@./bin/harness-cli --voice=full

## harness-voice-duration: Modo voz com duracao personalizada — uso: make harness-voice-duration SEC=5
harness-voice-duration: harness-build
	@./bin/harness-cli --voice --voice-duration=$(SEC)

# ── Governance & Quality ────────────────────────────────────────────────

## audit: Runs the SDD v2.3.0 Observable Governance health check (includes skill validation)
audit: skill-validate
	@python3 bin/check-health.py

## skill-validate: Validates all .skill.md files for metadata integrity and internal links
skill-validate:
	@python3 bin/validate-skills.py

## sync: Synchronizes skills to .agents and .gemini governance directories
sync:
	@chmod +x bin/sync-skills.sh
	@./bin/sync-skills.sh

## install-hooks: Installs the SDD pre-push hook into the local repository
install-hooks:
	@echo "Installing pre-push hook..."
	@cp bin/pre-push.sh .git/hooks/pre-hook
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