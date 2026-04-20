# Specification: Skill Installer CLI

**Feature ID:** FR-SKILL-INSTALLER
**Status:** DRAFT
**Complexity:** Medium
**Category:** Automation / Distribution

## Goal
Prover uma interface de linha de comando (CLI) que permita a instalação granular de skills individuais em projetos externos ou diretórios específicos, facilitando a portabilidade do Hub sem a necessidade de copiar manualmente arquivos.

## 🎯 Requisitos Funcionais

| ID | Título | Descrição |
|----|--------|-----------|
| FR-1 | **Listagem de Skills** | O comando deve listar todas as skills disponíveis no hub local com suas versões. |
| FR-2 | **Instalação Direta** | O comando `install <skill_name>` deve copiar a skill especificada para o diretório de destino. |
| FR-3 | **Destino Customizável** | Permitir definir o diretório de destino via flag `--target` (default: `./.agent/skills/`). |
| FR-4 | **Resolução de Dependências** | Se uma skill depende de outra (declarado no metadata), o instalador deve perguntar se deseja instalar as dependências. |
| FR-5 | **Modo de Sobrescrita** | Flag `--force` para permitir sobrescrever skills existentes. |

## ✅ Critérios de Aceitação (BDD)

### Cenário 1: Instalação bem-sucedida de uma skill
**Given** que eu tenho o Hub de Skills acessível
**When** eu executo `uv run scripts/installer.py install python-uv --target ./test_project`
**Then** a pasta `python-uv` deve ser criada em `./test_project/python-uv`
**And** todos os arquivos (SKILL.md, README.md, etc.) devem estar presentes.

### Cenário 2: Erro ao instalar skill inexistente
**Given** que eu tento instalar uma skill que não existe
**When** eu executo `uv run scripts/installer.py install non-existent-skill`
**Then** o sistema deve retornar um erro claro: "Skill 'non-existent-skill' não encontrada no Hub."

## 🛠️ Restrições Técnicas
- Linguagem: Python 3.12+
- Dependências: `pathlib`, `shutil`, `argparse`.
- Deve usar `scripts/utils.py` para localizar as skills.
- Operação atômica (usar `safe_copytree`).

## 📊 Matriz de Rastreabilidade
- **FR-2** -> `safe_copytree` implementation.
- **FR-4** -> Metadata analysis (`get_skill_metadata`).
