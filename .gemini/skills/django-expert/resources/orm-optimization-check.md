# Checklist: Otimização do ORM Django

Use esta checklist para identificar e corrigir problemas de performance em suas queries Django.

## 1. Problemas de N+1
- [ ] **Identificação**: O número de queries aumenta proporcionalmente ao número de itens na lista?
- [ ] **Solução (ForeignKey/OneToOne)**: Use `.select_related('nome_do_campo')` para realizar um JOIN SQL.
- [ ] **Solução (ManyToMany/Reverse ForeignKey)**: Use `.prefetch_related('nome_do_campo')` para realizar queries separadas e fazer o join no Python.

## 2. Seleção de Campos (Payload)
- [ ] **Apenas campos necessários**: Use `.only('campo1', 'campo2')` para carregar apenas colunas específicas.
- [ ] **Exclusão de campos pesados**: Use `.defer('campo_pesado')` para campos como `TextField` ou `BinaryField` que não são necessários de imediato.
- [ ] **Dicionários/Tuplas**: Use `.values()` ou `.values_list()` quando não precisar de instâncias de modelos (muito mais rápido).

## 3. Agregações e Cálculos
- [ ] **Cálculo no Banco**: Use `.annotate()` com funções de agregação (`Sum`, `Count`, `Avg`) em vez de calcular em loops Python.
- [ ] **Contagem Simples**: Use `.exists()` para verificar presença e `.count()` em vez de `len(queryset)`.

## 4. Querysets Lazy
- [ ] **Evite avaliação prematura**: Querysets são preguiçosos. Não os converta em listas (`list(qs)`) até que seja estritamente necessário.
- [ ] **Iteração Eficiente**: Use `.iterator()` para processar grandes volumes de dados sem carregar tudo na memória.

## 5. Debugging Tools
- [ ] **django-debug-toolbar**: Essencial para ver o SQL gerado em tempo real no browser.
- [ ] **nplusone**: Biblioteca para detectar automaticamente problemas de N+1 durante o desenvolvimento.
- [ ] **Query Inspector**: Logue o SQL no console durante os testes ou desenvolvimento local:
  ```python
  import logging
  l = logging.getLogger('django.db.backends')
  l.setLevel(logging.DEBUG)
  l.addHandler(logging.StreamHandler())
  ```

## Regra de Ouro
> "Mova a lógica de dados para o banco de dados e a lógica de apresentação para o Python."
