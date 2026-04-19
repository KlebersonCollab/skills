import os
from pathlib import Path
import re
from datetime import datetime

def consolidate_changelogs():
    root_dir = Path(".")
    exclude = [".git", ".github", ".specs", "scripts", "node_modules", ".venv", "dist_staging", "artifacts"]
    skills = [d for d in root_dir.iterdir() if d.is_dir() and d.name not in exclude and (d / "CHANGELOG.md").exists()]
    
    hub_changelog = [
        "# Hub Global Changelog",
        f"*Última atualização: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}*",
        "",
        "Este arquivo consolida as atualizações mais recentes de todas as skills do Hub.",
        ""
    ]
    
    for skill in sorted(skills):
        changelog_path = skill / "CHANGELOG.md"
        content = changelog_path.read_text()
        
        # Extract the 3 most recent versions/sections (H2 entries)
        entries = re.split(r"(?=## \[)", content)
        # Entry 0 is usually header text, skip it if it doesn't match
        recent_entries = [e for e in entries if e.startswith("## [")][:3]
        
        if recent_entries:
            hub_changelog.append(f"## 🧩 {skill.name}")
            for entry in recent_entries:
                hub_changelog.append(entry.strip())
                hub_changelog.append("")
            hub_changelog.append("---")
            hub_changelog.append("")

    Path("CHANGELOG-HUB.md").write_text("\n".join(hub_changelog))
    print(f"✅ CHANGELOG-HUB.md gerado com sucesso!")

if __name__ == "__main__":
    consolidate_changelogs()
