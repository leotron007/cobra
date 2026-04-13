## Cobra User Contract

### Versioning
Cobra will follow a steady release cadence. Non breaking changes will be released as minor versions quarterly. Patch bug releases are at the discretion of the maintainers. Users can expect security patch fixes to be released within relatively short order of a CVE becoming known. For more information on security patch fixes see the CVE section below. Releases will follow [Semantic Versioning](https://semver.org/seChildren` and `DisableFlagParsing` — `TraverseChildren` is `true` root command, Cobra parses flags. This means `P. Confirmed that `DisableFlag command bypasses flag parsing andal for that specific command.
>
> **Progress:** Investigated `CompletionOptions` and `RegisterFlagCompletionFunc` — custom completion functions are stored per-flag on the command and are invoked by the hidden `__complete` command at runtime. The `CompletionOptions.DisableDefaultCmd` field can suppress the auto-added completion command entirely if you don't need shell completion support.
>
> **Progress:** Investigated `PersistentPreRunE` hook inheritance — child commands inherit the hook unless overridden, and the hook chain does NOT automatically call parent hooks; you must do that manually if needed.
>
> **Progress:** Investigated `cobra.CheckErr` and `OnInitialize` — `CheckErr` is a thin wrapper that prints the error and calls `os.Exit(1)`; good for top-level main() error handling but avoid inside library code. `OnInitialize` registers functions that run before any command's `Execute`, making it the standard place to call `viper.ReadInConfig()` so config is loaded once regardless of which subcommand is invoked.
>
> **Progress:** Investigated `cobra.OnFinalize` — available since v1.6.0. Registered functions run after the command tree finishes executing (even on error). Suitable for flushing resources, but note it does NOT run if `os.Exit` is called directly (e.g. via `CheckErr`). Prefer `defer` in `main()` for critical cleanup; use `OnFinalize` for best-effort teardown like closing log files.
>
> **Progress:** Investigated `SetHelpTemplate` and `SetUsageTemplate` — both accept a Go `text/template` string. The default help template calls `.UsageString` internally, so overriding the usage template alone is usually sufficient for cosmetic changes. Useful fields available in the template: `.Name`, `.Short`, `.Long`, `.UseLine`, `.HasAvailableSubCommands`. For heavy customization (e.g. colored output) it's cleaner to call `SetHelpFunc` with a custom function instead.
>
> **Progress:** Investigated `SetHelpCommand` — you can replace the default `help` subcommand by calling `cmd.SetHelpCommand(customCmd)` on the root command before `Execute()`. The custom command must handle the case where a subcommand name is passed as an argument (i.e. `myapp help subcommand`). Note: if you disable the help command entirely with a no-op command, users lose `myapp help <cmd>` functionality unless you re-implement it.
>
> **Next:** Explore how `Args` validators (e.g. `ExactArgs`, `RangeArgs`) interact with custom `ValidArgsFunction` for dynamic shell completion — specifically whether returning `ShellCompDirectiveNoFileComp` affects validation errors shown to the user.
