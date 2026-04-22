import os
import sys


def check_layer_dependencies():
    """
    Exemplo simplificado de Fitness Function que verifica se
    a camada de Domínio está importando algo da camada de Infra.
    """
    domain_dir = "lib/domain"
    infra_layer_keyword = "infra/"
    violations = []

    for root, _, files in os.walk(domain_dir):
        for file in files:
            if file.endswith((".dart", ".py", ".ts")):
                path = os.path.join(root, file)
                with open(path, "r") as f:
                    for line_no, line in enumerate(f, 1):
                        if infra_layer_keyword in line:
                            violations.append(
                                f"{path}:{line_no} -> Domain layer cannot depend on Infra layer."
                            )

    if violations:
        print("❌ Architectural Fitness Function Failed:")
        for v in violations:
            print(v)
        sys.exit(1)
    else:
        print("✅ Architectural Fitness Function Passed: Layer integrity maintained.")


if __name__ == "__main__":
    check_layer_dependencies()
