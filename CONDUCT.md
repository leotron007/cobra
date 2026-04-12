## Cobra User Contract

### Versioning
Cobra will follow a steady release cadence. Non breaking changes will be released as minor versions quarterly. Patch bug releases are at the discretion of the maintainers. Users can expect security patch fixes to be released within relatively short order of a CVE becoming known. For more information on security patch fixes see the CVE section below. Releases will follow [Semantic Versioning](https://semver.org/) **Progress:** Investigated `TraverseChildren` — when set to `true` on the root command, Cobra parses flags on parent commands before executing a the child's `PreRunE` fires. Confirmed that `DisableFlagParsing` on a child command takes precedence and will bypass traversal for that specific command.
>
CompletionOptions` and `RegisterFlagCompletionFunc` — custom completion functions are stored per-flag on the command and are invoked by the built-in `__complete` command at runtime. The `CompletionOptions.DisableDefaultCmd` field can suppress the auto-added completion command entirely if you don't need shell completion support.
>
> **Progress:** Investigated `PersistentPreRunE` hook inheritance — child commands inherit the hook unless overridden, and the hook chain does NOT automatically call parent hooks; you must do that manually if needed.
>
> **Progress:** Investigated `cobra.CheckErr` and `OnInitialize` — `CheckErr` is a thin wrapper that prints the error and calls `os.Exit(1)`; good for top-level main() error handling but avoid inside library code. `OnInitialize` registers functions that run before any command's `Execute`, making it the standard place to call `viper.ReadInConfig()` so config is loaded once regardless of which subcommand is invoked.
>
> **Progress:** Investigated `cobra.OnFinalize` — available since v1.6.0. Registered functions run after the command tree finishes executing (even on error). Suitable for flushing resources, but note it does NOT run if `os.Exit` is called directly (e.g. via `CheckErr`). Prefer `defer` in `main()` for critical cleanup; use `OnFinalize` for best-effort teardown like closing log files.
>
> **Next:** Explore custom help templates (`SetHelpTemplate`) and usage templates to see how far the built-in formatting can be pushed before needing a fully custom help command.
