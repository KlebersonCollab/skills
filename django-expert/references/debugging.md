# Systematic Debugging in Django

Senior developers don't "try until it works", they debug systematically.

## 1. Indispensable Tools

- **Django Debug Toolbar**: Essential for analyzing SQL queries and template performance.
- **IPDB / PDB**: Use `breakpoint()` to pause execution and inspect the state.
- **Django Extensions (Shell Plus)**: To quickly test ORM logic in the terminal.

## 2. The Debugging Method

1.  **Reproduce**: Have a test or step-by-step process that fails consistently.
2.  **Isolate**: Reduce the problem to the smallest possible piece of code.
3.  **Inspect**: Use `logging.debug()` or the Debug Toolbar to see the real data.
4.  **Fix and Validate**: After the fix, ensure the test now passes.

## 3. Professional Logging

Configure `LOGGING` in `settings.py` to capture errors in production:

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
