import re
import os
import sys


def extract_patterns(report_path):
    """
    Extrai padrões de [ERROR] e [FIXED] de um arquivo de relatório.
    """
    if not os.path.exists(report_path):
        return []

    with open(report_path, "r", encoding="utf-8") as f:
        content = f.read()

    # Busca por blocos do tipo:
    # [ERROR]: Descrição
    # [FIXED]: Solução
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
    Formata o aprendizado para o padrão do LEARNINGS.md.
    """
    title = learning["error"].split("\n")[0][
        :50
    ]  # Usa a primeira linha do erro como título
    entry = f"## [PATTERN] {title}\n"
    entry += f" - **Contexto**: {learning['error']}\n"
    entry += f" - **Solução**: {learning['fix']}\n\n"
    return entry


def suggest_learnings(report_file):
    """
    Função principal de sugestão de aprendizados.
    """
    patterns = extract_patterns(report_file)
    if not patterns:
        return "Nenhum novo padrão identificado."

    output = "# Sugestões de Aprendizados (Auto-Distilled)\n\n"
    for p in patterns:
        output += format_learning_entry(p)

    return output


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Uso: python3 distiller.py <report_file>")
        sys.exit(1)

    report = sys.argv[1]
    suggestions = suggest_learnings(report)
    print(suggestions)
