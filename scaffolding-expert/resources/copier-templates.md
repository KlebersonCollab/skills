# Copier Templates Guide

## What is Copier?

Copier is a library and CLI app for writing project templates. It works with smart file copying, rendering templates, and project configuration.

## Installation

```bash
uvx copier copy <template-url> <destination>
```

## Template Structure

```
my-template/
├── copier.yaml           # Template configuration
├── {{project_name}}/     # Project directory (variable)
│   ├── {{module_name}}/ # Module directory (variable)
│   │   └── __init__.py
│   └── pyproject.toml
└── README.md
```

## copier.yaml Configuration

```yaml
# Template metadata
project_name:
  type: str
  help: Name of the project

module_name:
  type: str
  help: Name of the Python module
  default: "{{ project_name|lower|replace(' ', '_') }}"

author:
  type: str
  help: Author name
  default: "Your Name"

license:
  type: str
  help: Project license
  default: "MIT"
  choices:
    - MIT
    - Apache-2.0
    - GPL-3.0

include_tests:
  type: bool
  help: Include test directory
  default: true

# Conditional file inclusion
_templates_suffix: ""
_exclude:
  - "*.pyc"
  - "__pycache__"
  - ".git"
```

## Variable Types

### String
```yaml
project_name:
  type: str
  help: Project name
  default: "my-project"
```

### Boolean
```yaml
include_tests:
  type: bool
  help: Include tests
  default: true
```

### Choice
```yaml
license:
  type: str
  help: License type
  choices:
    - MIT
    - Apache-2.0
    - GPL-3.0
```

### Multiple Choice
```yaml
features:
  type: list
  help: Features to include
  choices:
    - auth
    - database
    - api
  default: []
```

## Template Filters

### String Manipulation
```yaml
# Lowercase
module_name: "{{ project_name|lower }}"

# Replace spaces
slug: "{{ project_name|replace(' ', '-') }}"

# Capitalize
class_name: "{{ project_name|capitalize }}"

# Title case
display_name: "{{ project_name|title }}"
```

### Conditional
```yaml
# Conditional inclusion
{% if include_tests %}
tests/
{% endif %}
```

## File Templates

### Python File
```python
# {{module_name}}/__init__.py
"""
{{project_name}}
{{'=' * project_name|length}}

{{description}}
"""

__version__ = "0.1.0"
__author__ = "{{author}}"
```

### Configuration File
```toml
# pyproject.toml
[project]
name = "{{project_name}}"
version = "0.1.0"
description = "{{description}}"
authors = [{name = "{{author}}"}]

[project.license]
text = "{{license}}"
```

### README
```markdown
# {{project_name}}

{{description}}

## Installation

```bash
pip install {{project_name}}
```

## Usage

```python
import {{module_name}}

# Your code here
```

## License

{{license}}

## Author

{{author}}
```

## Conditional Files

### Using _exclude
```yaml
_exclude:
  - "*.pyc"
  - "__pycache__"
  - ".git"
  - "tests/"  # Exclude tests directory
```

### Using _include
```yaml
_include:
  - "src/"
  - "tests/"
```

### Using Conditional Directories
```
{% if include_tests %}
tests/
    __init__.py
    test_{{module_name}}.py
{% endif %}
```

## Advanced Features

### Multi-Choice Templates
```yaml
database:
  type: str
  help: Database to use
  choices:
    - postgresql
    - mysql
    - sqlite
    - none
```

### Template Dependencies
```yaml
# copier.yaml
dependencies:
  - copier-template-configuration
```

### Pre-commit Hooks
```yaml
# copier.yaml
_tasks:
  - "uv sync"
  - "pre-commit install"
```

## Example: FastAPI Template

### copier.yaml
```yaml
project_name:
  type: str
  help: FastAPI project name

module_name:
  type: str
  help: Python module name
  default: "{{ project_name|lower|replace(' ', '_') }}"

database:
  type: str
  help: Database choice
  choices:
    - postgresql
    - sqlite
  default: sqlite

include_auth:
  type: bool
  help: Include authentication
  default: false

include_tests:
  type: bool
  help: Include tests
  default: true
```

### Template Files
```
fastapi-template/
├── copier.yaml
├── {{project_name}}/
│   ├── {{module_name}}/
│   │   ├── __init__.py
│   │   ├── main.py
│   │   └── models.py
│   ├── tests/
│   │   └── test_main.py
│   └── pyproject.toml
└── README.md
```

## Best Practices

1. **Use descriptive variable names** - Make it clear what each variable does
2. **Provide sensible defaults** - Reduce user input where possible
3. **Use choices for enums** - Restrict input to valid options
4. **Document your template** - Include a README with usage instructions
5. **Test your template** - Verify it works with different inputs
6. **Version your template** - Use git tags for versioning

## Resources

- [Copier Documentation](https://copier.readthedocs.io/)
- [Jinja2 Template Designer](https://jinja.palletsprojects.com/en/3.1.x/templates/)
- [Python Project Templates](https://github.com/copier-org/copier/tree/master/copier/templates)