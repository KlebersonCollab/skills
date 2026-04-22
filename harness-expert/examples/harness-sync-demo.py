import re
import os


def sync_tasks(task_file, activity_log):
    """
    Simula o Harness sincronizando o tasks.md baseado em logs de atividade.
    """
    if not os.path.exists(task_file):
        print(f"Erro: Arquivo {task_file} não encontrado.")
        return

    with open(task_file, "r") as f:
        content = f.read()

    # Exemplo simples: se o log contém "COMPLETED: <task_name>", marca no tasks.md
    completed_tasks = re.findall(r"COMPLETED: (.*)", activity_log)

    for task in completed_tasks:
        # Busca a linha da task e marca como concluída [x]
        pattern = rf"- \[ \] {re.escape(task)}"
        replacement = f"- [x] {task}"
        content = re.sub(pattern, replacement, content)

    with open(task_file, "w") as f:
        f.write(content)

    print(f"Harness Sync: {len(completed_tasks)} tarefas sincronizadas.")


if __name__ == "__main__":
    # Exemplo de uso
    sample_log = "COMPLETED: Setup Project\nCOMPLETED: Define Architecture"
    # sync_tasks("tasks.md", sample_log)
    print("Script harness-sync carregado. Pronto para integração.")
