"""
Skill Validator with SOLID architecture.
Allows adding new audit checks in a modular way.
"""

import re
import sys
from abc import ABC, abstractmethod
from dataclasses import dataclass, field
from pathlib import Path
from typing import List


# Add current directory to path to import utils
sys.path.append(str(Path(__file__).parent))
from utils import get_all_skills, safe_read_file, get_skill_metadata


@dataclass
class ValidationResult:
    """Represents the result of an individual check."""

    check_name: str
    is_valid: bool
    errors: List[str] = field(default_factory=list)
    warnings: List[str] = field(default_factory=list)


class BaseValidator(ABC):
    """Base class for all skill validators."""

    @abstractmethod
    def validate(self, skill_path: Path) -> ValidationResult:
        pass


class StructuralValidator(BaseValidator):
    """Validates if mandatory files exist."""

    MANDATORY_FILES = ["SKILL.md", "README.md", "CHANGELOG.md"]

    def validate(self, skill_path: Path) -> ValidationResult:
        missing = [f for f in self.MANDATORY_FILES if not (skill_path / f).exists()]
        errors = [f"Mandatory file missing: {f}" for f in missing]
        return ValidationResult("Structural Integrity", not errors, errors)


class FrontmatterValidator(BaseValidator):
    """Validates the YAML frontmatter in SKILL.md."""

    def validate(self, skill_path: Path) -> ValidationResult:
        metadata = get_skill_metadata(skill_path)
        errors = []

        if not metadata:
            return ValidationResult(
                "YAML Frontmatter",
                False,
                ["Frontmatter missing or corrupted in SKILL.md"],
            )

        # Validate name
        if metadata.get("name") != skill_path.name:
            errors.append(
                f"Frontmatter name ({metadata.get('name')}) does not match directory name ({skill_path.name})"
            )

        # Validate version (semver format)
        version = str(metadata.get("version", ""))
        if not re.match(r"^\d+\.\d+\.\d+$", version):
            errors.append(f"Invalid version format: {version}. Use X.Y.Z")

        return ValidationResult("YAML Frontmatter", not errors, errors)


class ContentCompletenessValidator(BaseValidator):
    """Validates the presence of mandatory H2 sections in SKILL.md."""

    MANDATORY_SECTIONS = [
        "## Goal",
        "## Output Structure",
        "## Quality Rules",
        "## Prohibited",
    ]

    def validate(self, skill_path: Path) -> ValidationResult:
        content = safe_read_file(skill_path / "SKILL.md")
        if not content:
            return ValidationResult(
                "Content Completeness", False, ["Could not read SKILL.md"]
            )

        missing = [sec for sec in self.MANDATORY_SECTIONS if sec not in content]
        errors = [f"H2 section missing in SKILL.md: {sec}" for sec in missing]

        return ValidationResult("Content Completeness", not errors, errors)


class VersionSyncValidator(BaseValidator):
    """Validates if the version in SKILL.md matches the latest version in CHANGELOG.md."""

    def validate(self, skill_path: Path) -> ValidationResult:
        metadata = get_skill_metadata(skill_path)
        changelog_content = safe_read_file(skill_path / "CHANGELOG.md")

        if not metadata or not changelog_content:
            return ValidationResult(
                "Version Sync", True
            )  # Other validators will catch missing files

        errors = []
        # Search for the most recent version in changelog: ## [X.Y.Z]
        version_match = re.search(r"## \[(.*?)\]", changelog_content)
        if version_match:
            latest_changelog_version = version_match.group(1)
            skill_version = str(metadata.get("version"))

            if skill_version != latest_changelog_version:
                errors.append(
                    f"Version mismatch: SKILL.md ({skill_version}) vs CHANGELOG.md ({latest_changelog_version})"
                )
        else:
            errors.append(
                "No version found in CHANGELOG.md (expected format: ## [X.Y.Z])"
            )

        return ValidationResult("Version Sync", not errors, errors)


class ChangelogFormatValidator(BaseValidator):
    """Validates the date format in CHANGELOG.md."""

    def validate(self, skill_path: Path) -> ValidationResult:
        content = safe_read_file(skill_path / "CHANGELOG.md")
        if not content:
            return ValidationResult("Changelog Format", True)

        errors = []
        # Search for version entry and date: ## [X.Y.Z] - YYYY-MM-DD
        # Group 2 optionally captures the date to verify its existence
        version_entry = re.search(r"## \[(.*?)\]( - \d{4}-\d{2}-\d{2})?", content)

        if version_entry and not version_entry.group(2):
            version = version_entry.group(1)
            errors.append(
                f"Date missing in CHANGELOG.md for version [{version}]. Expected: ## [{version}] - YYYY-MM-DD"
            )

        return ValidationResult("Changelog Format", not errors, errors)


class GovernanceValidator(BaseValidator):
    """Validates adherence to governance mandates (Full SDD Hook)."""

    MANDATORY_HOOK_ITEMS = [
        "Context Check",
        "Spec Check",
        "Plan Check",
        "Contract Check",
        "Task Check",
    ]

    def validate(self, skill_path: Path) -> ValidationResult:
        content = safe_read_file(skill_path / "SKILL.md")
        if not content:
            return ValidationResult("Governance", True)

        errors = []
        if "🔒 Prerequisites (Mandatory)" not in content:
            errors.append("Mandatory section missing: ## 🔒 Prerequisites (Mandatory)")
        else:
            for item in self.MANDATORY_HOOK_ITEMS:
                if item not in content:
                    errors.append(f"Governance item missing in hook: {item}")

        return ValidationResult("Governance", not errors, errors)


class RecommendedStructureValidator(BaseValidator):
    """Checks for recommended but optional folders."""

    RECOMMENDED_DIRS = ["references", "resources", "examples"]

    def validate(self, skill_path: Path) -> ValidationResult:
        missing = [d for d in self.RECOMMENDED_DIRS if not (skill_path / d).exists()]
        warnings = [f"Recommended folder missing: {d}" for d in missing]
        return ValidationResult("Recommended Structure", True, warnings=warnings)


class SkillAuditor:
    """Orchestrator that executes all validators."""

    def __init__(self):
        self.validators: List[BaseValidator] = [
            StructuralValidator(),
            FrontmatterValidator(),
            ContentCompletenessValidator(),
            GovernanceValidator(),
            VersionSyncValidator(),
            ChangelogFormatValidator(),
            RecommendedStructureValidator(),
        ]

    def audit_skill(self, skill_path: Path) -> List[ValidationResult]:
        return [v.validate(skill_path) for v in self.validators]


def main():
    root_dir = Path(".")
    skills = get_all_skills(root_dir)
    auditor = SkillAuditor()

    total_failed = 0

    for skill in skills:
        print(f"Validating skill: {skill.name}...", end=" ", flush=True)
        results = auditor.audit_skill(skill)

        is_skill_valid = all(r.is_valid and not r.errors for r in results)

        if is_skill_valid:
            print("\033[92mPASS\033[0m")
        else:
            print("\033[91mFAIL\033[0m")
            total_failed += 1

        # Display Errors and Warnings
        for r in results:
            for err in r.errors:
                print(f"  - \033[91m[Error]\033[0m {r.check_name}: {err}")
            for warn in r.warnings:
                print(f"  - \033[93m[Warning]\033[0m {r.check_name}: {warn}")

    if total_failed > 0:
        print(f"\n\033[91mAudit failed for {total_failed} skill(s).\033[0m")
        sys.exit(1)
    else:
        print(
            "\n\033[92mHub is healthy! All skills passed mandatory checks.\033[0m"
        )
        sys.exit(0)


if __name__ == "__main__":
    main()
