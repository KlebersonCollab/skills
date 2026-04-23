#!/usr/bin/env python3
import sys
import argparse
from pathlib import Path
from scripts.compressor_utility import MarkdownCompressor

def main():
    parser = argparse.ArgumentParser(description="Hub CLI - AI Agent Skills Hub Utility")
    subparsers = parser.add_subparsers(dest="command", help="Available commands")

    # Compress command
    compress_parser = subparsers.add_parser("compress", help="Compress a markdown file to save tokens")
    compress_parser.add_argument("file", help="Path to the markdown file to compress")

    args = parser.parse_args()

    if args.command == "compress":
        file_path = Path(args.file)
        if not file_path.exists():
            print(f"Error: File {args.file} not found.")
            sys.exit(1)
        
        compressor = MarkdownCompressor()
        compressor.compress_file(file_path)
    else:
        parser.print_help()

if __name__ == "__main__":
    main()
