from pathlib import Path


def auto_fix_memory(target_path="."):
    """
    Ensures the memory and state structure (.specs/project/) exists
    with basic templates if files are missing.
    """
    root = Path(target_path)
    specs_project = root / ".specs" / "project"

    # Basic templates
    templates = {
        "STATE.md": "# Current Project State\n\n## Current Status\n- [ ] Initializing project...\n",
        "MEMORY.md": "# Project Persistent Memory\n\n## Conventions\n- Add long-term patterns and decisions here.\n",
        "LEARNINGS.md": "# Project Learnings\n\n## Lessons Learned\n- Log of learnings and avoided errors.\n",
        "ROADMAP.md": "# Project Roadmap\n\n## Next Steps\n- [ ] Define project objectives.\n",
    }

    if not specs_project.exists():
        print(f"Creating directory: {specs_project}")
        specs_project.mkdir(parents=True, exist_ok=True)

    fixed_count = 0
    for filename, content in templates.items():
        file_path = specs_project / filename
        if not file_path.exists():
            print(f"Repairing missing file: {file_path}")
            file_path.write_text(content)
            fixed_count += 1

    if fixed_count == 0:
        print("✅ Memory structure is healthy.")
    else:
        print(f"✅ Repair completed: {fixed_count} files created.")


if __name__ == "__main__":
    auto_fix_memory()
