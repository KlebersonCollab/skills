# Specification: Golang Expert Skill

## 🎯 Goal
Criar uma skill de nível "Expert" para desenvolvimento em Go (Golang), compilando o conhecimento de 35 skills específicas do ecossistema Samber e da comunidade Go em uma única unidade de conhecimento robusta, idiomática e alinhada aos padrões do Skills Hub.

## 📋 Acceptance Criteria (BDD)

### Cenário 1: Estrutura da Skill
- **Given** que estou criando a skill `golang-expert`.
- **When** eu utilizar o `skill-factory-bootstrap`.
- **Then** a estrutura de arquivos deve seguir rigorosamente o padrão do Hub (`SKILL.md`, `README.md`, `CHANGELOG.md`, `references/`, `examples/`).

### Cenário 2: Consolidação de Conhecimento
- **Given** a lista de 35 skills de Go do repositório `samber/cc-skills-golang`.
- **When** eu compilar o conteúdo na nova skill.
- **Then** todos os 5 pilares (Fundamentos, Concorrência, Ferramentas, Arquitetura e Ecossistema) devem estar documentados no `SKILL.md`.

### Cenário 3: Validação de Conformidade
- **Given** que a skill foi criada.
- **When** eu executar o `skill-factory-validator`.
- **Then** o relatório deve emitir o status `COMPLIANT` com score 100/100.

## 🛠️ Constraints
- Utilizar obrigatoriamente o framework **SDD**.
- Seguir os mandatos globais de **Clean Code** (SOLID, DRY, KISS).
- A skill deve ser escrita em **Português Brasileiro**.
- Versão inicial deve ser `1.0.0`.

## 🏗️ Technical Architecture
- A skill será um arquivo `SKILL.md` autossuficiente com referências externas para guias detalhados na pasta `references/`.
- Exemplos de código devem usar o pacote nativo `testing` e, opcionalmente, `testify`.
- Foco total em Go moderno (v1.20+), utilizando Generics onde apropriado.
