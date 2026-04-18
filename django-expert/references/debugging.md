# Systematic Debugging in Django

Desenvolvedores seniores não "tentam até funcionar", eles debugaram sistematicamente.

## 1. Ferramentas Indispensáveis

- **Django Debug Toolbar**: Essencial para analisar queries SQL e performance de templates.
- **IPDB / PDB**: Use `breakpoint()` para pausar a execução e inspecionar o estado.
- **Django Extensions (Shell Plus)**: Para testar lógica de ORM rapidamente no terminal.

## 2. O Método de Depuração

1.  **Reproduzir**: Tenha um teste ou passo-a-passo que falha de forma consistente.
2.  **Isolar**: Reduza o problema até o menor pedaço de código possível.
3.  **Inspecionar**: Use `logging.debug()` ou o Debug Toolbar para ver os dados reais.
4.  **Corrigir e Validar**: Após o fix, garanta que o teste agora passa.

## 3. Logging Profissional

Configure o `LOGGING` no `settings.py` para capturar erros em produção:

```python
LOGGING = {
    'version': 1,
    'handlers': {
        'console': {'class': 'logging.StreamHandler'},
    },
    'loggers': {
        'django': {'handlers': ['console'], 'level': 'INFO'},
    },
}
```
