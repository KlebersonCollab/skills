# Spec: Python com UV Skill

## 1. Overview
Criar uma skill para uso e boas práticas de Python utilizando UV como gerenciador de pacotes e versionamento de Python.

## 2. Functional Requirements

### FR-1: Instalação e Configuração do UV
- A skill deve guiar na instalação do UV
- Configuração de ambientes Python com UV
- Gerenciamento de versões de Python

### FR-2: Gerenciamento de Dependências
- Criação e gerenciamento de `pyproject.toml`
- Instalação de pacotes com UV
- Gerenciamento de ambientes virtuais
- Lockfile (`uv.lock`) e reprodução de ambientes

### FR-3: Boas Práticas de Python Moderno
- Estrutura de projetos Python modernos
- Type hints e type checking
- Formatação de código (ruff)
- Linting e análise estática
- Testes automatizados

### FR-4: Workflows de Desenvolvimento
- Desenvolvimento local com UV
- CI/CD com UV
- Packaging e publicação
- Monorepo management

### FR-5: Integração com Ferramentas Modernas
- Integração com IDEs (VS Code, PyCharm)
- Ferramentas de qualidade de código
- Ferramentas de documentação

## 3. Acceptance Criteria

### AC-1: Documentação Completa
- SKILL.md com estrutura padrão
- README.md com exemplos práticos
- CHANGELOG.md vazio (1.0.0)

### AC-2: Padrão Skill Factory
- Seguir templates do Skill Factory
- Frontmatter correto (name, version, description, category)
- Diretório `python-uv/` na raiz

### AC-3: Conteúdo Técnico Corrigo
- Referências à documentação oficial do UV
- Exemplos práticos e testáveis
- Boas práticas validadas
- Padrões de projeto Python moderno

### AC-4: Registro no Hub
- Adicionar na tabela do README.md raiz
- Atualizar badge de contagem de skills
- Link correto para documentação

## 4. Non-Functional Requirements

### NFR-1: Qualidade
- Conformidade com padrões do Skills Hub
- Validação pelo Skill Factory Validator
- Documentação clara e concisa

### NFR-2: Usabilidade
- Exemplos práticos e reproduzíveis
- Passo a passo claro
- Referências à documentação oficial

### NFR-3: Manutenibilidade
- Versionamento semântico (iniciar em 1.0.0)
- Estrutura modular (preparada para sub-skills futuras)
- Documentação atualizável

## 5. Out of Scope
- Tutoriais avançados de Python
- Configuração específica de CI/CD providers
- Integração com outras linguagens
- Gerenciamento de infraestrutura

## 6. Dependencies
- Skill Factory (para scaffolding)
- SDD (para metodologia de desenvolvimento)
- Documentação oficial do UV (https://docs.astral.sh/uv/)

## 7. Risks
- **R1**: UV ainda é relativamente novo - mitigar referenciando documentação oficial
- **R2**: Múltiplas formas de usar Python - focar em padrões modernos e UV
- **R3**: Configurações específicas de SO - documentar instalação cross-platform