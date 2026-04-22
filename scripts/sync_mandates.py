import re
import sys
from pathlib import Path

# Adiciona o diretório atual ao path para importar utils
sys.path.append(str(Path(__file__).parent))
from utils import get_all_skills, safe_copytree, safe_read_file, logger


def sync_all():
    root_dir = Path(".")
    mandates_file = root_dir / ".specs" / "codebase" / "GLOBAL_MANDATES.md"

    mandates_content = safe_read_file(mandates_file)
    if mandates_content is None:
        logger.error(f"Mandatos globais não encontrados em {mandates_file}")
        return

    # 1. Configuração dos Agentes
    agent_configs = [
        {"name": "GEMINI", "folder": ".gemini", "file": ".gemini/GEMINI.md"},
        {"name": "CLAUDE", "folder": ".claude", "file": ".claude/CLAUDE.md"},
        {"name": "AGENT", "folder": ".agent", "file": ".agent/AGENT.md"},
    ]

    # 2. Identificação dinâmica das skills na raiz usando utils
    skills = get_all_skills(root_dir)

    marker_start = "<!-- GLOBAL_MANDATES_START -->"
    marker_end = "<!-- GLOBAL_MANDATES_END -->"
    new_block = f"{marker_start}\n\n{mandates_content}\n\n{marker_end}"

    for config in agent_configs:
        agent_dir = root_dir / config["folder"]
        agent_file = root_dir / config["file"]

        if not agent_dir.exists():
            continue

        # 2.1 Sincronizar Mandatos no arquivo principal
        if agent_file.exists():
            content = safe_read_file(agent_file)
            if content and marker_start in content and marker_end in content:
                pattern = f"{marker_start}.*?{marker_end}"
                new_content = re.sub(pattern, new_block, content, flags=re.DOTALL)
            else:
                new_content = f"# {config['name']} Project Instructions\n\n{new_block}\n\n--- \n*Mandato atualizado em Abril de 2026.*"

            agent_file.write_text(new_content)
            logger.info(f"Synced mandates to {agent_file}")

        # 2.2 Sincronizar Diretórios de Skills
        agent_skills_dir = agent_dir / "skills"
        agent_skills_dir.mkdir(exist_ok=True, parents=True)

        for skill_path in skills:
            skill_name = skill_path.name
            dst = agent_skills_dir / skill_name

            # Usa a cópia segura modularizada em utils.py
            if safe_copytree(skill_path, dst):
                logger.info(
                    f"   + Synced skill: {skill_name} -> {config['folder']}/skills/"
                )
            else:
                logger.error(f"   ! Failed to sync skill: {skill_name}")


if __name__ == "__main__":
    sync_all()
