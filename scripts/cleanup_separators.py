import os
from pathlib import Path
import re

def cleanup_skills():
    root = Path(".")
    for skill_dir in root.iterdir():
        if not skill_dir.is_dir(): continue
        skill_file = skill_dir / "SKILL.md"
        if not skill_file.exists(): continue
        
        content = skill_file.read_text()
        # Fix double separators like ---\n\n---\n\n
        content = re.sub(r"---\s*\n\s*---\s*\n+", "---\n\n", content)
        # Fix Prerequisites block double separators if any
        # content = re.sub(r"---\s*\n\s*## 🔒 Prerequisites", "---\n\n## 🔒 Prerequisites", content)
        
        skill_file.write_text(content)

if __name__ == "__main__":
    cleanup_skills()
