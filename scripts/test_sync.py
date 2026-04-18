import os
import shutil
import subprocess
from pathlib import Path

def test_sync_functionality():
    root_dir = Path(".")
    test_skill_name = "test-sync-skill"
    test_skill_dir = root_dir / test_skill_name
    
    print(f"🚀 Iniciando teste de integração para o script de sincronização...")

    try:
        # 1. Criar uma skill de teste temporária
        test_skill_dir.mkdir(exist_ok=True)
        (test_skill_dir / "SKILL.md").write_text("---\nname: test-sync-skill\nversion: 0.1.0\n---")
        print(f"   - Skill de teste criada: {test_skill_name}")

        # 2. Rodar o script de sincronização
        print(f"   - Executando scripts/sync_mandates.py...")
        subprocess.run(["python3", "scripts/sync_mandates.py"], check=True, capture_output=True)

        # 3. Verificar se a skill foi propagada para os agentes
        agents = [".agent", ".gemini", ".claude"]
        for agent in agents:
            dst = root_dir / agent / "skills" / test_skill_name
            if not dst.exists():
                raise AssertionError(f"❌ Falha: Skill não encontrada em {agent}/skills/")
            if not (dst / "SKILL.md").exists():
                raise AssertionError(f"❌ Falha: SKILL.md ausente em {agent}/skills/")
            print(f"   ✅ Verificado em {agent}/skills/")

        print(f"🎉 Teste de sincronização concluído com SUCESSO!")

    finally:
        # 4. Limpeza
        if test_skill_dir.exists():
            shutil.rmtree(test_skill_dir)
        for agent in [".agent", ".gemini", ".claude"]:
            dst = root_dir / agent / "skills" / test_skill_name
            if dst.exists():
                shutil.rmtree(dst)
        print(f"   - Limpeza concluída.")

if __name__ == "__main__":
    test_sync_functionality()
