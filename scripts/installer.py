#!/usr/bin/env python3
"""
Skill Installer CLI
Allows granular installation of Hub skills (local or remote).
"""

import argparse
import sys
import subprocess
import tempfile
from pathlib import Path

# Add current directory to path to import utils
sys.path.append(str(Path(__file__).parent))
from utils import get_all_skills, safe_copytree, get_skill_metadata


def list_skills():
    """Lists all available skills in the local Hub."""
    skills = get_all_skills()
    print("\n📦 Available Skills in Hub:\n")
    print(f"{'Name':<25} | {'Version':<10} | {'Category':<15}")
    print("-" * 55)

    for skill_path in skills:
        meta = get_skill_metadata(skill_path)
        name = meta.get("name", skill_path.name)
        version = meta.get("version", "v0.0.0")
        category = meta.get("category", "uncategorized")

        print(f"{name:<25} | {version:<10} | {category:<15}")
    print(f"\nTotal: {len(skills)} skills.\n")


def install_remote_skill(
    name: str,
    target: Path,
    force: bool = False,
    repo_url: str = "https://github.com/KlebersonCollab/skills.git",
):
    """Downloads and installs a skill using Git Sparse-Checkout for granular extraction."""
    dest_path = target / name

    if dest_path.exists() and not force:
        print(
            f"⚠️  Warning: Skill '{name}' already exists in {dest_path}. Use --force to overwrite."
        )
        sys.exit(1)

    print(f"🚀 Downloading '{name}' from remote repository via Git Sparse-Checkout...")
    try:
        with tempfile.TemporaryDirectory() as tmpdir:
            # Clone without downloading blobs (optimized)
            cmd_clone = [
                "git",
                "clone",
                "--filter=blob:none",
                "--sparse",
                repo_url,
                tmpdir,
            ]
            res = subprocess.run(cmd_clone, capture_output=True, text=True)
            if res.returncode != 0:
                print(f"❌ Error cloning remote repository:\n{res.stderr}")
                sys.exit(1)

            # Add possible paths to sparse-checkout
            cmd_sparse = [
                "git",
                "-C",
                tmpdir,
                "sparse-checkout",
                "add",
                name,
                f".agents/skills/{name}",
                f".agent/skills/{name}",
            ]
            res = subprocess.run(cmd_sparse, capture_output=True, text=True)
            if res.returncode != 0:
                print(f"❌ Error in sparse-checkout:\n{res.stderr}")
                sys.exit(1)

            # Identify where the skill was actually downloaded
            src_skill = None
            possible_paths = [
                Path(tmpdir) / name,
                Path(tmpdir) / ".agents/skills" / name,
                Path(tmpdir) / ".agent/skills" / name,
            ]

            for path in possible_paths:
                if (path / "SKILL.md").exists():
                    src_skill = path
                    break

            if not src_skill:
                print(f"❌ Error: Skill '{name}' not found in remote repository.")
                sys.exit(1)

            print(f"🚀 Installing {name} in {dest_path}...")
            if safe_copytree(src_skill, dest_path):
                print(f"✅ Skill '{name}' successfully installed from remote!")
            else:
                print(f"❌ Failed to copy skill '{name}'.")
                sys.exit(1)

    except Exception as e:
        print(f"❌ Unexpected failure during remote download: {e}")
        sys.exit(1)


def install_local_skill(name: str, target: Path, force: bool = False):
    """Installs a skill from the local repository."""
    skills = get_all_skills()
    skill_to_install = next(
        (
            s
            for s in skills
            if s.name == name or get_skill_metadata(s).get("name") == name
        ),
        None,
    )

    if not skill_to_install:
        print(f"❌ Error: Skill '{name}' not found in local Hub.")
        print(
            "💡 Tip: Try using --remote to fetch the skill from the official repository."
        )
        sys.exit(1)

    dest_path = target / skill_to_install.name

    if dest_path.exists() and not force:
        print(
            f"⚠️  Warning: Skill '{name}' already exists in {dest_path}. Use --force to overwrite."
        )
        sys.exit(1)

    print(f"🚀 Installing {name} in {dest_path}...")
    if safe_copytree(skill_to_install, dest_path):
        print(f"✅ Skill '{name}' successfully installed!")
    else:
        print(f"❌ Failed to install skill '{name}'.")
        sys.exit(1)


def install_skill(name: str, target: Path, force: bool = False, remote: bool = False):
    if remote:
        install_remote_skill(name, target, force)
    else:
        install_local_skill(name, target, force)


def main():
    parser = argparse.ArgumentParser(description="Skill Installer CLI")
    subparsers = parser.add_subparsers(dest="command", help="Available commands")

    # List Command
    subparsers.add_parser("list", help="Lists all skills in the local Hub")

    # Install Command
    install_parser = subparsers.add_parser("install", help="Installs a skill")
    install_parser.add_argument("name", help="Name of the skill to install")
    install_parser.add_argument(
        "--target",
        type=Path,
        default=Path("./.agent/skills"),
        help="Target directory (default: ./.agent/skills)",
    )
    install_parser.add_argument(
        "--force", action="store_true", help="Overwrite if skill already exists"
    )
    install_parser.add_argument(
        "--remote",
        action="store_true",
        help="Download directly from remote repository via Git Sparse-Checkout",
    )

    args = parser.parse_args()

    if args.command == "list":
        list_skills()
    elif args.command == "install":
        install_skill(args.name, args.target, args.force, args.remote)
    else:
        parser.print_help()


if __name__ == "__main__":
    main()
