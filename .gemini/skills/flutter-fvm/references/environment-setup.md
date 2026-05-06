# Environment Setup (FVM)

This guide details how to set up the Flutter development environment using FVM.

## 1. FVM Installation

FVM can be installed in several ways:

### macOS / Linux
```bash
brew tap leoafarias/fvm
brew install fvm
```

### Windows
```powershell
choco install fvm
# Or via Pub
dart pub global activate fvm
```

## 2. Configuring the PATH

Ensure the Dart/Pub binaries directory is in your PATH if you installed via Pub:

```bash
# macOS/Linux (zsh)
export PATH="$PATH":"$HOME/.pub-cache/bin"
```

## 3. Managing SDK Versions

### Install a specific version
```bash
fvm install 3.19.0
```

### Install the stable version
```bash
fvm install stable
```

### List installed versions
```bash
fvm list
```

## 4. Pinning the Version to the Project

In the root directory of your Flutter project, run:

```bash
fvm use stable
```

This will create a `.fvm` folder with the `fvm_config.json` file and a `flutter_sdk` symbolic link pointing to the selected version.

## 5. Health Check (FVM Doctor)

Run the command below to check if everything is configured correctly for FVM:

```bash
fvm doctor
```

To check the Flutter state via FVM:

```bash
fvm flutter doctor
```

## 6. Recommended .gitignore

Add the following to your `.gitignore` so as not to version the local SDK cache:

```gitignore
.fvm/flutter_sdk
```

But **do not** add `fvm_config.json`; it must be versioned so the team uses the same version.
## Environment Checklist
- [ ] Is the `fvm_config.json` file versioned?
- [ ] Is the `.fvm/flutter_sdk` folder in `.gitignore`?
- [ ] Did you run `fvm flutter doctor` and are all dependencies OK?
- [ ] Is the Flutter version in the project the same as in `fvm list`?


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.366473Z"
evidence_checksum: "NONE"
```
