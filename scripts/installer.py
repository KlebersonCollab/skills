#!/usr/bin/env python3
"""
Skill Installer CLI
Permite a instalação granular de skills do Hub (local ou remoto).
"""

import argparse
import sys
import subprocess
import tempfile
from pathlib import Path

# Adiciona o diretório atual ao path para importar utils
sys.path.append(str(Path(__file__).parent))
from utils import get_all_skills, safe_copytree, get_skill_metadata


def list_skills():
    """Lista todas as skills disponíveis no Hub local."""
    skills = get_all_skills()
    print("\n📦 Skills Disponíveis no Hub:\n")
    print(f"{'Nome':<25} | {'Versão':<10} | {'Categoria':<15}")
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
    """Baixa e instala uma skill usando Git Sparse-Checkout para extração granular."""
    dest_path = target / name

    if dest_path.exists() and not force:
        print(
            f"⚠️  Aviso: A skill '{name}' já existe em {dest_path}. Use --force para sobrescrever."
        )
        sys.exit(1)

    print(f"🚀 Baixando '{name}' do repositório remoto via Git Sparse-Checkout...")
    try:
        with tempfile.TemporaryDirectory() as tmpdir:
            # Clone sem baixar blobs (otimizado)
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
                print(f"❌ Erro ao clonar o repositório remoto:\n{res.stderr}")
                sys.exit(1)

            # Adiciona os caminhos possíveis ao sparse-checkout
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
                print(f"❌ Erro no sparse-checkout:\n{res.stderr}")
                sys.exit(1)

            # Identifica onde a skill foi efetivamente baixada
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
                print(f"❌ Erro: Skill '{name}' não encontrada no repositório remoto.")
                sys.exit(1)

            print(f"🚀 Instalando {name} em {dest_path}...")
            if safe_copytree(src_skill, dest_path):
                print(f"✅ Skill '{name}' instalada com sucesso a partir do remoto!")
            else:
                print(f"❌ Falha ao copiar a skill '{name}'.")
                sys.exit(1)

    except Exception as e:
        print(f"❌ Falha inesperada durante download remoto: {e}")
        sys.exit(1)


def install_local_skill(name: str, target: Path, force: bool = False):
    """Instala uma skill a partir do repositório local."""
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
        print(f"❌ Erro: Skill '{name}' não encontrada no Hub local.")
        print(
            "💡 Dica: Tente usar --remote para buscar a skill no repositório oficial."
        )
        sys.exit(1)

    dest_path = target / skill_to_install.name

    if dest_path.exists() and not force:
        print(
            f"⚠️  Aviso: A skill '{name}' já existe em {dest_path}. Use --force para sobrescrever."
        )
        sys.exit(1)

    print(f"🚀 Instalando {name} em {dest_path}...")
    if safe_copytree(skill_to_install, dest_path):
        print(f"✅ Skill '{name}' instalada com sucesso!")
    else:
        print(f"❌ Falha ao instalar skill '{name}'.")
        sys.exit(1)


def install_skill(name: str, target: Path, force: bool = False, remote: bool = False):
    if remote:
        install_remote_skill(name, target, force)
    else:
        install_local_skill(name, target, force)


def main():
    parser = argparse.ArgumentParser(description="Skill Installer CLI")
    subparsers = parser.add_subparsers(dest="command", help="Comandos disponíveis")

    # Comando List
    subparsers.add_parser("list", help="Lista todas as skills do Hub local")

    # Comando Install
    install_parser = subparsers.add_parser("install", help="Instala uma skill")
    install_parser.add_argument("name", help="Nome da skill para instalar")
    install_parser.add_argument(
        "--target",
        type=Path,
        default=Path("./.agent/skills"),
        help="Diretório de destino (default: ./.agent/skills)",
    )
    install_parser.add_argument(
        "--force", action="store_true", help="Sobrescreve se a skill já existir"
    )
    install_parser.add_argument(
        "--remote",
        action="store_true",
        help="Faz o download diretamente do repositório remoto via Git Sparse-Checkout",
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
