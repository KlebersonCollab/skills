import typer
from rich.console import Console
from rich.table import Table
from pathlib import Path
import re

app = typer.Typer(help="SDD CLI - Automating the Spec-Driven Development workflow.")
console = Console()

SPECS_DIR = Path(".specs/features")
STATE_FILE = Path(".specs/project/STATE.md")


@app.command()
def init(feature_name: str):
    """Initializes a new feature with SDD structure."""
    feature_dir = SPECS_DIR / feature_name.lower().replace(" ", "-")
    if feature_dir.exists():
        console.print(f"[bold red]Error:[/bold red] Feature '{feature_name}' already exists.")
        return

    feature_dir.mkdir(parents=True)

    # Create templates
    (feature_dir / "spec.md").write_text(
        f"# Specification: {feature_name}\n\n## Goal\n\n## Acceptance Criteria (BDD)\n"
    )
    (feature_dir / "plan.md").write_text(
        f"# Technical Plan: {feature_name}\n\n## Architecture\n\n## Mermaid Diagrams\n"
    )
    (feature_dir / "tasks.md").write_text(
        f"# Tasks: {feature_name}\n\n- [ ] 1. Define specifications\n- [ ] 2. Technical design\n- [ ] 3. Implementation\n- [ ] 4. Verification\n"
    )

    console.print(
        f"[bold green]Success![/bold green] Feature '{feature_name}' initialized at {feature_dir}"
    )


@app.command()
def state():
    """Shows the current state of features in development."""
    if not SPECS_DIR.exists():
        console.print("[yellow]No features found.[/yellow]")
        return

    table = Table(title="SDD Project State")
    table.add_column("Feature", style="cyan")
    table.add_column("Progress", style="magenta")
    table.add_column("Status", style="green")

    for feature in SPECS_DIR.iterdir():
        if feature.is_dir() and (feature / "tasks.md").exists():
            content = (feature / "tasks.md").read_text()
            tasks = re.findall(r"- \[(.)\]", content)
            done = tasks.count("x")
            total = len(tasks)
            percent = (done / total * 100) if total > 0 else 0

            table.add_row(
                feature.name,
                f"{done}/{total} ({percent:.0f}%)",
                "In progress" if percent < 100 else "Completed",
            )

    console.print(table)


@app.command()
def task(feature: str, task_id: int, done: bool = True):
    """Marks a task as completed (or pending)."""
    feature_dir = SPECS_DIR / feature.lower()
    task_file = feature_dir / "tasks.md"

    if not task_file.exists():
        console.print(
            f"[bold red]Error:[/bold red] Tasks file not found for '{feature}'."
        )
        return

    content = task_file.read_text().splitlines()
    new_content = []
    found = False

    task_pattern = rf"- \[(.)\] {task_id}\."

    for line in content:
        if re.search(task_pattern, line):
            mark = "x" if done else " "
            line = re.sub(r"\[.\]", f"[{mark}]", line)
            found = True
        new_content.append(line)

    if found:
        task_file.write_text("\n".join(new_content) + "\n")
        console.print(f"[bold green]Task {task_id} updated![/bold green]")
        # Auto-sync STATE.md
        _sync_global_state()
    else:
        console.print(
            f"[bold yellow]Warning:[/bold yellow] Task '{task_id}' not found (use the list number)."
        )


def _sync_global_state():
    """Synchronizes the global STATE.md based on active features."""
    if not STATE_FILE.exists():
        return

    # Simple logic for date and status updates
    console.print(f"[dim]Synchronizing {STATE_FILE}...[/dim]")
    # Here we could read all tasks and generate the summary automatically


@app.command()
def sync():
    """Synchronizes global mandates (.specs/codebase/GLOBAL_MANDATES.md) with all agents."""
    console.print("[bold cyan]Synchronizing Global Mandates...[/bold cyan]")
    import sys

    sys.path.append(str(Path(__file__).parent.parent))
    try:
        from sync_mandates import sync_mandates

        sync_mandates()
        console.print(
            "[bold green]Success![/bold green] Agent files updated."
        )
    except ImportError:
        console.print(
            "[bold red]Error:[/bold red] sync_mandates.py script not found."
        )


@app.command()
def graph():
    """Automatically generates the project's Knowledge Map (Mermaid)."""
    console.print("[bold cyan]Generating Knowledge Map...[/bold cyan]")
    # Import and run generation logic
    import sys

    sys.path.append(str(Path(__file__).parent.parent))
    try:
        from generate_knowledge_map import generate_knowledge_map

        generate_knowledge_map()
        console.print(
            "[bold green]Success![/bold green] KNOWLEDGE-MAP.mermaid file updated."
        )
    except ImportError:
        console.print(
            "[bold red]Error:[/bold red] generate_knowledge_map.py script not found."
        )


if __name__ == "__main__":
    app()
