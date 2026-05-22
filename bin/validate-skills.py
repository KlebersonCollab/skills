#!/usr/bin/env python3
"""
validate-skills.py — Valida todos os .skill.md do ecossistema.

Verificações:
- Bloco @sdd-state presente e com campos obrigatórios.
- Versão consistente (2.3.0).
- evidence_checksum não vazio.
- Parameters schema válido (se presente).
- Links internos existentes.

Nota: Tenta usar PyYAML; se indisponível, usa parser manual com suporte a multilinha.
      O ÚLTIMO bloco @sdd-state do arquivo é considerado o oficial.
"""

import os
import re
import sys

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
EXPECTED_VERSION = "2.3.0"
SKILL_DIRS = [d for d in os.listdir(ROOT)
              if os.path.isdir(os.path.join(ROOT, d)) and not d.startswith('.')]

# Tenta importar yaml
try:
    import yaml
    HAS_YAML = True
except ImportError:
    HAS_YAML = False

errors = []
warnings = []


def find_skill_files(base_dir):
    """Retorna lista de caminhos de arquivos .skill.md dentro de um diretório de skill."""
    skill_files = []
    for root, dirs, files in os.walk(base_dir):
        dirs[:] = [d for d in dirs if d not in ('.venv', '__pycache__', '.git', 'node_modules')]
        for f in files:
            if f.endswith('.skill.md'):
                skill_files.append(os.path.join(root, f))
    return skill_files


def parse_yaml_block(text):
    """
    Parsing manual de bloco YAML simples (chave: valor).
    Suporta valores multilinha (linhas seguintes indentadas).
    Remove aspas.
    """
    result = {}
    current_key = None
    current_value_lines = []
    in_multiline = False
    for line in text.splitlines():
        stripped = line.strip()
        if not stripped or stripped.startswith('#'):
            if in_multiline and not stripped:
                # linha vazia dentro do valor multilinha
                current_value_lines.append('')
            continue
        # Verifica se é uma nova chave (não indentada ou indentação menor)
        # Consideramos que uma linha sem indentação inicial é nova chave
        if not line.startswith(' ') and not line.startswith('\t'):
            # Se estávamos em multilinha, finaliza o valor anterior
            if in_multiline and current_key:
                result[current_key] = '\n'.join(current_value_lines).strip()
                current_value_lines = []
                in_multiline = False
            # Nova chave
            match = re.match(r'^([a-zA-Z_][a-zA-Z0-9_]*)\s*:\s*(.*)', stripped)
            if match:
                current_key = match.group(1)
                value = match.group(2).strip()
                if value:
                    # Valor na mesma linha
                    if len(value) >= 2 and value[0] == value[-1] and value[0] in ('"', "'"):
                        value = value[1:-1]
                    result[current_key] = value
                    current_key = None
                else:
                    # Valor começa na próxima linha (identado)
                    in_multiline = True
                    current_value_lines = []
        elif in_multiline and current_key:
            # Linha indentada: parte do valor multilinha
            current_value_lines.append(line.rstrip())
    # Finaliza última chave multilinha
    if in_multiline and current_key:
        result[current_key] = '\n'.join(current_value_lines).strip()
    return result


def extract_sdd_state(content):
    """
    Extrai o ÚLTIMO bloco yaml dentro de <!-- @sdd-state -->.
    Retorna dicionário ou None.
    """
    # Encontra todos os matches
    pattern = r'<!--\s*@sdd-state\s*-->\s*```\s*yaml\s*\n(.*?)```'
    matches = list(re.finditer(pattern, content, re.DOTALL))
    if not matches:
        return None
    # Pega o último match
    last_match = matches[-1]
    yaml_text = last_match.group(1)
    if HAS_YAML:
        try:
            return yaml.safe_load(yaml_text)
        except yaml.YAMLError:
            return None
    else:
        return parse_yaml_block(yaml_text)


def extract_frontmatter_meta(content):
    """Extrai o frontmatter YAML (entre ---) e retorna como dicionário."""
    match = re.match(r'^---\s*\n(.*?)\n---', content, re.DOTALL)
    if not match:
        return None
    fm_text = match.group(1)
    if HAS_YAML:
        try:
            return yaml.safe_load(fm_text)
        except yaml.YAMLError:
            return None
    else:
        return parse_yaml_block(fm_text)


def check_internal_links(filepath, content):
    """Verifica se links para arquivos dentro do mesmo diretório existem."""
    base_dir = os.path.dirname(filepath)
    links = re.findall(r'\[([^\]]+)\]\(([^)]+)\)', content)
    for text, link in links:
        if link.startswith('http://') or link.startswith('https://') or link.startswith('#'):
            continue
        full_path = os.path.normpath(os.path.join(base_dir, link))
        if not os.path.isfile(full_path) and not os.path.isdir(full_path):
            if '#' in full_path:
                continue
            warnings.append(f"Link quebrado em {filepath}: '{link}' -> {full_path}")


def validate_skill_file(filepath):
    with open(filepath, 'r', encoding='utf-8') as f:
        content = f.read()

    # 1. Bloco @sdd-state (último)
    state = extract_sdd_state(content)
    if state is None:
        errors.append(f"{filepath}: Bloco @sdd-state ausente ou mal formatado.")
        return

    # 2. Campos obrigatórios
    required_fields = ['version', 'feature_id', 'phase', 'status', 'last_update', 'evidence_checksum']
    for field in required_fields:
        if field not in state or not state[field]:
            errors.append(f"{filepath}: Campo obrigatório '{field}' ausente ou vazio no @sdd-state.")

    # 3. Version consistente
    if state.get('version') != EXPECTED_VERSION:
        warnings.append(f"{filepath}: Versão '{state.get('version')}' diferente da esperada '{EXPECTED_VERSION}'.")

    # 4. evidence_checksum não vazio
    if not state.get('evidence_checksum', '').strip():
        warnings.append(f"{filepath}: evidence_checksum vazio.")

    # 5. Phase válida
    valid_phases = ['DISCOVERY', 'SPECIFY', 'IMPLEMENT', 'VERIFY']
    if state.get('phase') not in valid_phases:
        errors.append(f"{filepath}: Fase '{state.get('phase')}' inválida. Esperada uma de {valid_phases}.")

    # 6. Status válido
    valid_status = ['IN_PROGRESS', 'COMPLETED', 'BLOCKED']
    if state.get('status') not in valid_status:
        errors.append(f"{filepath}: Status '{state.get('status')}' inválido. Esperado {valid_status}.")

    # 7. Parameters (se existir frontmatter com parameters)
    meta = extract_frontmatter_meta(content)
    if meta and 'parameters' in meta:
        params_val = meta['parameters']
        # Se for string, verificar se contém indícios de estrutura
        if isinstance(params_val, str):
            # Verifica se a string contém "type:" e "required:" (mesmo que multilinha)
            if 'type:' in params_val and 'required:' in params_val:
                pass
            else:
                warnings.append(f"{filepath}: 'parameters' pode estar incompleto (sem 'type' ou 'required').")
        elif isinstance(params_val, dict):
            for param_name, param_info in params_val.items():
                if not isinstance(param_info, dict):
                    errors.append(f"{filepath}: Parâmetro '{param_name}' deve ser um dicionário com 'type' e 'required'.")
                else:
                    if 'type' not in param_info:
                        errors.append(f"{filepath}: Parâmetro '{param_name}' sem campo 'type'.")
                    if 'required' not in param_info:
                        warnings.append(f"{filepath}: Parâmetro '{param_name}' sem campo 'required'.")

    # 8. Links internos
    check_internal_links(filepath, content)


def main():
    print("🔍 Validando skills do ecossistema...\n")

    skill_count = 0
    for skill_dir in SKILL_DIRS:
        skill_path = os.path.join(ROOT, skill_dir)
        for sf in find_skill_files(skill_path):
            skill_count += 1
            validate_skill_file(sf)

    print(f"📁 Skills encontradas: {skill_count}")

    if errors:
        print(f"\n❌ ERROS ({len(errors)}):")
        for e in errors:
            print(f"   • {e}")
    else:
        print("✅ Nenhum erro encontrado.")

    if warnings:
        print(f"\n⚠️  AVISOS ({len(warnings)}):")
        for w in warnings:
            print(f"   • {w}")
    else:
        print("✅ Nenhum aviso.")

    return len(errors)


if __name__ == '__main__':
    sys.exit(main())
