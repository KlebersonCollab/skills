import re
import os
import sys


def analyze_tasks(file_path):
    """
    Analyzes the tasks.md file and separates completed tasks from pending ones.
    """
    if not os.path.exists(file_path):
        return None, None

    with open(file_path, "r", encoding="utf-8") as f:
        lines = f.readlines()

    completed = []
    pending = []

    # Regex to identify tasks: "- [ ] Task" or "- [x] Task"
    task_pattern = re.compile(r"^\s*-\s*\[([ xX])\]\s*(.*)")

    for line in lines:
        match = task_pattern.match(line)
        if match:
            status = match.group(1).lower()
            task_desc = match.group(2).strip()

            if status == "x":
                completed.append(task_desc)
            else:
                pending.append(task_desc)

    return completed, pending


def generate_compact_state(completed, pending, feature_name="Context Compressor"):
    """
    Generates Markdown content for the compressed STATE.md.
    """
    summary = f"## 📜 Compressed History: {feature_name}\n"
    if completed:
        summary += "- **Completed**: " + ", ".join(completed) + ".\n"
    else:
        summary += "- No tasks completed at the moment.\n"

    summary += "\n## 🚧 Pending Tasks\n"
    if pending:
        for task in pending:
            summary += f"- [ ] {task}\n"
    else:
        summary += "- All tasks have been completed!\n"

    return summary


def update_state_file(state_path, compact_content):
    """
    Updates the STATE.md file with the compressed content.
    """
    # Maintains the original header if it exists
    header = '# Operational State\n\n> "State compressed by Context Compressor."\n\n'

    with open(state_path, "w", encoding="utf-8") as f:
        f.write(header + compact_content)

    return True


if __name__ == "__main__":
    import argparse
    from distiller import suggest_learnings

    parser = argparse.ArgumentParser(description="Harness Context Compressor")
    parser.add_argument("task_file", help="Path to tasks.md file")
    parser.add_argument("state_file", help="Path to STATE.md file")
    parser.add_argument(
        "--auto", action="store_true", help="Executes only if threshold is exceeded"
    )
    parser.add_argument(
        "--threshold",
        type=int,
        default=5,
        help="Number of completed tasks to trigger",
    )
    parser.add_argument(
        "--distill",
        help="Path to a validation-report.md for extracting learnings",
    )

    args = parser.parse_args()

    # 1. Knowledge Distillation (Phase 3)
    if args.distill:
        print(f"Harness: Starting Knowledge Distillation from {args.distill}...")
        suggestions = suggest_learnings(args.distill)
        print(suggestions)

    # 2. Context Compression (Phases 1 and 2)
    t_file = args.task_file
    s_file = args.state_file

    comp, pend = analyze_tasks(t_file)
    if comp is not None:
        if args.auto and len(comp) < args.threshold:
            print(
                f"Harness: {len(comp)} tasks completed. Threshold of {args.threshold} not reached. Ignoring compression."
            )
            sys.exit(0)

        state_content = generate_compact_state(comp, pend)
        update_state_file(s_file, state_content)
        print(
            f"Harness Context Compressor: {len(comp)} tasks compressed into {s_file}"
        )
    else:
        print(f"Error: File {t_file} not found.")
