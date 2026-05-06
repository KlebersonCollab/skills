import os
import re

def extract_metadata(path):
    if not os.path.exists(path):
        return None
    
    with open(path, 'r', encoding='utf-8') as f:
        content = f.read()
    
    # Pattern to find SDD state blocks
    # We look for the marker and the following yaml block
    pattern = r'<!-- @sdd-state -->\s*```yaml\s*(.*?)\s*```'
    matches = re.finditer(pattern, content, re.DOTALL)
    
    valid_metadata = None
    
    for match in matches:
        start_idx = match.start()
        
        # Security check: verify this block is NOT inside another code block
        # We count the number of ``` before the match
        prefix = content[:start_idx]
        if prefix.count('```') % 2 != 0:
            # We are likely inside a code block (odd number of fences before us)
            continue
            
        yaml_content = match.group(1)
        metadata = {}
        for line in yaml_content.split('\n'):
            if ':' in line:
                key, value = line.split(':', 1)
                metadata[key.strip()] = value.strip().strip('"').strip("'")
        
        # Check if it's a real state (not a template)
        if metadata.get('feature_id') != 'FEAT-ID' and metadata.get('last_update') != 'ISO-TIMESTAMP':
            valid_metadata = metadata
            
    return valid_metadata

def get_targets():
    targets = []
    
    # 1. Root and Core Artifacts
    core_dirs = [".specs/project", ".specs/codebase"]
    root_files = ["README.md", "PROJECT-ONBOARDING.md"]
    
    for f in root_files:
        if os.path.exists(f):
            targets.append(f)
            
    for d in core_dirs:
        if os.path.exists(d):
            for f in sorted(os.listdir(d)):
                if f.endswith(".md"):
                    targets.append(os.path.join(d, f))

    # 2. Dynamic Skill Discovery
    exclude_dirs = {'.git', 'bin', 'node_modules', 'venv', '.agents', '.gemini', '.cursor'}
    
    for root, dirs, files in os.walk("."):
        # Prune excluded directories
        dirs[:] = [d for d in dirs if d not in exclude_dirs and not d.startswith('.')]
        
        if any(excl in root.split(os.sep) for excl in exclude_dirs):
            continue

        if "SKILL.md" in files:
            # Found a skill directory
            for sf in ["SKILL.md", "README.md", "CHANGELOG.md"]:
                path = os.path.join(root, sf)
                if os.path.exists(path):
                    targets.append(path)
            
            # Sub-skills
            for f in files:
                if f.endswith(".skill.md"):
                    targets.append(os.path.join(root, f))
            
            # References
            ref_path = os.path.join(root, "references")
            if os.path.exists(ref_path):
                for rf in sorted(os.listdir(ref_path)):
                    if rf.endswith(".md"):
                        targets.append(os.path.join(ref_path, rf))

    # De-duplicate and normalize
    seen = set()
    unique_targets = []
    for t in targets:
        nt = os.path.normpath(t)
        if nt.startswith("./"):
            nt = nt[2:]
        if nt not in seen:
            unique_targets.append(nt)
            seen.add(nt)
            
    return unique_targets

if __name__ == "__main__":
    print("--- 🛡️ SDD v2.3.0 Observable Governance Audit ---")
    
    targets = get_targets()
    
    print(f"{'FILE':<60} | {'VER':<8} | {'STATUS':<12}")
    print("-" * 85)
    
    errors = 0
    for target in targets:
        meta = extract_metadata(target)
        if meta:
            ver = meta.get('version', '???')
            status = meta.get('status', '???')
            print(f"✅ {target:<57} | {ver:<8} | {status:<12}")
        else:
            print(f"❌ {target:<57} | {'MISSING':<8} | {'FAILED':<12}")
            errors += 1
    
    print("\n" + "="*85)
    if errors == 0:
        print("🏆 RESULT: 100% COMPLIANT (v2.3.0)")
    else:
        print(f"⚠️ RESULT: {errors} FILES FAILED AUDIT")
    print("="*85)
