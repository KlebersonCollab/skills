import re
import sys
from pathlib import Path

def verify_versions():
    root_dir = Path(".")
    readme_path = root_dir / "README.md"
    
    if not readme_path.exists():
        print("Error: README.md not found.")
        return

    readme_content = readme_path.read_text()
    
    # Encontrar a tabela de skills no README
    # Regex para capturar: | # | **[Name](path/)** | Desc | `version` |
    pattern = r"\| \d+ \| \*\*\[.*?\]\((.*?)/?\) \*\* \| .*? \| `(.*?)` \|"
    matches = re.findall(pattern, readme_content)

    errors = []
    print(f"🔍 Iniciando verificação de consistência de versões...")

    for skill_path, readme_version in matches:
        skill_file = root_dir / skill_path / "SKILL.md"
        
        if not skill_file.exists():
            print(f"   ⚠️  Aviso: Diretório {skill_path} não encontrado (definido no README).")
            continue

        # Ler versão do frontmatter do SKILL.md
        skill_content = skill_file.read_text()
        version_match = re.search(r"version:\s*([^\s\n]+)", skill_content)
        
        if not version_match:
            errors.append(f"❌ {skill_path}: Versão não encontrada no SKILL.md")
            continue
            
        actual_version = version_match.group(1).strip("'\"")
        
        if actual_version != readme_version:
            errors.append(f"❌ {skill_path}: README (`{readme_version}`) != SKILL.md (`{actual_version}`)")
        else:
            print(f"   ✅ {skill_path}: {actual_version}")

    if errors:
        print("\n--- ERROS DE CONSISTÊNCIA ENCONTRADOS ---")
        for error in errors:
            print(error)
        sys.exit(1)
    else:
        print("\n🎉 Todas as versões estão consistentes!")

if __name__ == "__main__":
    verify_versions()
