import sys
import os
import argparse
from pathlib import Path

# Add project root to sys.path
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))
from scripts.compressor_utility import MarkdownCompressor

def main():
    parser = argparse.ArgumentParser(description="Hub CLI - AI Agent Skills Hub Utility")
    subparsers = parser.add_subparsers(dest="command", help="Available commands")

    # Compress command
    compress_parser = subparsers.add_parser("compress", help="Compress a markdown file to save tokens")
    compress_parser.add_argument("file", help="Path to the markdown file to compress")

    # Mode command
    mode_parser = subparsers.add_parser("mode", help="Switch between operational modes (low/high)")
    mode_parser.add_argument("value", choices=["low", "high", "caveman", "premium"], help="The mode to switch to")

    # Check state command
    subparsers.add_parser("check-state", help="Check if STATE.md needs compression")

    args = parser.parse_args()

    if args.command == "compress":
        file_path = Path(args.file)
        if not file_path.exists():
            print(f"Error: File {args.file} not found.")
            sys.exit(1)
        
        compressor = MarkdownCompressor()
        compressor.compress_file(file_path)
    elif args.command == "mode":
        mode_file = Path(".hub-mode")
        value = "premium" if args.value in ["high", "premium"] else "caveman"
        mode_file.write_text(value)
        print(f"Operational mode switched to: {value.upper()}")
    elif args.command == "check-state":
        state_file = Path(".specs/project/STATE.md")
        if state_file.exists() and state_file.stat().st_size > 5000: # Heuristic for ~1000 tokens
            print(f"STATE.md is large ({state_file.stat().st_size} bytes). Compressing...")
            compressor = MarkdownCompressor()
            compressor.compress_file(state_file)
        else:
            print("STATE.md size is optimal.")
    else:
        parser.print_help()

if __name__ == "__main__":
    main()
