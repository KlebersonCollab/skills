"""
Shared modules for Skills Hub scripts.
Centralizes navigation logic, safe file reading, and metadata extraction.
"""

import logging
import re
import shutil
import tempfile
from pathlib import Path
from typing import List, Set, Optional, Dict, Any

import yaml

# Standard logging configuration
logging.basicConfig(
    level=logging.INFO, format="%(asctime)s - %(levelname)s - %(message)s"
)
logger = logging.getLogger(__name__)

# Directories that should never be treated as skills
DEFAULT_EXCLUDES: Set[str] = {
    ".git",
    ".github",
    ".specs",
    "scripts",
    "node_modules",
    ".venv",
    ".claude",
    ".gemini",
    ".agent",
    "dist_staging",
    "artifacts",
    "brain",
    "scratch",
}


def get_all_skills(
    root: Path = Path("."), excludes: Optional[Set[str]] = None
) -> List[Path]:
    """
    Returns a list of Paths to directories containing a valid skill.
    A valid skill is identified by the presence of a SKILL.md file.
    """
    excludes = excludes or DEFAULT_EXCLUDES
    # Ensure root is a Path object
    root_path = Path(root)

    if not root_path.exists():
        logger.warning(f"Root directory not found: {root_path}")
        return []

    skills = []
    try:
        for d in root_path.iterdir():
            if d.is_dir() and d.name not in excludes:
                if (d / "SKILL.md").exists():
                    skills.append(d)
    except PermissionError as e:
        logger.error(f"Permission denied listing directories in {root_path}: {e}")
    except Exception as e:
        logger.error(f"Error listing skills in {root_path}: {e}")

    return sorted(skills, key=lambda x: x.name)


def safe_read_file(path: Path, encoding: str = "utf-8") -> Optional[str]:
    """Reads file content with exception handling."""
    try:
        return path.read_text(encoding=encoding)
    except FileNotFoundError:
        logger.error(f"File not found: {path}")
        return None
    except PermissionError:
        logger.error(f"Permission denied reading file: {path}")
        return None
    except Exception as e:
        logger.error(f"Unexpected error reading file {path}: {e}")
        return None


def get_skill_metadata(skill_path: Path) -> Dict[str, Any]:
    """
    Extracts metadata from the YAML frontmatter of the SKILL.md file.
    Returns an empty dictionary if it fails.
    """
    skill_md = skill_path / "SKILL.md"
    content = safe_read_file(skill_md)
    if not content:
        return {}

    try:
        # Search for the block between ---
        fm_match = re.search(r"^---\n(.*?)\n---", content, re.DOTALL)
        if fm_match:
            data = yaml.safe_load(fm_match.group(1))
            return data if isinstance(data, dict) else {}
    except Exception as e:
        logger.error(f"Error parsing YAML in {skill_md}: {e}")

    return {}


def safe_copytree(src: Path, dst: Path) -> bool:
    """
    Copies a directory tree atomically/safely.
    Uses a temporary directory to ensure that failures mid-copy
    do not leave the destination in an inconsistent state.
    """
    if not src.exists():
        logger.error(f"Source not found: {src}")
        return False

    try:
        with tempfile.TemporaryDirectory() as tmpdir:
            tmp_dst = Path(tmpdir) / src.name
            shutil.copytree(src, tmp_dst)

            # If destination already exists, remove it safely
            if dst.exists():
                shutil.rmtree(dst)

            # Ensure the destination's parent directory exists
            dst.parent.mkdir(parents=True, exist_ok=True)

            # Move the temporary directory to the final destination
            shutil.move(str(tmp_dst), str(dst))
            return True
    except Exception as e:
        logger.error(f"Failed safe copy from {src} to {dst}: {e}")
        return False
