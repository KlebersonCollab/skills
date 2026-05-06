import os
import re

def extract_metadata(path):
    if not os.path.exists(path):
        return None
    
    with open(path, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # Pattern para o bloco de metadados SDD
    # Busca blocos que NÃO estão dentro de outro bloco de código (evita templates em documentação)
    # Procuramos pela última ocorrência que pareça ser um estado real
    pattern = r'<!-- @sdd-state -->\s*```yaml\s*(.*?)\s*```'
    matches = re.findall(pattern, content, re.DOTALL)
    
    if matches:
        # Filtramos para pegar o último que não seja um template
        for yaml_content in reversed(matches):
            metadata = {}
            for line in yaml_content.split('\n'):
                if ':' in line:
                    key, value = line.split(':', 1)
                    metadata[key.strip()] = value.strip().strip('"').strip("'")
            
            # Se não for um template (contém placeholders), retornamos este
            if metadata.get('feature_id') != 'FEAT-ID' and metadata.get('last_update') != 'ISO-TIMESTAMP':
                return metadata
            
    return None

def get_targets():
    targets = []
    
    # 1. Arquivos na raiz
    root_files = ["README.md", "PROJECT-ONBOARDING.md"]
    for f in root_files:
        if os.path.exists(f):
            targets.append(f)
            
    # 2. Specs de Projeto
    spec_project = ".specs/project"
    if os.path.exists(spec_project):
        for f in sorted(os.listdir(spec_project)):
            if f.endswith(".md"):
                targets.append(os.path.join(spec_project, f))

    # 3. Codebase Specs
    spec_codebase = ".specs/codebase"
    if os.path.exists(spec_codebase):
        for f in sorted(os.listdir(spec_codebase)):
            if f.endswith(".md"):
                targets.append(os.path.join(spec_codebase, f))

    # 4. Descoberta Dinâmica de Skills
    # Varre todos os diretórios em busca de SKILL.md
    for root, dirs, files in os.walk("."):
        # Ignorar diretórios ocultos ou de sistema
        dirs[:] = [d for d in dirs if not d.startswith('.') and d not in ['bin', 'examples', 'resources', 'node_modules', 'venv']]
        
        if "SKILL.md" in files:
            # Encontramos uma skill!
            # Adicionamos os arquivos padrão de uma skill
            skill_files = ["SKILL.md", "README.md", "CHANGELOG.md"]
            for sf in skill_files:
                path = os.path.join(root, sf)
                if os.path.exists(path) and path not in targets:
                    targets.append(path)
            
            # Adicionamos arquivos .skill.md e referências
            for f in files:
                if f.endswith(".skill.md"):
                    path = os.path.join(root, f)
                    if path not in targets:
                        targets.append(path)
            
            # Checar referências se existirem
            ref_path = os.path.join(root, "references")
            if os.path.exists(ref_path):
                for rf in sorted(os.listdir(ref_path)):
                    if rf.endswith(".md"):
                        path = os.path.join(ref_path, rf)
                        if path not in targets:
                            targets.append(path)

    return sorted(list(set(targets)))

if __name__ == "__main__":
    print("--- 🛡️ SDD v2.3.0 Observable Governance Audit ---")
    
    targets = get_targets()
    
    print(f"{'FILE':<60} | {'VER':<8} | {'STATUS':<12}")
    print("-" * 85)
    
    errors = 0
    for target in targets:
        # Normaliza o path para exibição (remove ./ inicial)
        display_path = target[2:] if target.startswith("./") else target
        
        meta = extract_metadata(target)
        if meta:
            ver = meta.get('version', '???')
            status = meta.get('status', '???')
            print(f"✅ {display_path:<57} | {ver:<8} | {status:<12}")
        else:
            # Se o arquivo existe mas não tem metadados, é falha
            print(f"❌ {display_path:<57} | {'MISSING':<8} | {'FAILED':<12}")
            errors += 1
    
    print("\n" + "="*85)
    if errors == 0:
        print("🏆 RESULT: 100% COMPLIANT (v2.3.0)")
    else:
        print(f"⚠️ RESULT: {errors} FILES FAILED AUDIT")
    print("="*85)
