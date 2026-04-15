# Architectural Principles — Referência Completa

Princípios fundamentais para criar sistemas sustentáveis e robustos.

---

## 1. Simplicidade (The Core Principle)
> "Simplicidade é a sofisticação máxima." (Leonardo da Vinci)

- Comece com o mínimo necessário para resolver o problema.
- Adicione complexidade apenas quando houver evidências de necessidade (escala, performance, segurança).
- Lembre-se: É muito mais barato adicionar complexidade depois do que remover complexidade desnecessária agora.

## 2. DRY (Don't Repeat Yourself)
- Evite duplicar lógica de negócio ou conhecimento no sistema.
- Cada peça de conhecimento deve ter uma representação única e inequívoca dentro de um sistema.

## 3. KISS (Keep It Simple, Stupid)
- A maioria dos sistemas funciona melhor se forem mantidos simples em vez de complicados.
- Evite abstrações prematuras.

## 4. YAGNI (You Ain't Gonna Need It)
- Não implemente funcionalidades ou infraestrutura até que sejam realmente necessárias.
- Aplicar o YAGNI reduz o "ruído" arquitetural e o custo de manutenção.

## 5. SOLID
- **S**: Single Responsibility Principle (Uma classe deve ter um único motivo para mudar).
- **O**: Open/Closed Principle (Entidades abertas para extensão, fechadas para modificação).
- **L**: Liskov Substitution Principle (Subtipos devem ser substituíveis por seus tipos base).
- **I**: Interface Segregation Principle (Muitas interfaces específicas são melhores que uma geral).
- **D**: Dependency Inversion Principle (Dependa de abstrações, não de implementações).
