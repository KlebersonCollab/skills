import re
import sys
from pathlib import Path
from typing import Dict, List

# Add current directory to path to import utils
sys.path.append(str(Path(__file__).parent))
from utils import get_all_skills, get_skill_metadata


def generate_knowledge_map():
    root_dir = Path(".")
    output_file = root_dir / "KNOWLEDGE-MAP.mermaid"

    skills = get_all_skills(root_dir)

    # Category mapping for readable subgraph names
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

    # Colors per category (grouped)
    style_map = {
        "Architecture & Design": "fill:#f3e5f5,stroke:#4a148c",  # Purple
        "Ecosystems & DevOps": "fill:#e8f5e9,stroke:#1b5e20",  # Green
        "Core Frameworks": "fill:#e1f5fe,stroke:#01579b",  # Blue
        "Automation & Utils": "fill:#fff3e0,stroke:#e65100",  # Orange
        "Navigation & Orchestration": "fill:#fce4ec,stroke:#880e4f",  # Pink
    }

    nodes = []
    relations = set()
    subgraphs: Dict[str, List[str]] = {}

    # 1. Skill Processing
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

        # Group for Subgraphs
        if category_name not in subgraphs:
            subgraphs[category_name] = []
        subgraphs[category_name].append(skill_id)

        # 2. Relation Detection (Explicit via Metadata)
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

        # 3. Relation Detection (Implicit via Mentions - Fallback)
        for other in skills:
            if other.name != skill_path.name:
                # Search for mention of the other skill name with word boundaries
                pattern = rf"\b{re.escape(other.name)}\b"
                if re.search(pattern, skill_md_content):
                    other_id = other.name.replace("-", "_")
                    # If an explicit relation doesn't already exist, mark as implicit
                    explicit_rel = f"{skill_id} --(references)--> {other_id}"
                    explicit_use = f"{skill_id} --(uses)--> {other_id}"
                    if explicit_rel not in relations and explicit_use not in relations:
                        relations.add(explicit_rel)

    # 4. Mermaid Construction
    mermaid = ["graph TD", "    %% Nodes and Hierarchical Grouping"]

    for cat_name, skill_ids in subgraphs.items():
        mermaid.append(f'\n    subgraph "{cat_name}"')
        for s_id in skill_ids:
            # Find the correct label for the ID
            node_data = next(n for n in nodes if n["id"] == s_id)
            mermaid.append(f"        {s_id}[{node_data['label']}]")
        mermaid.append("    end")

    mermaid.append("\n    %% Relations")
    # Prioritize (uses) over (references) to clean the graph
    final_relations = set()
    for rel in relations:
        # If (uses) exists, (references) is redundant
        if "--(references)-->" in rel:
            base = rel.replace("--(references)-->", "--(uses)-->")
            if base in relations:
                continue
        final_relations.add(rel)

    for rel in sorted(list(final_relations)):
        mermaid.append(f"    {rel}")

    # 5. Subgraph Styling
    mermaid.append("\n    %% Styling")
    for cat_name, style in style_map.items():
        # Mermaid doesn't easily support direct styling of subgraphs by name in all versions,
        # so we apply it to the nodes within it.
        if cat_name in subgraphs:
            for s_id in subgraphs[cat_name]:
                mermaid.append(f"    style {s_id} {style}")

    # 6. Legend
    mermaid.append("\n    %% Legend")
    mermaid.append('    subgraph "Legend"')
    mermaid.append('        L1["🛡️ SDD Compliant"]')
    mermaid.append('        L2["Uses: Direct Dependency"]')
    mermaid.append('        L3["References: Documentation Mention"]')
    mermaid.append("    end")

    # 7. Save
    output_file.write_text("\n".join(mermaid) + "\n")
    print(f"✅ Knowledge Map (v3 - Distilled) generated at {output_file}")


if __name__ == "__main__":
    generate_knowledge_map()
