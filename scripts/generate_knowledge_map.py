import re
import sys
from pathlib import Path
from typing import Dict, List

# Adiciona o diretório atual ao path para importar utils
sys.path.append(str(Path(__file__).parent))
from utils import get_all_skills, get_skill_metadata


def generate_knowledge_map():
    root_dir = Path(".")
    output_file = root_dir / "KNOWLEDGE-MAP.mermaid"

    skills = get_all_skills(root_dir)

    # Mapeamento de categorias para nomes legíveis de subgrafos
    category_names = {
        "api-design": "Architecture & Design",
        "architecture": "Architecture & Design",
        "devops-automation": "Ecosystems & DevOps",
        "site-reliability-engineering": "Ecosystems & DevOps",
        "development-python": "Ecosystems & DevOps",
        "development": "Ecosystems & DevOps",
        "development-workflow": "Core Frameworks",
        "automation": "Automation & Utils",
        "orchestration": "Navigation & Orchestration",
        "design-thinking": "Core Frameworks",
        "code-quality": "Architecture & Design",
        "project-management": "Core Frameworks",
        "skill-management": "Core Frameworks",
        "knowledge-management": "Core Frameworks",
    }

    # Cores por categoria (agrupadas)
    style_map = {
        "Architecture & Design": "fill:#f3e5f5,stroke:#4a148c",  # Roxo
        "Ecosystems & DevOps": "fill:#e8f5e9,stroke:#1b5e20",  # Verde
        "Core Frameworks": "fill:#e1f5fe,stroke:#01579b",  # Azul
        "Automation & Utils": "fill:#fff3e0,stroke:#e65100",  # Laranja
        "Navigation & Orchestration": "fill:#fce4ec,stroke:#880e4f",  # Rosa
    }

    nodes = []
    relations = set()
    subgraphs: Dict[str, List[str]] = {}

    # 1. Processamento de Skills
    for skill_path in skills:
        meta = get_skill_metadata(skill_path)
        skill_md_content = (skill_path / "SKILL.md").read_text()

        name = meta.get("name", skill_path.name)
        version = meta.get("version", "v0.0.0")
        category_slug = meta.get("category", "uncategorized")
        category_name = category_names.get(category_slug, "Miscellaneous")

        skill_id = skill_path.name.replace("-", "_")

        # Badges
        has_sdd = "🔒 Prerequisites (Mandatory)" in skill_md_content
        badge = " 🛡️" if has_sdd else ""

        label = f'"{name} ({version}){badge}"'
        nodes.append({"id": skill_id, "label": label, "category": category_name})

        # Agrupar para Subgrafos
        if category_name not in subgraphs:
            subgraphs[category_name] = []
        subgraphs[category_name].append(skill_id)

        # 2. Detecção de Relações (Explícitas via Metadata)
        uses = meta.get("uses", [])
        if isinstance(uses, str):
            uses = [uses]

        for dep in uses:
            dep_id = dep.replace("-", "_")
            relations.add(f"{skill_id} --(uses)--> {dep_id}")

        references = meta.get("references", [])
        if isinstance(references, str):
            references = [references]

        for ref in references:
            ref_id = ref.replace("-", "_")
            relations.add(f"{skill_id} --(references)--> {ref_id}")

        # 3. Detecção de Relações (Implícitas via Menções - Fallback)
        for other in skills:
            if other.name != skill_path.name:
                # Procura menção ao nome da outra skill com word boundaries
                pattern = rf"\b{re.escape(other.name)}\b"
                if re.search(pattern, skill_md_content):
                    other_id = other.name.replace("-", "_")
                    # Se já não houver uma relação explícita, a gente marca implícita
                    explicit_rel = f"{skill_id} --(references)--> {other_id}"
                    explicit_use = f"{skill_id} --(uses)--> {other_id}"
                    if explicit_rel not in relations and explicit_use not in relations:
                        relations.add(explicit_rel)

    # 4. Construção do Mermaid
    mermaid = ["graph TD", "    %% Nodes and Hierarchical Grouping"]

    for cat_name, skill_ids in subgraphs.items():
        mermaid.append(f'\n    subgraph "{cat_name}"')
        for s_id in skill_ids:
            # Encontrar o label correto para o ID
            node_data = next(n for n in nodes if n["id"] == s_id)
            mermaid.append(f"        {s_id}[{node_data['label']}]")
        mermaid.append("    end")

    mermaid.append("\n    %% Relations")
    # Priorizar (uses) sobre (references) para limpar o grafo
    final_relations = set()
    for rel in relations:
        # Se existe um (uses), não precisa do (references)
        if "--(references)-->" in rel:
            base = rel.replace("--(references)-->", "--(uses)-->")
            if base in relations:
                continue
        final_relations.add(rel)

    for rel in sorted(list(final_relations)):
        mermaid.append(f"    {rel}")

    # 5. Estilização de Subgrafos
    mermaid.append("\n    %% Styling")
    for cat_name, style in style_map.items():
        # Mermaid não suporta estilo direto em subgraph via nome facilmente em todas versões,
        # então aplicamos aos nós dentro dele.
        if cat_name in subgraphs:
            for s_id in subgraphs[cat_name]:
                mermaid.append(f"    style {s_id} {style}")

    # 6. Legenda
    mermaid.append("\n    %% Legend")
    mermaid.append('    subgraph "Legenda"')
    mermaid.append('        L1["🛡️ SDD Compliant"]')
    mermaid.append('        L2["Uses: Dependência Direta"]')
    mermaid.append('        L3["References: Menção em Docs"]')
    mermaid.append("    end")

    # 7. Salvar
    output_file.write_text("\n".join(mermaid) + "\n")
    print(f"✅ Knowledge Map (v3 - Distilled) gerado em {output_file}")


if __name__ == "__main__":
    generate_knowledge_map()
