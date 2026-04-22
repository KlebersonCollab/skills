import sys
from pathlib import Path
import pytest
from unittest.mock import patch, MagicMock
import subprocess

# Adiciona a raiz ao PYTHONPATH
sys.path.insert(0, str(Path(__file__).parent.parent))
from scripts.installer import install_local_skill, install_remote_skill, install_skill

@pytest.fixture
def mock_hub(tmp_path):
    """Cria uma estrutura fake de Skills Hub para testes."""
    skill_a = tmp_path / "skill-a"
    skill_a.mkdir()
    (skill_a / "SKILL.md").write_text("---\nname: skill-a\nversion: 1.0\n---\nConteudo", encoding="utf-8")
    return tmp_path

@patch("scripts.installer.get_all_skills")
def test_install_local_skill_success(mock_get_all_skills, mock_hub, tmp_path):
    # Setup mock
    skill_a = mock_hub / "skill-a"
    mock_get_all_skills.return_value = [skill_a]
    
    target_dir = tmp_path / "target"
    target_dir.mkdir()
    
    install_local_skill("skill-a", target=target_dir)
    assert (target_dir / "skill-a" / "SKILL.md").exists()

@patch("scripts.installer.sys.exit")
@patch("scripts.installer.get_all_skills")
def test_install_local_skill_not_found(mock_get_all_skills, mock_exit, mock_hub, tmp_path):
    mock_exit.side_effect = SystemExit(1)
    mock_get_all_skills.return_value = []
    with pytest.raises(SystemExit):
        install_local_skill("non-existent", target=tmp_path)
    mock_exit.assert_called_once_with(1)

@patch("scripts.installer.sys.exit")
@patch("scripts.installer.get_all_skills")
def test_install_local_skill_exists_no_force(mock_get_all_skills, mock_exit, mock_hub, tmp_path):
    mock_exit.side_effect = SystemExit(1)
    skill_a = mock_hub / "skill-a"
    mock_get_all_skills.return_value = [skill_a]
    
    target_dir = tmp_path / "target"
    target_dir.mkdir()
    
    # Primeira instalacao
    install_local_skill("skill-a", target=target_dir)
    
    # Segunda instalacao sem force
    with pytest.raises(SystemExit):
        install_local_skill("skill-a", target=target_dir)
    mock_exit.assert_called_once_with(1)

@patch("scripts.installer.sys.exit")
@patch("scripts.installer.subprocess.run")
def test_install_remote_skill_clone_error(mock_subprocess_run, mock_exit, tmp_path):
    mock_exit.side_effect = SystemExit(1)
    # Simula erro no git clone
    mock_res = MagicMock()
    mock_res.returncode = 1
    mock_res.stderr = "git error"
    mock_subprocess_run.return_value = mock_res
    
    with pytest.raises(SystemExit):
        install_remote_skill("skill-remote", target=tmp_path)
    mock_exit.assert_called_once_with(1)

@patch("scripts.installer.subprocess.run")
@patch("scripts.installer.safe_copytree")
def test_install_remote_skill_success(mock_safe_copytree, mock_subprocess_run, tmp_path, monkeypatch):
    # Simula clone com sucesso
    mock_res = MagicMock()
    mock_res.returncode = 0
    mock_subprocess_run.return_value = mock_res
    
    # Prepara mock_copytree
    mock_safe_copytree.return_value = True
    
    # Em vez de mockar Path.exists de forma global (que quebra coisas), vamos apenas criar o diretório e o SKILL.md no tmpdir que ele vai criar internamente.
    # Como não sabemos o tmpdir, vamos mockar tempfile.TemporaryDirectory para retornar algo previsivel
    import tempfile
    
    test_tmp = tmp_path / "fake-tmp"
    test_tmp.mkdir()
    fake_skill_dir = test_tmp / "skill-remote"
    fake_skill_dir.mkdir()
    (fake_skill_dir / "SKILL.md").write_text("---")
    
    # Mockando TemporaryDirectory context manager
    class MockTmpDir:
        def __enter__(self):
            return str(test_tmp)
        def __exit__(self, exc, value, tb):
            pass
            
    with patch("scripts.installer.tempfile.TemporaryDirectory", return_value=MockTmpDir()):
        install_remote_skill("skill-remote", target=tmp_path)
        
    assert mock_safe_copytree.called

@patch("scripts.installer.install_remote_skill")
@patch("scripts.installer.install_local_skill")
def test_install_skill_router(mock_local, mock_remote, tmp_path):
    install_skill("test", tmp_path, remote=True)
    mock_remote.assert_called_once_with("test", tmp_path, False)
    mock_local.assert_not_called()
    
    install_skill("test", tmp_path, remote=False)
    mock_local.assert_called_once_with("test", tmp_path, False)
