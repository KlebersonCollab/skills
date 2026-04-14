# Technical Plan: Python com UV Skill

## 1. Architecture

### 1.1 Skill Structure
```
python-uv/
├── README.md          # Documentação completa
├── SKILL.md           # Definição técnica principal
├── CHANGELOG.md       # Histórico de versões
└── (futuras sub-skills)
```

### 1.2 Frontmatter Template
```yaml
name: python-uv
version: 1.0.0
description: "Skill para uso e boas práticas de Python utilizando UV como gerenciador de pacotes e versionamento"
category: development-workflow
```

## 2. Content Outline

### 2.1 SKILL.md Sections
1. **Goal**: Capacitar desenvolvedores a usar Python moderno com UV
2. **Output Structure**: Guias práticos, comandos, configurações
3. **Quality Rules**: Referência à documentação oficial, exemplos testáveis
4. **Prohibited**: Não assumir configurações específicas de SO

### 2.2 README.md Sections
1. **Visão Geral**: O que é UV e por que usar
2. **Instalação**: Passo a passo cross-platform
3. **Configuração Inicial**: Primeiros passos com UV
4. **Gerenciamento de Dependências**: pyproject.toml, lockfiles
5. **Boas Práticas**: Type hints, formatação, testes
6. **Workflows**: Desenvolvimento, CI/CD, packaging
7. **Referências**: Links para documentação oficial

## 3. Implementation Tasks

### T1: Scaffolding da Skill
- Usar Skill Factory para criar estrutura básica
- Parâmetros: skill_name="python-uv", description="Skill para uso e boas práticas de Python utilizando UV como gerenciador de pacotes e versionamento", category="development-workflow"

### T2: Conteúdo Técnico Principal
- Pesquisar documentação oficial do UV
- Compilar comandos essenciais
- Criar exemplos práticos
- Documentar boas práticas

### T3: Validação
- Executar Skill Factory Validator
- Corrigir não conformidades
- Garantir padrão do hub

### T4: Registro
- Atualizar README.md raiz
- Adicionar à tabela de skills
- Incrementar badge de contagem

## 4. Technical Considerations

### 4.1 Cross-Platform Support
- Documentar instalação para macOS, Linux, Windows
- Comandos compatíveis com todos os SOs
- Soluções para problemas comuns

### 4.2 Version Management
- UV gerencia versões de Python
- pyproject.toml como fonte única da verdade
- Lockfiles para reprodução exata

### 4.3 Modern Python Practices
- Type hints obrigatórios
- Ruff para formatação e linting
- pytest para testes
- mypy para type checking

## 5. Success Metrics

### 5.1 Completeness
- [ ] Scaffolding criado com sucesso
- [ ] Conteúdo técnico abrangente
- [ ] Validação aprovada
- [ ] Registro atualizado

### 5.2 Quality
- [ ] Referências à documentação oficial
- [ ] Exemplos práticos e testáveis
- [ ] Boas práticas documentadas
- [ ] Estrutura seguindo padrões do hub

## 6. Timeline
- **Fase 1 (Scaffolding)**: 15 minutos
- **Fase 2 (Conteúdo)**: 30 minutos
- **Fase 3 (Validação)**: 10 minutos
- **Fase 4 (Registro)**: 5 minutos

**Total estimado**: 60 minutos

## 7. Dependencies
- Skill Factory (para scaffolding)
- SDD (para metodologia)
- Documentação oficial do UV
- Conhecimento de Python moderno