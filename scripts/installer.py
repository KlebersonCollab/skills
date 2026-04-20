#!/usr/bin/env python3
"""
Skill Installer CLI
Permite a instalação granular de skills do Hub.
"""
import argparse
import sys
from pathlib import Path

# Adiciona o diretório atual ao path para importar utils
sys.path.append(str(Path(__file__).parent))
from utils import get_all_skills, safe_copytree, get_skill_metadata, logger

def list_skills():
    """Lista todas as skills disponíveis no Hub."""
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

def install_skill(name: str, target: Path, force: bool = False):
    """Instala uma skill específica no diretório alvo."""
    skills = get_all_skills()
    skill_to_install = next((s for s in skills if s.name == name or get_skill_metadata(s).get("name") == name), None)
    
    if not skill_to_install:
        print(f"❌ Erro: Skill '{name}' não encontrada no Hub.")
        sys.exit(1)
        
    dest_path = target / skill_to_install.name
    
    if dest_path.exists() and not force:
        print(f"⚠️  Aviso: A skill '{name}' já existe em {dest_path}. Use --force para sobrescrever.")
        sys.exit(1)
        
    print(f"🚀 Instalando {name} em {dest_path}...")
    if safe_copytree(skill_to_install, dest_path):
        print(f"✅ Skill '{name}' instalada com sucesso!")
    else:
        print(f"❌ Falha ao instalar skill '{name}'.")
        sys.exit(1)

def main():
    parser = argparse.ArgumentParser(description="Skill Installer CLI")
    subparsers = parser.add_subparsers(dest="command", help="Comandos disponíveis")
    
    # Comando List
    subparsers.add_parser("list", help="Lista todas as skills do Hub")
    
    # Comando Install
    install_parser = subparsers.add_parser("install", help="Instala uma skill")
    install_parser.add_argument("name", help="Nome da skill para instalar")
    install_parser.add_argument("--target", type=Path, default=Path("./.agent/skills"), help="Diretório de destino (default: ./.agent/skills)")
    install_parser.add_argument("--force", action="store_true", help="Sobrescreve se a skill já existir")
    
    args = parser.parse_args()
    
    if args.command == "list":
        list_skills()
    elif args.command == "install":
        install_skill(args.name, args.target, args.force)
    else:
        parser.print_help()

if __name__ == "__main__":
    main()
