#!/usr/bin/env python3
import os
import re
import yaml
import sys

def audit_files():
    violations = []
    project_root = os.getcwd()
    specs_dir = os.path.join(project_root, ".specs")
    
    # 1. Check for Triad of Memory
    triad_files = ["STATE.md", "MEMORY.md", "LEARNINGS.md"]
    triad_path = os.path.join(specs_dir, "project")
    for f in triad_files:
        if not os.path.exists(os.path.join(triad_path, f)):
            violations.append(f"CRITICAL: Missing Triad file {f} in .specs/project/")

    # 2. Scan for Markdown files and validate metadata
    for root, dirs, files in os.walk(project_root):
        # Ignore common directories
        if any(d in root for d in [".git", "node_modules", "bin", "docs"]):
            continue
            
        for file in files:
            if file.endswith(".md"):
                file_path = os.path.join(root, file)
                with open(file_path, "r", encoding="utf-8") as f:
                    content = f.read()
                    
                    # Check for metadata block
                    if "<!-- @sdd-state -->" not in content:
                        # Only enforce on .specs and skill folders
                        if ".specs" in file_path or any(os.path.exists(os.path.join(root, "SKILL.md")) for root, dirs, files in os.walk(root)):
                             violations.append(f"METADATA: Missing <!-- @sdd-state --> in {os.path.relpath(file_path)}")
                        continue
                    
                    # Extract YAML
                    try:
                        match = re.search(r"```yaml\n(.*?)\n```", content.split("<!-- @sdd-state -->")[1], re.DOTALL)
                        if match:
                            data = yaml.safe_load(match.group(1))
                            
                            # Validation: status vs evidence
                            if data.get("status") == "COMPLETED" and data.get("evidence_checksum") == "NONE":
                                # Exception for project-wide rehydration files which might not have evidence yet
                                if "project/" not in file_path:
                                    violations.append(f"EVIDENCE: {os.path.relpath(file_path)} is COMPLETED but has evidence_checksum: NONE")
                    except Exception as e:
                        violations.append(f"PARSE: Failed to parse metadata in {os.path.relpath(file_path)}: {e}")

    return violations

if __name__ == "__main__":
    print("🔍 Starting Governance Audit...")
    results = audit_files()
    if results:
        print("\n❌ Violations found:")
        for v in results:
            print(f"  - {v}")
        sys.exit(1)
    else:
        print("\n✅ All mandates are compliant.")
        sys.exit(0)
