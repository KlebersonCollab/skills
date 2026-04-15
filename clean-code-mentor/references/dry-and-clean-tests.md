# DRY & Clean Tests: Quality Assurance

Princípios para garantir a coesão do código base e a confiabilidade da suíte de testes.

---

## 1. DRY (Don't Repeat Yourself)
**"Toda peça de conhecimento deve ter uma representação única, não ambígua e autoritativa dentro do sistema."**

- **O que NÃO é DRY**: Duplicação de lógica (ex: calcular o mesmo imposto em dois lugares diferentes).
- **O que É DRY**: Centralizar o cálculo em uma única função ou serviço que é reutilizado por todos.
- **Sinais de Violação**:
    - Mudanças em um lugar exigem mudanças em múltiplos outros lugares ("Shotgun Surgery").
    - Padrões de código idênticos em classes não relacionadas.
- **Estratégia**: Abstraia a lógica duplicada para utilitários, mixins ou serviços especializados.
- **Atenção**: Não confunda duplicação de *código* (ex: dois `for` parecidos) com duplicação de *conhecimento* (ex: a mesma regra de negócio em dois lugares). Forçar o DRY em lógicas que evoluem de forma diferente pode levar a um acoplamento indesejado.

## 2. Clean Tests (FIRST)
Testes devem ser tratados com o mesmo rigor de qualidade que o código de produção.

- **F: Fast**: Testes devem rodar em segundos. Se forem lentos, ninguém os rodará.
- **I: Independent**: Um teste não deve depender do resultado de outro teste ou do estado deixado por ele.
- **R: Repeatable**: Devem passar em qualquer ambiente (Local, CI, Docker), a qualquer momento.
- **S: Self-validating**: O teste deve ter um resultado booleano (Pass/Fail) claro. Nada de logs para verificar manualmente.
- **T: Timely**: Devem ser escritos preferencialmente no ciclo de TDD (ou junto com o código de produção).

---

## Estrutura de Teste Limpo (AAA)

Siga sempre o padrão **Arrange, Act, Assert** para legibilidade:

1.  **Arrange**: Configure o cenário, crie objetos e mocks.
2.  **Act**: Execute o método ou ação que está sendo testada.
3.  **Assert**: Verifique se o resultado foi o esperado.

**Dica**: Um bom teste deve contar uma história. Use nomes descritivos (ex: `test_deve_lancar_excecao_quando_usuario_for_nulo`).
