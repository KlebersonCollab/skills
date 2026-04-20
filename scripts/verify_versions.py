import re
import sys
from pathlib import Path

# Adiciona o diretório atual ao path para importar utils
sys.path.append(str(Path(__file__).parent))
from utils import safe_read_file, get_skill_metadata, logger

def verify_versions():
    root_dir = Path(".")
    readme_path = root_dir / "README.md"
    
    readme_content = safe_read_file(readme_path)
    if readme_content is None:
        logger.error("README.md não encontrado.")
        return

    # Regex para capturar: | # | **[Name](path/)** | Desc | `version` |
    pattern = r"\| \d+ \| \*\*\[.*?\]\((.*?)/?\) \*\* \| .*? \| `(.*?)` \|"
    matches = re.findall(pattern, readme_content)

    errors = []
    logger.info("Iniciando verificação de consistência de versões...")

    for skill_path_str, readme_version in matches:
        skill_path = root_dir / skill_path_str
        metadata = get_skill_metadata(skill_path)
        
        if not metadata:
            logger.warning(f"Metadata não encontrada para {skill_path_str}")
            continue
            
        actual_version = str(metadata.get("version", "")).strip("'\"")
        
        if actual_version != readme_version:
            errors.append(f"❌ {skill_path_str}: README (`{readme_version}`) != SKILL.md (`{actual_version}`)")
        else:
            logger.info(f"   ✅ {skill_path_str}: {actual_version}")

    if errors:
        print("\n--- ERROS DE CONSISTÊNCIA ENCONTRADOS ---")
        for error in errors:
            print(error)
        sys.exit(1)
    else:
        print("\n🎉 Todas as versões estão consistentes!")

if __name__ == "__main__":
    verify_versions()
