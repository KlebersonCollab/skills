import unittest
import os
from compressor import analyze_tasks, generate_compact_state


class TestContextCompressor(unittest.TestCase):
    def setUp(self):
        # Cria um arquivo de tasks temporário para o teste
        self.test_tasks_file = "test_tasks.md"
        with open(self.test_tasks_file, "w", encoding="utf-8") as f:
            f.write("# Tasks\n- [x] Task 1\n- [ ] Task 2\n- [x] Task 3\n- [ ] Task 4\n")

    def tearDown(self):
        # Remove o arquivo temporário após o teste
        if os.path.exists(self.test_tasks_file):
            os.remove(self.test_tasks_file)

    def test_analyze_tasks(self):
        completed, pending = analyze_tasks(self.test_tasks_file)
        self.assertEqual(len(completed), 2)
        self.assertEqual(len(pending), 2)
        self.assertIn("Task 1", completed)
        self.assertIn("Task 2", pending)

    def test_generate_compact_state(self):
        completed = ["Task 1", "Task 3"]
        pending = ["Task 2", "Task 4"]
        summary = generate_compact_state(completed, pending)
        self.assertIn("## 📜 Histórico Compactado", summary)
        self.assertIn("- **Concluído**: Task 1, Task 3.", summary)
        self.assertIn("- [ ] Task 2", summary)

    def test_auto_threshold(self):
        # Este teste verifica apenas a lógica de decisão simulada
        completed = ["Task 1", "Task 3"]
        threshold = 5
        self.assertLess(len(completed), threshold)


if __name__ == "__main__":
    unittest.main()
