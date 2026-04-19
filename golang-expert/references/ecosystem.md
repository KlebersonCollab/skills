# Pilar 5: Ecosystem & Samber

> O ecossistema Go evoluiu rapidamente com a introdução de Generics e logging estruturado nativo. Conhecer as ferramentas certas diferencia um desenvolvedor de um Expert.

---

## 🚀 Modernizing Go (Modernização)

Go 1.18 (Generics) e 1.21 (slog) mudaram as regras do jogo:
- **Generics**: Utilize tipos genéricos para criar estruturas de dados e funções reutilizáveis sem sacrificar a segurança de tipos (ex: bibliotecas de utilitários).
- **Native Structured Logging**: Substitua `logrus` ou `zap` pelo pacote nativo `slog` sempre que possível para reduzir dependências externas.
- **errors.Join**: Utilize a nova funcionalidade de agregação de múltiplos erros do Go 1.20+.

---

## 🧩 Samber Ecosystem (Produtividade Expert)

A suíte de bibliotecas `samber` é referência em Go moderno por utilizar Generics para resolver problemas comuns com elegância:

### 1. samber/lo (Lodash para Go)
O canivete suíço para manipulação de coleções:
- **Map / Filter / Reduce**: Transforme slices e maps sem loops verbosos.
- **Find / Uniq / GroupBy**: Operações complexas de busca e agrupamento em uma única linha.

### 2. samber/mo (Monads & Funcional)
Traga padrões de programação funcional para Go de forma idiomática:
- **Option[T]**: Lide com valores nulos/ausentes sem `nil` checks constantes.
- **Result[T]**: Encapsule resultados e erros em uma única estrutura.

### 3. samber/do (Dependency Injection)
Injeção de dependência simples e baseada em tipos:
- **Service Discovery**: Registre e recupere dependências sem reflexão pesada ou configurações globais complexas.
- **Lifecycle Management**: Gerencie o ciclo de vida (inicialização/finalização) de seus serviços.

### 4. samber/oops (Erro Rico)
Tratamento de erros de nível industrial:
- **Stack Traces**: Obtenha o rastro exato de onde o erro ocorreu.
- **Contextual Data**: Anexe metadados (User ID, Request ID) aos erros automaticamente.

---

## 📚 Outras Bibliotecas Populares

- **Echo / Chi / Gin**: Frameworks e roteadores HTTP para todos os gostos e necessidades.
- **sqlc**: Compila SQL puro em código Go tipado — a melhor forma de interagir com bancos SQL hoje.
- **GoMock / Testify**: Essenciais para garantir a testabilidade de sistemas complexos.

---

## 📈 Staying Updated (Evolução Contínua)

O ecossistema Go não para. Um Expert deve acompanhar:
- **Go Blog**: O canal oficial para anúncios de novas versões e propostas de mudanças (Gopher proposals).
- **Go Weekly**: Newsletter semanal com os melhores artigos e bibliotecas da comunidade.
- **Github samber/cc-skills-golang**: O repositório base que inspirou esta skill, sempre atualizado com novos padrões.

---

## 💡 Dica Expert
Não tente reinventar a roda. Se precisar de uma função utilitária para slices, verifique primeiro o `samber/lo`. Se precisar de tratamento de erros robusto, vá de `samber/oops`. A produtividade vem de saber o que usar e quando usar.
