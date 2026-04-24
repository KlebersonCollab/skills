import os
import sys
from datetime import datetime
from pathlib import Path

# Add current directory to path to import utils
sys.path.append(str(Path(__file__).parent))
from utils import logger, DEFAULT_EXCLUDES


def check_memory_sync():
    """
    Checks if memory files (STATE.md, LEARNINGS.md) have been updated
    recently compared to code changes in the repository.
    """
    root_dir = Path(__file__).parent.parent
    specs_dir = root_dir / ".specs" / "project"

    # Critical memory files
    memory_files = [
        specs_dir / "STATE.md",
        specs_dir / "LEARNINGS.md",
        specs_dir / "MEMORY.md",
    ]

    # 1. Verify existence
    for f in memory_files:
        if not f.exists():
            logger.error(f"Memory file missing: {f.name}")
            return False

    # 2. Check if code commits occurred without memory updates
    # Get last modification time of memory files
    last_memory_update = max(os.path.getmtime(f) for f in memory_files)
    last_memory_dt = datetime.fromtimestamp(last_memory_update)

    logger.info(f"Last Memory Update: {last_memory_dt.strftime('%Y-%m-%d %H:%M:%S')}")

    # Check if any file outside .specs was changed AFTER memory files
    unsynced_files = []

    # Exclusions based on utils.py + .specs
    local_excludes = DEFAULT_EXCLUDES.copy()
    local_excludes.add(".specs")

    for root, dirs, files in os.walk(root_dir):
        # Filter excluded directories
        dirs[:] = [d for d in dirs if d not in local_excludes]

        for file in files:
            file_path = Path(root) / file
            # Monitor only relevant code/doc files
            if file_path.suffix in [".py", ".md", ".js", ".ts", ".tsx", ".go", ".dart"]:
                # Ignore root files like Makefile or CHANGELOG-HUB
                if file_path.parent == root_dir and file_path.name != "README.md":
                    continue

                mtime = os.path.getmtime(file_path)
                if mtime > last_memory_update:
                    unsynced_files.append(file_path.relative_to(root_dir))

    if unsynced_files:
        print("\n⚠️  WARNING: CONTEXT DRIFT DETECTED!")
        print(
            "The following files were modified but MEMORY (STATE/LEARNINGS) was not updated:"
        )
        for f in unsynced_files[:10]:
            print(f"  - {f}")
        if len(unsynced_files) > 10:
            print(f"  ... and {len(unsynced_files) - 10} more files.")

        print(
            "\n👉 Required Action: Update files in .specs/project/ to reflect current state."
        )
        return False

    print("✅ Memory is in sync with codebase changes.")
    return True


if __name__ == "__main__":
    success = check_memory_sync()
    if not success:
        # Warning only by default
        pass
