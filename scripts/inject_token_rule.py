import os
import re

# Rule to inject
RULE = "0. **Mode Check**: Verificar o modo operacional atual (`.hub-mode`) e aplicar as diretrizes da skill `token-distiller`."

def inject_rule(file_path):
    with open(file_path, 'r', encoding='utf-8') as f:
        content = f.read()

    # Search for Prerequisites section
    if "🔒 Prerequisites (Mandatory)" in content:
        # Avoid duplicate injection
        if "token-distiller" in content:
            print(f"Skipping {file_path} - already updated.")
            return

        # Inject after the header
        new_content = re.sub(
            r"(## 🔒 Prerequisites \(Mandatory\)\nEsta skill opera DENTRO do framework \*\*SDD\*\*. Antes de iniciar qualquer execução técnica:\n)",
            r"\g<1>" + RULE + "\n",
            content
        )
        
        with open(file_path, 'w', encoding='utf-8') as f:
            f.write(new_content)
        print(f"Updated {file_path}")
    else:
        print(f"Prerequisites section not found in {file_path}")

# List of root skill directories
skills = [
    "api-architect", "architecture", "azure-devops", "benchmark-expert",
    "brainstorming", "clean-code-mentor", "devsecops-expert", "django-expert",
    "fastapi-expert", "flutter-fvm", "frontend-expert", "git-workflow",
    "golang-expert", "golang-testing-expert", "harness-expert", "knowledge-architect",
    "observability-expert", "python-uv", "scaffolding-expert", "sdd",
    "skill-factory", "swarm-facilitator", "youtube-transcript"
]

for skill in skills:
    skill_path = os.path.join(skill, "SKILL.md")
    if os.path.exists(skill_path):
        inject_rule(skill_path)
