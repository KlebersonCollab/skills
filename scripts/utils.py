"""
Módulos compartilhados para scripts do Hub de Skills.
Centraliza lógica de navegação, leitura segura de arquivos e extração de metadados.
"""
import logging
import re
import shutil
import tempfile
from pathlib import Path
from typing import List, Set, Optional, Dict, Any

import yaml

# Configuração de logging padrão
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

# Diretórios que nunca devem ser tratados como skills
DEFAULT_EXCLUDES: Set[str] = {
    ".git", ".github", ".specs", "scripts", 
    "node_modules", ".venv", ".claude", ".gemini", ".agent",
    "dist_staging", "artifacts", "brain", "scratch"
}

def get_all_skills(root: Path = Path("."), excludes: Optional[Set[str]] = None) -> List[Path]:
    """
    Retorna uma lista de Paths para diretórios que contêm uma skill válida.
    Uma skill válida é identificada pela presença de um arquivo SKILL.md.
    """
    excludes = excludes or DEFAULT_EXCLUDES
    # Garante que root é um objeto Path
    root_path = Path(root)
    
    if not root_path.exists():
        logger.warning(f"Diretório raiz não encontrado: {root_path}")
        return []

    skills = []
    try:
        for d in root_path.iterdir():
            if d.is_dir() and d.name not in excludes:
                if (d / "SKILL.md").exists():
                    skills.append(d)
    except PermissionError as e:
        logger.error(f"Permissão negada ao listar diretórios em {root_path}: {e}")
    except Exception as e:
        logger.error(f"Erro ao listar skills em {root_path}: {e}")
        
    return sorted(skills, key=lambda x: x.name)

def safe_read_file(path: Path, encoding: str = "utf-8") -> Optional[str]:
    """Lê o conteúdo de um arquivo com tratamento de exceções."""
    try:
        return path.read_text(encoding=encoding)
    except FileNotFoundError:
        logger.error(f"Arquivo não encontrado: {path}")
        return None
    except PermissionError:
        logger.error(f"Permissão negada ao ler arquivo: {path}")
        return None
    except Exception as e:
        logger.error(f"Erro inesperado ao ler arquivo {path}: {e}")
        return None

def get_skill_metadata(skill_path: Path) -> Dict[str, Any]:
    """
    Extrai metadados do frontmatter YAML do arquivo SKILL.md.
    Retorna um dicionário vazio se falhar.
    """
    skill_md = skill_path / "SKILL.md"
    content = safe_read_file(skill_md)
    if not content:
        return {}
    
    try:
        # Busca o bloco entre ---
        fm_match = re.search(r"^---\n(.*?)\n---", content, re.DOTALL)
        if fm_match:
            data = yaml.safe_load(fm_match.group(1))
            return data if isinstance(data, dict) else {}
    except Exception as e:
        logger.error(f"Erro ao parsear YAML em {skill_md}: {e}")
        
    return {}

def safe_copytree(src: Path, dst: Path) -> bool:
    """
    Copia uma árvore de diretórios de forma atômica/segura.
    Usa um diretório temporário para garantir que falhas no meio da cópia
    não deixem o destino em estado inconsistente.
    """
    if not src.exists():
        logger.error(f"Fonte não encontrada: {src}")
        return False

    try:
        with tempfile.TemporaryDirectory() as tmpdir:
            tmp_dst = Path(tmpdir) / src.name
            shutil.copytree(src, tmp_dst)
            
            # Se o destino já existe, remove de forma segura
            if dst.exists():
                shutil.rmtree(dst)
            
            # Garante que o diretório pai do destino existe
            dst.parent.mkdir(parents=True, exist_ok=True)
            
            # Move o diretório temporário para o destino final
            shutil.move(str(tmp_dst), str(dst))
            return True
    except Exception as e:
        logger.error(f"Falha na cópia segura de {src} para {dst}: {e}")
        return False
