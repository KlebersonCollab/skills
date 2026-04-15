# SOLID Principles: Deep Dive

Os cinco princípios fundamentais de design orientado a objetos que garantem a manutenibilidade e escalabilidade do software.

---

## 1. S: Single Responsibility Principle (SRP)
**"Uma classe deve ter apenas uma razão para mudar."**

- **Violation**: Classes "God" que lidam com lógica de UI, acesso a dados e regras de negócio.
- **Solution**: Decompor classes em componentes menores e especializados.
- **Check**: Se você tem dificuldade em nomear a classe de forma concisa ou se ela tem múltiplos "And" em sua descrição, ela provavelmente quebra o SRP.

## 2. O: Open/Closed Principle (OCP)
**"Entidades de software devem estar abertas para extensão, mas fechadas para modificação."**

- **Violation**: Uso excessivo de `switch` ou `if-else` baseados em tipos, forçando a alteração da lógica original para cada novo tipo adicionado.
- **Solution**: Utilizar interfaces, polimorfismo e padrões como *Strategy* ou *Decorator*.
- **Check**: Ao adicionar uma nova funcionalidade, eu preciso mexer no código que já está funcionando? Se sim, você quebrou o OCP.

## 3. L: Liskov Substitution Principle (LSP)
**"Subclasses devem ser substituíveis por suas superclasses sem quebrar o comportamento do programa."**

- **Violation**: Sobrescrever um método para lançar uma exceção de "Não implementado" ou forçar verificações de tipo (`is instance of`).
- **Solution**: Garantir que as subclasses respeitem o contrato (pré-condições e pós-condições) definido pela superclasse.
- **Check**: Se eu trocar a classe base por uma filha em qualquer lugar, o sistema continua estável?

## 4. I: Interface Segregation Principle (ISP)
**"Muitas interfaces específicas são melhores que uma interface única e genérica."**

- **Violation**: Uma interface "Gorda" que obriga uma classe a implementar métodos que ela não utilizará.
- **Solution**: Segregar interfaces em unidades funcionais menores e coesas.
- **Check**: Minha classe está implementando métodos vazios apenas para satisfazer a interface?

## 5. D: Dependency Inversion Principle (DIP)
**"Dependa de abstrações, não de implementações concretas."**

- **Violation**: Instanciar dependências diretamente dentro de uma classe (ex: `var repo = new SqlDatabase()`).
- **Solution**: Injeção de dependência e programação voltada a interfaces.
- **Check**: Posso trocar meu banco de dados ou serviço externo sem precisar recompilar as classes de negócio?

---

## Boas Práticas

- **Não exagere**: Aplicar todos os princípios de forma cega em projetos pequenos pode levar a um excesso de engenharia.
- **Contexto**: O SOLID é uma ferramenta para reduzir o custo de mudança; use-o onde a mudança é provável.
- **Simplicidade**: Se o SOLID tornar o código impossível de ler, revise sua implementação do KISS.
