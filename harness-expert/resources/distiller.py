import re
import os
import sys


def extract_patterns(report_path):
    """
    Extracts [ERROR] and [FIXED] patterns from a report file.
    """
    if not os.path.exists(report_path):
        return []

    with open(report_path, "r", encoding="utf-8") as f:
        content = f.read()

    # Search for blocks like:
    # [ERROR]: Description
    # [FIXED]: Solution
    pattern = re.compile(
        r"\[ERROR\]:\s*(.*?)\n\s*\[FIXED\]:\s*(.*?)(?=\n\n|\n\[|$)", re.DOTALL
    )
    matches = pattern.findall(content)

    learnings = []
    for error, fix in matches:
        learnings.append({"error": error.strip(), "fix": fix.strip()})

    return learnings


def format_learning_entry(learning):
    """
    Formats the learning for the LEARNINGS.md standard.
    """
    title = learning["error"].split("\n")[0][
        :50
    ]  # Uses the first line of the error as the title
    entry = f"## [PATTERN] {title}\n"
    entry += f" - **Context**: {learning['error']}\n"
    entry += f" - **Solution**: {learning['fix']}\n\n"
    return entry


def suggest_learnings(report_file):
    """
    Main learning suggestion function.
    """
    patterns = extract_patterns(report_file)
    if not patterns:
        return "No new patterns identified."

    output = "# Learning Suggestions (Auto-Distilled)\n\n"
    for p in patterns:
        output += format_learning_entry(p)

    return output


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python3 distiller.py <report_file>")
        sys.exit(1)

    report = sys.argv[1]
    suggestions = suggest_learnings(report)
    print(suggestions)
