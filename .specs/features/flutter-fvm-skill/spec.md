# Spec: Flutter with FVM Skill

## 1. Goal
Create a specialized skill for Flutter development using FVM (Flutter Version Management) as the version manager. This skill should mirror the quality and structure of the `python-uv` skill, providing a standard, efficient, and professional workflow for Flutter projects.

## 2. Target Audience
Software engineers using Gemini CLI to develop Flutter applications, requiring multi-version support and consistent project environments.

## 3. Core Requirements (Functional)
- **FVM Integration**: Detailed instructions on using FVM (`fvm use`, `fvm install`, `fvm flutter ...`).
- **Project Lifecycle**:
    - **Environment**: SDK installation, version pinning, doctor checks.
    - **Project**: Initialization (`fvm flutter create`), directory structure, `pubspec.yaml` management.
    - **Develop**: Dependency management (`fvm flutter pub add/get`), running/debugging, linting (`fvm flutter analyze`), formatting (`dart format`).
    - **Test**: Unit, widget, and integration testing.
    - **Deploy**: Building for multiple platforms (Android, iOS, Web, Desktop).
- **Command Comparison**: Table comparing standard `flutter` commands vs. `fvm flutter`.
- **Best Practices**: State management (BLoC, Riverpod, Provider), Clean Architecture, CI/CD with FVM.
- **Troubleshooting**: Common FVM issues (PATH, IDE config, SDK mismatches).

## 4. Non-Functional Requirements
- **Consistency**: Follow the same tone, style, and structure as the `python-uv` skill.
- **Brevity**: Concise explanations and clear command examples.
- **Correctness**: Ensure all commands are valid for the latest stable Flutter and FVM versions.

## 5. Acceptance Criteria (AC)
- [ ] `SKILL.md` exists in a `flutter-fvm/` directory.
- [ ] `README.md` and `CHANGELOG.md` are present.
- [ ] `references/` directory contains detailed guides for Environment, Project, and Deploy.
- [ ] The skill provides a clear "When to Use" section.
- [ ] The skill includes a "Workflow" section with 4 phases.
- [ ] A command comparison table between `flutter` and `fvm flutter` is included.
- [ ] Best practices for `pubspec.yaml` and state management are documented.
- [ ] VS Code and IntelliJ/Android Studio configuration for FVM is mentioned.

## 6. Architecture & Design
The skill will be organized as follows:
- `flutter-fvm/SKILL.md`: Main entry point.
- `flutter-fvm/README.md`: Basic introduction.
- `flutter-fvm/CHANGELOG.md`: Version history.
- `flutter-fvm/references/`:
    - `environment-setup.md`: FVM installation, PATH, SDK management.
    - `project-structure.md`: Folder organization, `pubspec.yaml`.
    - `dependency-management.md`: Pub commands, versioning.
    - `testing-and-quality.md`: Analyze, test, coverage.
    - `deployment-guide.md`: Build commands for different platforms.
    - `ide-integration.md`: VS Code and Android Studio setup.
- `flutter-fvm/examples/`: Sample `pubspec.yaml`, `analysis_options.yaml`, and CI/CD pipelines.
