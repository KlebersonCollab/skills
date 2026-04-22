import pytest
from pathlib import Path
from unittest.mock import patch
from scripts.generate_knowledge_map import generate_knowledge_map


@pytest.fixture
def mock_skills():
    # Retorna objetos Path mockados (com nomes de skill)
    return [
        Path("skill-alpha"),
        Path("skill-beta"),
    ]


@pytest.fixture
def mock_metadata():
    # Mapeamento do nome da skill para seu metadado
    return {
        "skill-alpha": {
            "name": "Alpha",
            "version": "1.0.0",
            "category": "architecture",
            "uses": ["skill-beta"],
            "references": ["skill-gama"],
        },
        "skill-beta": {
            "name": "Beta",
            "version": "1.0.0",
            "category": "development",
            "uses": [],
            "references": [],
        },
    }


@patch("scripts.generate_knowledge_map.get_all_skills")
@patch("scripts.generate_knowledge_map.get_skill_metadata")
@patch("pathlib.Path.read_text")
@patch("pathlib.Path.write_text")
def test_generate_knowledge_map(
    mock_write, mock_read, mock_get_meta, mock_get_all, mock_skills, mock_metadata
):
    mock_get_all.return_value = mock_skills

    # Simular metadados baseado na skill requisitada
    def get_meta_side_effect(skill_path):
        return mock_metadata.get(skill_path.name, {})

    mock_get_meta.side_effect = get_meta_side_effect

    # Simular o texto do MD contendo ou não menções
    def read_side_effect(*args, **kwargs):
        return "Fake SKILL.md com 🔒 Prerequisites (Mandatory) e menção a skill-beta"

    mock_read.side_effect = read_side_effect

    # Executar
    generate_knowledge_map()

    # Asserts
    mock_write.assert_called_once()
    written_content = mock_write.call_args[0][0]

    # Verificar se as tags de Mermaid estão no documento final
    assert "graph TD" in written_content
    # Verifica o node alpha
    assert 'skill_alpha["Alpha (1.0.0) 🛡️"]' in written_content
    # Verifica as relações explícitas
    assert "skill_alpha --(uses)--> skill_beta" in written_content
    assert "skill_alpha --(references)--> skill_gama" in written_content
