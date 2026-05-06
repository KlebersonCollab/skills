# IDE Integration: Flutter with FVM

For your IDE to correctly recognize the Flutter SDK managed by FVM, follow these steps.

## 1. VS Code (Recommended)

Create or edit the `.vscode/settings.json` file in the project root:

```json
{
  "dart.flutterSdkPath": ".fvm/flutter_sdk",
  "search.exclude": {
    "**/.fvm": true
  },
  "files.exclude": {
    "**/.fvm": true
  }
}
```

**Tip**: Add `.vscode/` to your `.gitignore` if you want to keep personal settings out of the repository, or version it if you want to ensure everyone uses the local SDK.

## 2. Android Studio / IntelliJ

For these IDEs, you must point the Flutter SDK path manually:

1. Open the project in Android Studio.
2. Go to `Preferences` (macOS) or `Settings` (Windows).
3. Navigate to `Languages & Frameworks > Flutter`.
4. In the `Flutter SDK path` field, enter the full path to the `.fvm/flutter_sdk` folder of your project.
   - Example: `/Users/my_user/Projects/my_app/.fvm/flutter_sdk`
5. Click `Apply`.

## 3. IDE Terminal

Whenever you open your IDE's integrated terminal, remember to use the `fvm` prefix.

To make it easier, you can create an alias in your `.zshrc` or `.bashrc`:

```bash
alias fl="fvm flutter"
```

Now you can use `fl run` instead of `fvm flutter run`.

## 4. Troubleshooting in VS Code

If VS Code continues to warn that the SDK was not found:

1. Run `fvm use stable` in the terminal to ensure the `.fvm/flutter_sdk` link exists.
2. In VS Code, press `Cmd+Shift+P` (macOS) or `Ctrl+Shift+P` (Windows).
3. Type `Flutter: Change SDK`.
4. Select the option that points to the FVM path in your project.


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.366928Z"
evidence_checksum: "NONE"
```
