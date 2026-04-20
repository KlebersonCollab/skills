import os
import re
import yaml
from pathlib import Path

def generate_knowledge_map():
    root_dir = Path(".")
    skills_dir = root_dir
    output_file = root_dir / "KNOWLEDGE-MAP.mermaid"
    
    nodes = []
    relations = set()
    categories = {}

    # Estilos de cores para as categorias reais
    style_map = {
        "api-design": "fill:#f3e5f5,stroke:#4a148c",          # Roxo (Architecture)
        "architecture": "fill:#f3e5f5,stroke:#4a148c",        # Roxo (Architecture)
        "devops-automation": "fill:#e8f5e9,stroke:#1b5e20",   # Verde (DevOps)
        "site-reliability-engineering": "fill:#e8f5e9,stroke:#1b5e20", # Verde (DevOps)
        "development-python": "fill:#fff3e0,stroke:#e65100",  # Laranja (Dev)
        "development": "fill:#fff3e0,stroke:#e65100",         # Laranja (Dev)
        "development-workflow": "fill:#fff3e0,stroke:#e65100",# Laranja (Dev)
        "automation": "fill:#fff3e0,stroke:#e65100",          # Laranja (Dev)
        "orchestration": "fill:#fff3e0,stroke:#e65100",       # Laranja (Dev)
        "design-thinking": "fill:#e1f5fe,stroke:#01579b",     # Azul (Methodology)
        "code-quality": "fill:#fce4ec,stroke:#880e4f",        # Rosa (Quality)
        "project-management": "fill:#fce4ec,stroke:#880e4f",  # Rosa (Mgmt)
        "skill-management": "fill:#fce4ec,stroke:#880e4f",    # Rosa (Mgmt)
        "knowledge-management": "fill:#e0f7fa,stroke:#006064" # Ciano (Knowledge)
    }


    exclude_dirs = {".git", ".github", ".specs", "scripts", "node_modules", ".venv", ".claude", ".gemini", ".agent"}
    skills_paths = [d for d in skills_dir.iterdir() if d.is_dir() and d.name not in exclude_dirs and (d / "SKILL.md").exists()]
    
    # 1. Extração de Dados e Criação de Nós
    for skill in skills_paths:
        skill_md = skill / "SKILL.md"
        content = skill_md.read_text()
        
        # Extrair frontmatter
        fm_match = re.search(r"^---\n(.*?)\n---", content, re.DOTALL)
        if fm_match:
            data = yaml.safe_load(fm_match.group(1))
            name = data.get("name", skill.name)
            version = data.get("version", "v0.0.0")
            category = data.get("category", "uncategorized")
        else:
            name = skill.name
            version = "v0.0.0"
            category = "uncategorized"

        skill_id = skill.name.replace("-", "_")
        
        # Verificar se tem o Hook Mandatório do SDD
        has_sdd = "🔒 Prerequisites (Mandatory)" in content
        badge = " 🛡️" if has_sdd else ""
        
        label = f'"{name} ({version}){badge}"'
        nodes.append({"id": skill_id, "label": label, "category": category})
        
        if category not in categories:
            categories[category] = []
        categories[category].append(skill_id)

        # 2. Detecção de Relações
        # Procurar por menções a outras skills
        for other in skills_paths:
            if other.name != skill.name and (f"/{other.name}/" in content or f" {other.name} " in content):
                other_id = other.name.replace("-", "_")
                relations.add(f"{skill_id} --(uses)--> {other_id}")

    # 3. Construção do Mermaid
    mermaid = ["graph TD", "    %% Nodes and Styles"]
    
    for node in nodes:
        mermaid.append(f'    {node["id"]}[{node["label"]}]')
    
    mermaid.append("\n    %% Categories Styling")
    for cat, ids in categories.items():
        style = style_map.get(cat, "fill:#f5f5f5,stroke:#333")
        for node_id in ids:
            mermaid.append(f"    style {node_id} {style}")

    mermaid.append("\n    %% Relations")
    for rel in sorted(list(relations)):
        mermaid.append(f"    {rel}")

    # 4. Salvar
    output_file.write_text("\n".join(mermaid) + "\n")
    print(f"✅ Knowledge Map (v2) gerado em {output_file}")

if __name__ == "__main__":
    generate_knowledge_map()
