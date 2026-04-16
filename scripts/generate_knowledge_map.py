import os
from pathlib import Path
import re

def generate_knowledge_map():
    root_dir = Path(".")
    features_dir = root_dir / ".specs" / "features"
    skills_dir = root_dir
    output_file = root_dir / "KNOWLEDGE-MAP.mermaid"
    
    nodes = set()
    relations = set()

    # 1. Mapear Skills como nós principais
    exclude_dirs = {".git", ".github", ".specs", "scripts", "node_modules", ".venv", ".claude", ".gemini", ".agent"}
    skills = [d for d in skills_dir.iterdir() if d.is_dir() and d.name not in exclude_dirs and (d / "SKILL.md").exists()]
    
    for skill in skills:
        skill_id = skill.name.replace("-", "_")
        nodes.add(f'{skill_id}["🧩 {skill.name}"]')
        
        # Tentar encontrar relações dentro do SKILL.md (links para outras skills)
        skill_md = skill / "SKILL.md"
        if skill_md.exists():
            content = skill_md.read_text()
            # Procurar por menções a outras skills na mesma pasta
            for other_skill in skills:
                if other_skill.name != skill.name and other_skill.name in content:
                    other_id = other_skill.name.replace("-", "_")
                    relations.add(f"{skill_id} --(uses)--> {other_id}")

    # 2. Mapear Features em .specs/features
    if features_dir.exists():
        for feature in features_dir.iterdir():
            if feature.is_dir():
                feat_id = f"feat_{feature.name.replace('-', '_')}"
                nodes.add(f'{feat_id}["🚀 Feature: {feature.name}"]')
                
                # Relacionar feature com skills mencionadas na spec ou plan
                for doc in ["spec.md", "plan.md", "tasks.md"]:
                    doc_path = feature / doc
                    if doc_path.exists():
                        doc_content = doc_path.read_text()
                        for skill in skills:
                            if skill.name in doc_content:
                                skill_id = skill.name.replace("-", "_")
                                relations.add(f"{feat_id} --(depends_on)--> {skill_id}")

    # 3. Gerar conteúdo Mermaid
    mermaid_content = ["graph TD", "    %% Nodes"]
    for node in sorted(list(nodes)):
        mermaid_content.append(f"    {node}")
    
    mermaid_content.append("\n    %% Relations")
    for rel in sorted(list(relations)):
        mermaid_content.append(f"    {rel}")

    # 4. Salvar arquivo
    output_file.write_text("\n".join(mermaid_content) + "\n")
    print(f"Knowledge Map generated successfully at {output_file}")

if __name__ == "__main__":
    generate_knowledge_map()
