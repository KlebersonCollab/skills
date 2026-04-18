import os
import re
import shutil
from pathlib import Path

def sync_all():
    root_dir = Path(".")
    mandates_file = root_dir / ".specs" / "codebase" / "GLOBAL_MANDATES.md"
    
    if not mandates_file.exists():
        print(f"Error: Mandates file {mandates_file} not found.")
        return

    mandates_content = mandates_file.read_text()
    
    # 1. Configuração dos Agentes
    agent_configs = [
        {"name": "GEMINI", "folder": ".gemini", "file": ".gemini/GEMINI.md"},
        {"name": "CLAUDE", "folder": ".claude", "file": ".claude/CLAUDE.md"},
        {"name": "AGENT", "folder": ".agent", "file": ".agent/AGENT.md"}
    ]

    # 2. Identificação dinâmica das skills na raiz
    skills = [
        d.name for d in root_dir.iterdir() 
        if d.is_dir() and (d / "SKILL.md").exists()
    ]
    
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
            content = agent_file.read_text()
            if marker_start in content and marker_end in content:
                pattern = f"{marker_start}.*?{marker_end}"
                new_content = re.sub(pattern, new_block, content, flags=re.DOTALL)
            else:
                new_content = f"# {config['name']} Project Instructions\n\n{new_block}\n\n--- \n*Mandato atualizado em 18 de Abril de 2026.*"
            agent_file.write_text(new_content)
            print(f"✅ Synced mandates to {agent_file}")

        # 2.2 Sincronizar Diretórios de Skills
        agent_skills_dir = agent_dir / "skills"
        agent_skills_dir.mkdir(exist_ok=True)
        
        for skill in skills:
            src = root_dir / skill
            dst = agent_skills_dir / skill
            
            # Se a skill não existe no agente ou é mais antiga, copiar
            # Para garantir integridade, vamos remover e copiar de novo (estilo mirror)
            if dst.exists():
                shutil.rmtree(dst)
            shutil.copytree(src, dst)
            print(f"   + Synced skill: {skill} -> {config['folder']}/skills/")

if __name__ == "__main__":
    sync_all()
