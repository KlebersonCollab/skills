import os
from pathlib import Path

def sync_mandates():
    root_dir = Path(".")
    mandates_file = root_dir / ".specs" / "codebase" / "GLOBAL_MANDATES.md"
    
    if not mandates_file.exists():
        print(f"Error: Mandates file {mandates_file} not found.")
        return

    mandates_content = mandates_file.read_text()
    
    agent_files = [
        root_dir / ".gemini" / "GEMINI.md",
        root_dir / ".claude" / "CLAUDE.md",
        root_dir / ".agent" / "AGENT.md"
    ]

    marker_start = "<!-- GLOBAL_MANDATES_START -->"
    marker_end = "<!-- GLOBAL_MANDATES_END -->"
    
    new_block = f"{marker_start}\n\n{mandates_content}\n\n{marker_end}"

    for agent_file in agent_files:
        if agent_file.exists():
            content = agent_file.read_text()
            
            # Se já existem os marcadores, substituir o bloco
            if marker_start in content and marker_end in content:
                pattern = f"{marker_start}.*?{marker_end}"
                new_content = re.sub(pattern, new_block, content, flags=re.DOTALL)
            else:
                # Se não existem, adicionar ao final ou substituir seções antigas
                # Para facilitar a primeira vez, vamos apenas sobrescrever com um cabeçalho fixo + mandatos
                agent_name = agent_file.parent.name.replace(".", "").upper()
                new_content = f"# {agent_name} Project Instructions\n\n{new_block}\n\n--- \n*Mandato atualizado em 16 de Abril de 2026.*"
            
            agent_file.write_text(new_content)
            print(f"Synced mandates to {agent_file}")

if __name__ == "__main__":
    import re
    sync_mandates()
