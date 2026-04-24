import shutil
import subprocess
from pathlib import Path


def test_sync_functionality():
    root_dir = Path(".")
    test_skill_name = "test-sync-skill"
    test_skill_dir = root_dir / test_skill_name

    print("🚀 Starting integration test for synchronization script...")

    try:
        # 1. Create a temporary test skill
        test_skill_dir.mkdir(exist_ok=True)
        (test_skill_dir / "SKILL.md").write_text(
            "---\nname: test-sync-skill\nversion: 0.1.0\n---"
        )
        print(f"   - Test skill created: {test_skill_name}")

        # 2. Run the sync script
        print("   - Executing scripts/sync_mandates.py...")
        subprocess.run(
            ["python3", "scripts/sync_mandates.py"], check=True, capture_output=True
        )

        # 3. Verify if the skill was propagated to agents
        agents = [".agent", ".gemini", ".claude"]
        for agent in agents:
            dst = root_dir / agent / "skills" / test_skill_name
            if not dst.exists():
                raise AssertionError(
                    f"❌ Failure: Skill not found in {agent}/skills/"
                )
            if not (dst / "SKILL.md").exists():
                raise AssertionError(f"❌ Failure: SKILL.md missing in {agent}/skills/")
            print(f"   ✅ Verified in {agent}/skills/")

        print("🎉 Synchronization test completed with SUCCESS!")

    finally:
        # 4. Cleanup
        if test_skill_dir.exists():
            shutil.rmtree(test_skill_dir)
        for agent in [".agent", ".gemini", ".claude"]:
            dst = root_dir / agent / "skills" / test_skill_name
            if dst.exists():
                shutil.rmtree(dst)
        print("   - Cleanup completed.")


if __name__ == "__main__":
    test_sync_functionality()
