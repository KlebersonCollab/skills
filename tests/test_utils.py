import sys
from pathlib import Path
import pytest

# Adiciona a raiz ao PYTHONPATH
sys.path.insert(0, str(Path(__file__).parent.parent))
from scripts.utils import (
    get_all_skills,
    safe_read_file,
    get_skill_metadata,
    safe_copytree,
)


@pytest.fixture
def mock_hub(tmp_path):
    """Cria uma estrutura fake de Skills Hub para testes."""
    # Cria uma skill válida
    skill_a = tmp_path / "skill-a"
    skill_a.mkdir()
    (skill_a / "SKILL.md").write_text(
        "---\nname: skill-a\nversion: 1.0\n---\nConteudo", encoding="utf-8"
    )

    # Cria uma skill invalida (sem SKILL.md)
    skill_b = tmp_path / "skill-b"
    skill_b.mkdir()

    # Cria diretorio ignorado
    node_modules = tmp_path / "node_modules"
    node_modules.mkdir()
    (node_modules / "SKILL.md").write_text("Ignore me", encoding="utf-8")

    return tmp_path


def test_get_all_skills(mock_hub):
    skills = get_all_skills(root=mock_hub)
    assert len(skills) == 1
    assert skills[0].name == "skill-a"


def test_safe_read_file(mock_hub):
    valid_file = mock_hub / "skill-a" / "SKILL.md"
    content = safe_read_file(valid_file)
    assert content is not None
    assert "Conteudo" in content

    missing_file = mock_hub / "missing.md"
    content = safe_read_file(missing_file)
    assert content is None


def test_get_skill_metadata(mock_hub):
    skill_a = mock_hub / "skill-a"
    meta = get_skill_metadata(skill_a)
    assert meta.get("name") == "skill-a"
    assert meta.get("version") == 1.0

    skill_b = mock_hub / "skill-b"
    meta2 = get_skill_metadata(skill_b)
    assert meta2 == {}


def test_safe_copytree(mock_hub):
    src = mock_hub / "skill-a"
    dst = mock_hub / "skill-a-copy"

    result = safe_copytree(src, dst)
    assert result is True
    assert dst.exists()
    assert (dst / "SKILL.md").exists()

    # Testa copia sobre destino existente
    result2 = safe_copytree(src, dst)
    assert result2 is True

    # Testa fonte inexistente
    missing_src = mock_hub / "missing"
    result3 = safe_copytree(missing_src, mock_hub / "missing-dst")
    assert result3 is False
