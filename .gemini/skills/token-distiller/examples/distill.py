#!/usr/bin/env python3
"""
Token Distiller - Context Micro-Compaction Tool
This script provides a practical implementation of the Token Distiller skill,
compressing large log files or context dumps into dense "Caveman" summaries.
"""

import sys
import re

def to_caveman(text: str) -> str:
    """
    Applies Low Token Mode (Caveman) rules:
    - Removes articles, fillers.
    - Uses fragments.
    """
    # Remove articles
    text = re.sub(r'\b(the|a|an)\b\s+', '', text, flags=re.IGNORECASE)
    # Remove fillers
    text = re.sub(r'\b(really|basically|just|very|actually)\b\s+', '', text, flags=re.IGNORECASE)
    
    # Simple summarization for the sake of the example
    lines = text.split('\n')
    distilled_lines = []
    
    for line in lines:
        line = line.strip()
        if not line:
            continue
            
        if "error" in line.lower() or "exception" in line.lower():
            distilled_lines.append(f"[ERROR DETECTED]: {line}")
        elif "success" in line.lower():
            distilled_lines.append("[SUCCESS]")
            
    # If it's general text, just truncate and join
    if not distilled_lines:
        words = text.split()
        if len(words) > 20:
            return " ".join(words[:20]) + "... [Truncated. Use /mode high for details]"
        return " ".join(words)
        
    return "\n".join(distilled_lines)

def main():
    if len(sys.argv) < 2:
        print("Usage: python3 distill.py <file.log>")
        sys.exit(1)
        
    filename = sys.argv[1]
    
    try:
        with open(filename, 'r') as f:
            content = f.read()
    except Exception as e:
        print(f"Error reading {filename}: {e}")
        sys.exit(1)
        
    print("--- 🪨 CAVEMAN DISTILLATION ---")
    distilled = to_caveman(content)
    print(distilled)
    
    # Calculate savings
    original_size = len(content)
    new_size = len(distilled)
    savings = 100 - ((new_size / original_size) * 100) if original_size > 0 else 0
    
    print(f"\n--- [Token Savings: {savings:.1f}%] ---")

if __name__ == "__main__":
    main()
