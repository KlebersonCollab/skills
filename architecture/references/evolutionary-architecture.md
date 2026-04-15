# Evolutionary Architecture & Fitness Functions

Este guia descreve como projetar sistemas que suportam mudanças guiadas e incrementais através de múltiplas dimensões.

---

## 1. O que é Arquitetura Evolutiva?

É uma arquitetura que suporta mudanças guiadas e incrementais como um princípio de primeira classe. Em vez de tentar prever todo o futuro (YAGNI), ela foca em proteger as características vitais do sistema enquanto ele evolui.

## 2. Fitness Functions (Funções de Aptidão)

Uma **Fitness Function** é qualquer mecanismo que fornece uma avaliação objetiva de uma característica de integridade arquitetural. Elas podem ser automatizadas (testes unitários, scripts) ou manuais (code reviews).

### Exemplos de Dimensões Protegidas:
- **Acoplamento**: Garantir que o módulo A não dependa do módulo B.
- **Segurança**: Verificar se nenhum endpoint está exposto sem autenticação.
- **Performance**: Validar se o tempo de resposta médio está abaixo de 200ms.
- **Qualidade**: Impedir complexidade ciclomática acima de um limite X.

## 3. Tipos de Fitness Functions

1.  **Atomic**: Executada em um único contexto (ex: um teste unitário verificando dependências de pacotes).
2.  **Holistic**: Avalia o sistema como um todo (ex: teste de carga para validar escalabilidade).
3.  **Continuous**: Executada no pipeline de CI/CD em cada commit.
4.  **Temporal**: Acionada por eventos de tempo (ex: verificar expiração de certificados).

## 4. Implementação Prática (Arquitetura como Código)

Utilize ferramentas de **Linter Arquitetural** ou scripts simples no CI:

### Exemplo: Validando Acoplamento (Python/Import Linter)
```toml
# setup.cfg ou import-linter.config
[importlinter]
root_package = meu_projeto

[[importlinter:contracts:forbidden]]
name = Camada de Dominio nao deve depender da Camada de Infra
source_modules = meu_projeto.domain
forbidden_modules = meu_projeto.infra
```

### Exemplo: Validando Ciclos (Dart/ArchTest)
```dart
void main() {
  test('Sem dependências circulares entre features', () {
    final arch = ArchTest.fromDirectory('lib/features');
    expect(arch, hasNoCircularDependencies());
  });
}
```

## 5. Boas Práticas

- **Feedback Rápido**: Prefira funções atômicas e contínuas que rodam em segundos.
- **Fail Fast**: Falhe o build se uma Fitness Function for violada.
- **Evolução do Teste**: À medida que a arquitetura muda, as Fitness Functions também devem ser atualizadas ou removidas.
