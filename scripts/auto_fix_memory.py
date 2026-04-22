from pathlib import Path


def auto_fix_memory(target_path="."):
    """
    Garante que a estrutura de memória e estado (.specs/project/) exista
    com templates básicos caso arquivos estejam ausentes.
    """
    root = Path(target_path)
    specs_project = root / ".specs" / "project"

    # Templates básicos
    templates = {
        "STATE.md": "# Current Project State\n\n## Status Atual\n- [ ] Inicializando projeto...\n",
        "MEMORY.md": "# Project Persistent Memory\n\n## Convenções\n- Adicione aqui padrões e decisões de longo prazo.\n",
        "LEARNINGS.md": "# Project Learnings\n\n## Lições Aprendidas\n- Log de aprendizados e erros evitados.\n",
        "ROADMAP.md": "# Project Roadmap\n\n## Próximos Passos\n- [ ] Definir objetivos do projeto.\n",
    }

    if not specs_project.exists():
        print(f"Criando diretório: {specs_project}")
        specs_project.mkdir(parents=True, exist_ok=True)

    fixed_count = 0
    for filename, content in templates.items():
        file_path = specs_project / filename
        if not file_path.exists():
            print(f"Reparando arquivo ausente: {file_path}")
            file_path.write_text(content)
            fixed_count += 1

    if fixed_count == 0:
        print("✅ Estrutura de memória está íntegra.")
    else:
        print(f"✅ Reparo concluído: {fixed_count} arquivos criados.")


if __name__ == "__main__":
    auto_fix_memory()
