import typer
import rich
from rich.console import Console
from rich.table import Table
import questionary
from pathlib import Path
import re
from datetime import datetime

app = typer.Typer(help="SDD CLI - Automatizando o workflow Spec-Driven Development.")
console = Console()

SPECS_DIR = Path(".specs/features")
STATE_FILE = Path(".specs/project/STATE.md")

@app.command()
def init(feature_name: str):
    """Inicializa uma nova feature com a estrutura SDD."""
    feature_dir = SPECS_DIR / feature_name.lower().replace(" ", "-")
    if feature_dir.exists():
        console.print(f"[bold red]Erro:[/bold red] Feature '{feature_name}' já existe.")
        return

    feature_dir.mkdir(parents=True)
    
    # Criar templates
    (feature_dir / "spec.md").write_text(f"# Specification: {feature_name}\n\n## Goal\n\n## Acceptance Criteria (BDD)\n")
    (feature_dir / "plan.md").write_text(f"# Technical Plan: {feature_name}\n\n## Architecture\n\n## Mermaid Diagrams\n")
    (feature_dir / "tasks.md").write_text(f"# Tasks: {feature_name}\n\n- [ ] 1. Define specifications\n- [ ] 2. Technical design\n- [ ] 3. Implementation\n- [ ] 4. Verification\n")
    
    console.print(f"[bold green]Sucesso![/bold green] Feature '{feature_name}' inicializada em {feature_dir}")

@app.command()
def state():
    """Mostra o estado atual das features em desenvolvimento."""
    if not SPECS_DIR.exists():
        console.print("[yellow]Nenhuma feature encontrada.[/yellow]")
        return

    table = Table(title="SDD Project State")
    table.add_column("Feature", style="cyan")
    table.add_column("Progresso", style="magenta")
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
                "Em andamento" if percent < 100 else "Concluída"
            )

    console.print(table)

@app.command()
def task(feature: str, task_id: int, done: bool = True):
    """Marca uma tarefa como concluída (ou pendente)."""
    feature_dir = SPECS_DIR / feature.lower()
    task_file = feature_dir / "tasks.md"
    
    if not task_file.exists():
        console.print(f"[bold red]Erro:[/bold red] Arquivo de tarefas não encontrado para '{feature}'.")
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
        console.print(f"[bold green]Tarefa {task_id} atualizada![/bold green]")
        # Auto-sync STATE.md
        _sync_global_state()
    else:
        console.print(f"[bold yellow]Aviso:[/bold yellow] Tarefa '{task_id}' não encontrada (use o número da lista).")

def _sync_global_state():
    """Sincroniza o STATE.md global com base nas features ativas."""
    if not STATE_FILE.exists():
        return
        
    # Lógica simples de atualização de data e status
    now = datetime.now().strftime("%d de %B de %Y")
    console.print(f"[dim]Sincronizando {STATE_FILE}...[/dim]")
    # Aqui poderíamos ler todas as tasks e gerar o resumo automaticamente

@app.command()
def sync():
    """Sincroniza os mandatos globais (.specs/codebase/GLOBAL_MANDATES.md) com todos os agentes."""
    console.print("[bold cyan]Sincronizando Mandatos Globais...[/bold cyan]")
    import sys
    sys.path.append(str(Path(__file__).parent.parent))
    try:
        from sync_mandates import sync_mandates
        sync_mandates()
        console.print("[bold green]Sucesso![/bold green] Arquivos dos agentes atualizados.")
    except ImportError:
        console.print("[bold red]Erro:[/bold red] Script sync_mandates.py não encontrado.")

@app.command()
def graph():
    """Gera automaticamente o Knowledge Map (Mermaid) do projeto."""
    console.print("[bold cyan]Gerando Knowledge Map...[/bold cyan]")
    # Importar e rodar a lógica do script de geração
    import sys
    sys.path.append(str(Path(__file__).parent.parent))
    try:
        from generate_knowledge_map import generate_knowledge_map
        generate_knowledge_map()
        console.print("[bold green]Sucesso![/bold green] Arquivo KNOWLEDGE-MAP.mermaid atualizado.")
    except ImportError:
        console.print("[bold red]Erro:[/bold red] Script generate_knowledge_map.py não encontrado.")

if __name__ == "__main__":
    app()
