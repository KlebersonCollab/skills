#!/usr/bin/env python3
"""
Token Distiller - Context Micro-Compaction Tool (v2.4.0)
Practical implementation of the advanced Token Distiller skill,
compressing text with three intensity levels and Auto-Clarity triggers.
"""

import sys
import re
import argparse

def check_auto_clarity(text: str) -> bool:
    """
    Checks if the input text triggers the Auto-Clarity safety valve.
    Returns True if a security warning, irreversible action, or multi-step
    instruction sequence is detected, indicating compression should be bypassed.
    """
    security_keywords = [
        "security warning", "vulnerability", "unauthorized access", "cve-", 
        "exploit", "secret", "private key", "auth bypass", "sql injection", "unsafe"
    ]
    irreversible_keywords = [
        "permanent", "cannot be undone", "drop table", "delete all", 
        "purge database", "destructive", "rm -rf", "irreversible"
    ]
    # Check for numbered lists or step patterns
    multi_step_patterns = [
        r"\b(step|phase)\s+\d+",
        r"^\s*\d+\.\s+",
        r"^\s*-\s+Step\s+\d+"
    ]
    
    text_lower = text.lower()
    for kw in security_keywords + irreversible_keywords:
        if kw in text_lower:
            return True
            
    for pat in multi_step_patterns:
        if re.search(pat, text, re.MULTILINE | re.IGNORECASE):
            return True
            
    return False

def abbreviate_text(text: str) -> str:
    """
    Abbreviates common prose words and strips conjunctions for ultra mode,
    while carefully keeping backticked code sections completely unchanged.
    """
    parts = text.split("`")
    for i in range(len(parts)):
        if i % 2 == 0:  # Outside backticks
            # Causality arrows
            parts[i] = re.sub(r"\s+\b(leads to|causes|results in|leads\s+to|results\s+in)\b\s+", " → ", parts[i], flags=re.IGNORECASE)
            
            # Common prose abbreviations
            parts[i] = re.sub(r"\bdatabase\b", "DB", parts[i], flags=re.IGNORECASE)
            parts[i] = re.sub(r"\bdatabases\b", "DBs", parts[i], flags=re.IGNORECASE)
            parts[i] = re.sub(r"\bauthentication\b", "auth", parts[i], flags=re.IGNORECASE)
            parts[i] = re.sub(r"\bconfiguration\b", "config", parts[i], flags=re.IGNORECASE)
            parts[i] = re.sub(r"\brequest\b", "req", parts[i], flags=re.IGNORECASE)
            parts[i] = re.sub(r"\brequests\b", "reqs", parts[i], flags=re.IGNORECASE)
            parts[i] = re.sub(r"\bresponse\b", "res", parts[i], flags=re.IGNORECASE)
            parts[i] = re.sub(r"\bresponses\b", "res", parts[i], flags=re.IGNORECASE)
            parts[i] = re.sub(r"\bfunction\b", "fn", parts[i], flags=re.IGNORECASE)
            parts[i] = re.sub(r"\bfunctions\b", "fns", parts[i], flags=re.IGNORECASE)
            parts[i] = re.sub(r"\bimplementation\b", "impl", parts[i], flags=re.IGNORECASE)
            parts[i] = re.sub(r"\bconnection\b", "conn", parts[i], flags=re.IGNORECASE)
            parts[i] = re.sub(r"\bconnections\b", "conns", parts[i], flags=re.IGNORECASE)
            
            # Strip conjunctions
            parts[i] = re.sub(r"\b(and|or|but|because|since|although)\b\s*", "", parts[i], flags=re.IGNORECASE)
    return "`".join(parts)

def apply_pattern_heuristics(text: str) -> str:
    """
    Tries to map explanation text into the strict response pattern:
    [thing] [action] [reason]. [next step].
    """
    # Pattern 1: X has Y because of Z. We need to W.
    match = re.search(
        r"(?:we have a|there is a)?\s*(.*?)\s+(?:has|uses|is)\s+(.*?)\s+because\s+(.*?)\.\s+(?:we need to|please|fix:)\s*(.*)",
        text,
        re.IGNORECASE
    )
    if match:
        thing, action, reason, next_step = match.groups()
        return f"[{thing.strip()}] [{action.strip()}] [{reason.strip()}]. [{next_step.strip()}]."
    
    return text

def to_caveman(text: str, level: str = "full") -> str:
    """
    Main distillation engine. Applies rules according to the selected intensity level.
    """
    if check_auto_clarity(text):
        return f"[AUTO-CLARITY TRIGGERED: Safety Valve Active - Bypassing Compression]\n{text}"

    # Pre-compaction fillers for all levels
    fillers = r"\b(really|basically|just|actually|simply|very|of_course|absolutely|sure|certainly|obviously|clearly)\b\s*"
    text = re.sub(fillers, "", text, flags=re.IGNORECASE)

    if level == "lite":
        # lite level keeps articles + full sentences, just removes filler/hedging
        return text.strip()

    # levels 'full' and 'ultra' drop articles
    articles = r"\b(the|a|an)\b\s*"
    text = re.sub(articles, "", text, flags=re.IGNORECASE)

    # Use short direct synonyms
    text = re.sub(r"\bimplement a solution for\b", "fix", text, flags=re.IGNORECASE)
    text = re.sub(r"\bimplement\b", "do", text, flags=re.IGNORECASE)
    text = re.sub(r"\bextensive\b", "big", text, flags=re.IGNORECASE)

    # Apply strict pattern heuristics where possible
    text = apply_pattern_heuristics(text)

    if level == "ultra":
        text = abbreviate_text(text)

    return text.strip()

def main():
    parser = argparse.ArgumentParser(description="Compacts text using Advanced Caveman Mode rules (v2.4.0).")
    parser.add_argument("--level", choices=["lite", "full", "ultra"], default="full", help="Intensity level of compression (default: full)")
    group = parser.add_mutually_exclusive_group(required=True)
    group.add_argument("--text", help="Direct string input to compress")
    group.add_argument("--file", help="File path containing text to compress")
    
    args = parser.parse_args()
    
    content = ""
    if args.text:
        content = args.text
    elif args.file:
        try:
            with open(args.file, 'r', encoding='utf-8') as f:
                content = f.read()
        except Exception as e:
            print(f"Error reading file {args.file}: {e}", file=sys.stderr)
            sys.exit(1)
            
    print(f"--- 🪨 CAVEMAN DISTILLATION (Level: {args.level.upper()}) ---")
    distilled = to_caveman(content, args.level)
    print(distilled)
    
    # Savings calculation
    orig_len = len(content)
    new_len = len(distilled)
    savings = 100 - ((new_len / orig_len) * 100) if orig_len > 0 else 0
    print(f"\n--- [Token Savings: {savings:.1f}%] ---")

if __name__ == "__main__":
    main()
