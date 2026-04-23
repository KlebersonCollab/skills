#!/usr/bin/env python3
import re
import sys
from pathlib import Path

class MarkdownCompressor:
    """
    Minifies Markdown prose while preserving technical structure.
    Inspired by Caveman-speak.
    """

    # Elements to preserve exactly
    PRESERVE_PATTERNS = [
        r'```[\s\S]*?```',  # Code blocks
        r'`[^`]*?`',         # Inline code
        r'\[.*?\]\(.*?\)',   # Links
        r'^#+ .*$',          # Headings
        r'\|.*?\|',          # Tables
        r'https?://\S+',     # Raw URLs
        r'/[a-zA-Z0-9._/-]+' # File paths (heuristic)
    ]

    # Fillers to remove
    FILLER_WORDS = [
        r'\bjust\b', r'\breally\b', r'\bbasically\b', r'\bactually\b',
        r'\bsimply\b', r'\bvery\b', r'\bof course\b', r'\bhappy to\b',
        r'\bsure\b', r'\bcertainly\b', r'\bplease\b', r'\bkindly\b',
        r'\brealmente\b', r'\bsimplesmente\b', r'\bbasicamente\b', r'\brealmente\b',
        r'\bapenas\b', r'\bsó\b', r'\bcom certeza\b', r'\bcom certeza\b'
    ]

    # Articles to remove (in Portuguese and English)
    ARTICLES = [
        r'\bthe\b', r'\ba\b', r'\ban\b',
        r'\bo\b', r'\ba\b', r'\bos\b', r'\bas\b', r'\bum\b', r'\buma\b',
        r'\buns\b', r'\bumas\b', r'\bde\b', r'\bdo\b', r'\bda\b', r'\bdos\b', r'\bdas\b'
    ]

    def __init__(self):
        self.placeholder_map = {}
        self.counter = 0

    def _replace_with_placeholder(self, match):
        placeholder = f"__PRESERVE_{self.counter}__"
        self.placeholder_map[placeholder] = match.group(0)
        self.counter += 1
        return placeholder

    def compress_text(self, text: str) -> str:
        # 1. Protect technical elements
        pattern = '|'.join(f'({p})' for p in self.PRESERVE_PATTERNS)
        text = re.sub(pattern, self._replace_with_placeholder, text, flags=re.MULTILINE)

        # 2. Linguistic compression
        # Remove filler words (case insensitive)
        for filler in self.FILLER_WORDS:
            text = re.sub(filler, '', text, flags=re.IGNORECASE)

        # Remove articles
        for article in self.ARTICLES:
            text = re.sub(article, '', text, flags=re.IGNORECASE)

        # 3. Clean up whitespace
        text = re.sub(r' +', ' ', text)
        text = re.sub(r'\n{3,}', '\n\n', text)
        text = text.strip()

        # 4. Restore technical elements
        for placeholder, original in self.placeholder_map.items():
            text = text.replace(placeholder, original)

        return text

    def compress_file(self, input_path: Path):
        if not input_path.exists():
            print(f"Error: {input_path} does not exist.")
            return

        original_content = input_path.read_text()
        compressed_content = self.compress_text(original_content)

        # Save backup
        backup_path = input_path.with_suffix('.original.md')
        if not backup_path.exists():
            backup_path.write_text(original_content)
            print(f"Backup created: {backup_path}")

        # Overwrite original
        input_path.write_text(compressed_content)
        
        reduction = 100 - (len(compressed_content) / len(original_content) * 100)
        print(f"Compressed {input_path.name}: {len(original_content)} -> {len(compressed_content)} bytes ({reduction:.1f}% saved)")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python compressor.py <file>")
        sys.exit(1)

    file_to_compress = Path(sys.argv[1])
    compressor = MarkdownCompressor()
    compressor.compress_file(file_to_compress)
