import re
import os
import sys

def analyze_tasks(file_path):
    """
    Analisa o arquivo tasks.md e separa tarefas concluídas de pendentes.
    """
    if not os.path.exists(file_path):
        return None, None

    with open(file_path, 'r', encoding='utf-8') as f:
        lines = f.readlines()

    completed = []
    pending = []
    
    # Regex para identificar tarefas: "- [ ] Task" ou "- [x] Task"
    task_pattern = re.compile(r"^\s*-\s*\[([ xX])\]\s*(.*)")

    for line in lines:
        match = task_pattern.match(line)
        if match:
            status = match.group(1).lower()
            task_desc = match.group(2).strip()
            
            if status == 'x':
                completed.append(task_desc)
            else:
                pending.append(task_desc)

    return completed, pending

def generate_compact_state(completed, pending, feature_name="Context Compressor"):
    """
    Gera o conteúdo em Markdown para o STATE.md compactado.
    """
    summary = f"## 📜 Histórico Compactado: {feature_name}\n"
    if completed:
        summary += "- **Concluído**: " + ", ".join(completed) + ".\n"
    else:
        summary += "- Nenhuma tarefa concluída no momento.\n"
        
    summary += "\n## 🚧 Tarefas Pendentes\n"
    if pending:
        for task in pending:
            summary += f"- [ ] {task}\n"
    else:
        summary += "- Todas as tarefas foram concluídas!\n"
        
    return summary

def update_state_file(state_path, compact_content):
    """
    Atualiza o arquivo STATE.md com o conteúdo compactado.
    """
    # Mantém o cabeçalho original se existir
    header = "# Operational State\n\n> \"Estado compactado pelo Context Compressor.\"\n\n"
    
    with open(state_path, 'w', encoding='utf-8') as f:
        f.write(header + compact_content)
    
    return True

if __name__ == "__main__":
    import argparse
    from distiller import suggest_learnings
    
    parser = argparse.ArgumentParser(description="Harness Context Compressor")
    parser.add_argument("task_file", help="Caminho do arquivo tasks.md")
    parser.add_argument("state_file", help="Caminho do arquivo STATE.md")
    parser.add_argument("--auto", action="store_true", help="Executa apenas se exceder o threshold")
    parser.add_argument("--threshold", type=int, default=5, help="Número de tasks concluídas para disparar")
    parser.add_argument("--distill", help="Caminho de um validation-report.md para extração de aprendizados")
    
    args = parser.parse_args()
    
    # 1. Destilação de Conhecimento (Fase 3)
    if args.distill:
        print(f"Harness: Iniciando Destilação de Conhecimento de {args.distill}...")
        suggestions = suggest_learnings(args.distill)
        print(suggestions)

    # 2. Compactação de Contexto (Fase 1 e 2)
    t_file = args.task_file
    s_file = args.state_file
    
    comp, pend = analyze_tasks(t_file)
    if comp is not None:
        if args.auto and len(comp) < args.threshold:
            print(f"Harness: {len(comp)} tarefas concluídas. Threshold de {args.threshold} não atingido. Ignorando compactação.")
            sys.exit(0)
            
        state_content = generate_compact_state(comp, pend)
        update_state_file(s_file, state_content)
        print(f"Harness Context Compressor: {len(comp)} tarefas compactadas em {s_file}")
    else:
        print(f"Erro: Arquivo {t_file} não encontrado.")
