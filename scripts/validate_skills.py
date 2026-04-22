"""
Validador de Skills com arquitetura SOLID.
Permite a adição de novos checks de auditoria de forma modular.
"""
import re
import sys
from abc import ABC, abstractmethod
from dataclasses import dataclass, field
from pathlib import Path
from typing import List, Dict, Any

import yaml

# Adiciona o diretório atual ao path para importar utils
sys.path.append(str(Path(__file__).parent))
from utils import get_all_skills, safe_read_file, get_skill_metadata, logger

@dataclass
class ValidationResult:
    """Representa o resultado de um check individual."""
    check_name: str
    is_valid: bool
    errors: List[str] = field(default_factory=list)
    warnings: List[str] = field(default_factory=list)

class BaseValidator(ABC):
    """Classe base para todos os validadores de skills."""
    @abstractmethod
    def validate(self, skill_path: Path) -> ValidationResult:
        pass

class StructuralValidator(BaseValidator):
    """Valida se os arquivos obrigatórios existem."""
    MANDATORY_FILES = ["SKILL.md", "README.md", "CHANGELOG.md"]
    
    def validate(self, skill_path: Path) -> ValidationResult:
        missing = [f for f in self.MANDATORY_FILES if not (skill_path / f).exists()]
        errors = [f"Arquivo obrigatório ausente: {f}" for f in missing]
        return ValidationResult("Integridade Estrutural", not errors, errors)

class FrontmatterValidator(BaseValidator):
    """Valida o frontmatter YAML do SKILL.md."""
    def validate(self, skill_path: Path) -> ValidationResult:
        metadata = get_skill_metadata(skill_path)
        errors = []
        
        if not metadata:
            return ValidationResult("YAML Frontmatter", False, ["Frontmatter ausente ou corrompido em SKILL.md"])
        
        # Valida nome
        if metadata.get("name") != skill_path.name:
            errors.append(f"Nome no frontmatter ({metadata.get('name')}) não bate com diretório ({skill_path.name})")
            
        # Valida versão (formato semver)
        version = str(metadata.get("version", ""))
        if not re.match(r"^\d+\.\d+\.\d+$", version):
            errors.append(f"Formato de versão inválido: {version}. Use X.Y.Z")
            
        return ValidationResult("YAML Frontmatter", not errors, errors)

class ContentCompletenessValidator(BaseValidator):
    """Valida a presença de seções H2 obrigatórias no SKILL.md."""
    MANDATORY_SECTIONS = ["## Goal", "## Output Structure", "## Quality Rules", "## Prohibited"]
    
    def validate(self, skill_path: Path) -> ValidationResult:
        content = safe_read_file(skill_path / "SKILL.md")
        if not content:
            return ValidationResult("Completude de Conteúdo", False, ["Não foi possível ler SKILL.md"])
            
        missing = [sec for sec in self.MANDATORY_SECTIONS if sec not in content]
        errors = [f"Seção H2 ausente no SKILL.md: {sec}" for sec in missing]
        
        return ValidationResult("Completude de Conteúdo", not errors, errors)

class VersionSyncValidator(BaseValidator):
    """Valida se a versão no SKILL.md bate com a última versão no CHANGELOG.md."""
    def validate(self, skill_path: Path) -> ValidationResult:
        metadata = get_skill_metadata(skill_path)
        changelog_content = safe_read_file(skill_path / "CHANGELOG.md")
        
        if not metadata or not changelog_content:
            return ValidationResult("Sincronia de Versão", True) # Outros validadores pegam falha de arquivo
            
        errors = []
        # Busca a versão mais recente no changelog: ## [X.Y.Z]
        version_match = re.search(r"## \[(.*?)\]", changelog_content)
        if version_match:
            latest_changelog_version = version_match.group(1)
            skill_version = str(metadata.get("version"))
            
            if skill_version != latest_changelog_version:
                errors.append(f"Versão divergente: SKILL.md ({skill_version}) vs CHANGELOG.md ({latest_changelog_version})")
        else:
            errors.append("Nenhuma versão encontrada no CHANGELOG.md (formato esperado: ## [X.Y.Z])")
            
        return ValidationResult("Sincronia de Versão", not errors, errors)

class ChangelogFormatValidator(BaseValidator):
    """Valida o formato da data no CHANGELOG.md."""
    def validate(self, skill_path: Path) -> ValidationResult:
        content = safe_read_file(skill_path / "CHANGELOG.md")
        if not content:
            return ValidationResult("Formato do Changelog", True)
            
        errors = []
        # Busca entrada de versão e data: ## [X.Y.Z] - YYYY-MM-DD
        # O grupo 2 captura a data opcionalmente para verificar sua existência
        version_entry = re.search(r"## \[(.*?)\]( - \d{4}-\d{2}-\d{2})?", content)
        
        if version_entry and not version_entry.group(2):
            version = version_entry.group(1)
            errors.append(f"Data ausente no CHANGELOG.md para a versão [{version}]. Esperado: ## [{version}] - YYYY-MM-DD")
            
        return ValidationResult("Formato do Changelog", not errors, errors)

class GovernanceValidator(BaseValidator):
    """Valida a adesão aos mandatos de governança (Full SDD Hook)."""
    MANDATORY_HOOK_ITEMS = [
        "Context Check",
        "Spec Check",
        "Plan Check",
        "Contract Check",
        "Task Check"
    ]
    
    def validate(self, skill_path: Path) -> ValidationResult:
        content = safe_read_file(skill_path / "SKILL.md")
        if not content:
            return ValidationResult("Governança", True)
            
        errors = []
        if "🔒 Prerequisites (Mandatory)" not in content:
            errors.append("Seção obrigatória ausente: ## 🔒 Prerequisites (Mandatory)")
        else:
            for item in self.MANDATORY_HOOK_ITEMS:
                if item not in content:
                    errors.append(f"Item de governança ausente no hook: {item}")
            
        return ValidationResult("Governança", not errors, errors)

class RecommendedStructureValidator(BaseValidator):
    """Verifica pastas recomendadas mas não obrigatórias."""
    RECOMMENDED_DIRS = ["references", "resources", "examples"]
    
    def validate(self, skill_path: Path) -> ValidationResult:
        missing = [d for d in self.RECOMMENDED_DIRS if not (skill_path / d).exists()]
        warnings = [f"Pasta recomendada ausente: {d}" for d in missing]
        return ValidationResult("Estrutura Recomendada", True, warnings=warnings)

class SkillAuditor:
    """Orquestrador que executa todos os validadores."""
    def __init__(self):
        self.validators: List[BaseValidator] = [
            StructuralValidator(),
            FrontmatterValidator(),
            ContentCompletenessValidator(),
            GovernanceValidator(),
            VersionSyncValidator(),
            ChangelogFormatValidator(),
            RecommendedStructureValidator()
        ]
        
    def audit_skill(self, skill_path: Path) -> List[ValidationResult]:
        return [v.validate(skill_path) for v in self.validators]

def main():
    root_dir = Path(".")
    skills = get_all_skills(root_dir)
    auditor = SkillAuditor()
    
    total_failed = 0
    
    for skill in skills:
        print(f"Validating skill: {skill.name}...", end=" ", flush=True)
        results = auditor.audit_skill(skill)
        
        is_skill_valid = all(r.is_valid and not r.errors for r in results)
        
        if is_skill_valid:
            print("\033[92mPASS\033[0m")
        else:
            print("\033[91mFAIL\033[0m")
            total_failed += 1
            
        # Exibe Erros e Warnings
        for r in results:
            for err in r.errors:
                print(f"  - \033[91m[Erro]\033[0m {r.check_name}: {err}")
            for warn in r.warnings:
                print(f"  - \033[93m[Aviso]\033[0m {r.check_name}: {warn}")
    
    if total_failed > 0:
        print(f"\n\033[91mAuditoria falhou para {total_failed} skill(s).\033[0m")
        sys.exit(1)
    else:
        print("\n\033[92mHub está saudável! Todas as skills passaram nos checks mandatórios.\033[0m")
        sys.exit(0)

if __name__ == "__main__":
    main()
