import os
import sys
import re
import yaml
from pathlib import Path

def validate_skill(skill_path):
    skill_dir = Path(skill_path)
    skill_name = skill_dir.name
    results = {}
    errors = []

    # Check 1: Structural Integrity
    mandatory_files = ["SKILL.md", "README.md", "CHANGELOG.md"]
    missing_files = [f for f in mandatory_files if not (skill_dir / f).exists()]
    results["Structural Integrity"] = "PASS" if not missing_files else "FAIL"
    if missing_files:
        errors.append(f"Missing mandatory files: {', '.join(missing_files)}")

    # Check 2 & 3: Frontmatter and Content Rigor
    skill_md_path = skill_dir / "SKILL.md"
    if skill_md_path.exists():
        content = skill_md_path.read_text()
        
        # Frontmatter
        try:
            frontmatter_match = re.search(r"^---\n(.*?)\n---", content, re.DOTALL)
            if frontmatter_match:
                data = yaml.safe_load(frontmatter_match.group(1))
                if data.get("name") != skill_name:
                    errors.append(f"Frontmatter name '{data.get('name')}' does not match directory '{skill_name}'")
                if not re.match(r"^\d+\.\d+\.\d+$", str(data.get("version", ""))):
                    errors.append(f"Invalid version format: {data.get('version')}")
            else:
                errors.append("Missing YAML frontmatter in SKILL.md")
        except Exception as e:
            errors.append(f"Error parsing frontmatter: {str(e)}")

        # H2 Rigor
        h2_sections = ["## Goal", "## Output Structure", "## Quality Rules", "## Prohibited"]
        missing_h2 = [h2 for h2 in h2_sections if h2 not in content]
        if missing_h2:
            errors.append(f"Missing H2 sections in SKILL.md: {', '.join(missing_h2)}")
            results["Content Completeness"] = "FAIL"
        else:
            results["Content Completeness"] = "PASS"

    # Check 4: Version Sync
    changelog_path = skill_dir / "CHANGELOG.md"
    if changelog_path.exists() and skill_md_path.exists():
        changelog_content = changelog_path.read_text()
        version_match = re.search(r"## \[(.*?)\]", changelog_content)
        if version_match:
            latest_version = version_match.group(1)
            # Re-read frontmatter for sync
            frontmatter_match = re.search(r"^---\n(.*?)\n---", skill_md_path.read_text(), re.DOTALL)
            if frontmatter_match:
                data = yaml.safe_load(frontmatter_match.group(1))
                if str(data.get("version")) != latest_version:
                    errors.append(f"Version mismatch: SKILL.md ({data.get('version')}) vs CHANGELOG.md ({latest_version})")

    # Check 5: Recommended Directories
    recommended_dirs = ["references", "resources", "examples"]
    missing_dirs = [d for d in recommended_dirs if not (skill_dir / d).exists()]
    if missing_dirs:
        # Warning instead of error (don't fail the build for this yet)
        print(f"\n    \033[93mWarning\033[0m: Missing recommended directories in {skill_name}: {', '.join(missing_dirs)}", end="")

    # Check 6: Changelog Format (## [Version] - YYYY-MM-DD)
    changelog_path = skill_dir / "CHANGELOG.md"
    if changelog_path.exists():
        changelog_content = changelog_path.read_text()
        # Find the first version entry
        version_entry = re.search(r"## \[(.*?)\]( - \d{4}-\d{2}-\d{2})?", changelog_content)
        if version_entry and not version_entry.group(2):
            errors.append(f"Missing date in CHANGELOG.md for version [{version_entry.group(1)}]. Expected: ## [{version_entry.group(1)}] - YYYY-MM-DD")

    # Final Verdict for this skill
    status = "PASS" if not errors else "FAIL"
    return status, errors

def main():
    root_dir = Path(".")
    # Ignorar pastas que começam com ponto ou não são skills
    exclude = [".git", ".github", ".specs", "scripts", "node_modules", ".venv"]
    
    skills = [d for d in root_dir.iterdir() if d.is_dir() and d.name not in exclude and (d / "SKILL.md").exists()]
    
    total_failed = 0
    for skill in skills:
        print(f"Validating skill: {skill.name}...", end=" ")
        status, errors = validate_skill(skill)
        if status == "PASS":
            print("\033[92mPASS\033[0m")
        else:
            print("\033[91mFAIL\033[0m")
            total_failed += 1
            for err in errors:
                print(f"  - {err}")
    
    if total_failed > 0:
        print(f"\n\033[91mValidation failed for {total_failed} skill(s).\033[0m")
        sys.exit(1)
    else:
        print("\n\033[92mAll skills are compliant!\033[0m")
        sys.exit(0)

if __name__ == "__main__":
    main()
