import unittest
import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))
from scripts.compressor_utility import MarkdownCompressor


class TestCompressor(unittest.TestCase):
    def setUp(self):
        self.compressor = MarkdownCompressor()

    def test_preserve_code_blocks(self):
        text = "Aqui está um código:\n```python\nprint('hello')\n```\nMuito legal."
        compressed = self.compressor.compress_text(text)
        self.assertIn("```python\nprint('hello')\n```", compressed)

    def test_remove_fillers_and_articles(self):
        text = "Eu realmente acho que o problema é simplesmente a configuração do servidor."
        compressed = self.compressor.compress_text(text)
        # Expected: "Eu acho problema configuração servidor." (approx)
        self.assertNotIn("realmente", compressed)
        self.assertNotIn("simplesmente", compressed)
        self.assertNotIn(" o ", compressed)

    def test_preserve_headings(self):
        text = "# Meu Título\nProsa explicativa aqui."
        compressed = self.compressor.compress_text(text)
        self.assertIn("# Meu Título", compressed)

    def test_preserve_links(self):
        text = "Veja o [link](https://google.com) para mais info."
        compressed = self.compressor.compress_text(text)
        self.assertIn("[link](https://google.com)", compressed)


if __name__ == "__main__":
    unittest.main()
