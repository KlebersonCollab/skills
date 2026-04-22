# Skill: Git Workflow

## Overview
Esta skill define os padrões de controle de versão e fluxos de trabalho colaborativos para o projeto, garantindo um histórico limpo, rastreável e integrado ao ciclo **SDD (Spec-Driven Development)**.

## O Que Você Pode Fazer
- Escolher a melhor **estratégia de branching** para o seu contexto.
- Escrever **mensagens de commit** profissionais usando Conventional Commits.
- Realizar **rebases** e resolver conflitos com segurança.
- Abrir **Pull Requests** de alta qualidade com templates padronizados.
- Garantir que o Git reflita o progresso das tarefas no SDD.

## Guia Rápido de Comandos

### Iniciando uma Feature
```bash
git checkout main
git pull origin main
git checkout -b feature/minha-feature
```

### Committing (Conventional Commits + Inglês Mandatório)
> **Nota**: Embora a documentação esteja em Português, as mensagens de commit devem ser sempre em **Inglês**.

```bash
git add .
git commit -m "feat(auth): implement password recovery flow"
```

### Sincronizando com a Main (Rebase)
```bash
git fetch origin
git rebase origin/main
# Se houver conflitos, resolva e:
# git add .
# git rebase --continue
```

## Anti-patterns (O que NÃO fazer)
- ❌ Commits gigantes com muitas alterações não relacionadas.
- ❌ Mensagens de commit como "fix", "update", "wip".
- ❌ Deixar branches de feature vivas por semanas.
- ❌ Forçar push (`push -f`) em branches compartilhadas sem aviso.

## Referências
- [Conventional Commits](references/conventional-commits.md)
- [Branching Strategies](references/branching-strategies.md)
- [Template de Commit](examples/gitmessage.template)
- [Template de PR](examples/pull-request.md)
