import os
import sys
from datetime import datetime, timedelta
from pathlib import Path

# Adiciona o diretório atual ao path para importar utils
sys.path.append(str(Path(__file__).parent))
from utils import logger, DEFAULT_EXCLUDES

def check_memory_sync():
    """
    Verifica se os arquivos de memória (STATE.md, LEARNINGS.md) foram atualizados 
    recentemente em comparação com as alterações de código no repositório.
    """
    root_dir = Path(__file__).parent.parent
    specs_dir = root_dir / ".specs" / "project"
    
    # Arquivos críticos de memória
    memory_files = [
        specs_dir / "STATE.md",
        specs_dir / "LEARNINGS.md",
        specs_dir / "MEMORY.md"
    ]
    
    # 1. Verificar existência
    for f in memory_files:
        if not f.exists():
            logger.error(f"Memory file missing: {f.name}")
            return False

    # 2. Verificar se houve commits de código sem atualização de memória
    now = datetime.now()
    # Pegar data da última modificação dos arquivos de memória
    last_memory_update = max(os.path.getmtime(f) for f in memory_files)
    last_memory_dt = datetime.fromtimestamp(last_memory_update)
    
    logger.info(f"Last Memory Update: {last_memory_dt.strftime('%Y-%m-%d %H:%M:%S')}")

    # Verificar se algum arquivo fora de .specs foi alterado DEPOIS da memória
    unsynced_files = []
    
    # Exclusões baseadas no utils.py + .specs
    local_excludes = DEFAULT_EXCLUDES.copy()
    local_excludes.add(".specs")
    
    for root, dirs, files in os.walk(root_dir):
        # Filtra diretórios excluídos
        dirs[:] = [d for d in dirs if d not in local_excludes]
        
        for file in files:
            file_path = Path(root) / file
            # Monitorar apenas arquivos relevantes de código/doc
            if file_path.suffix in ['.py', '.md', '.js', '.ts', '.tsx', '.go', '.dart']:
                # Ignorar arquivos na raiz como Makefile ou CHANGELOG-HUB
                if file_path.parent == root_dir and file_path.name != "README.md":
                    continue
                    
                mtime = os.path.getmtime(file_path)
                if mtime > last_memory_update:
                    unsynced_files.append(file_path.relative_to(root_dir))

    if unsynced_files:
        print("\n⚠️  WARNING: CONTEXT DRIFT DETECTED!")
        print("Os seguintes arquivos foram alterados mas a MEMÓRIA (STATE/LEARNINGS) não foi atualizada:")
        for f in unsynced_files[:10]:
            print(f"  - {f}")
        if len(unsynced_files) > 10:
            print(f"  ... e mais {len(unsynced_files) - 10} arquivos.")
        
        print("\n👉 Ação Requerida: Atualize os arquivos em .specs/project/ para refletir o estado atual.")
        return False
    
    print("✅ Memory is in sync with codebase changes.")
    return True

if __name__ == "__main__":
    success = check_memory_sync()
    if not success:
        # Apenas avisa por padrão
        pass
