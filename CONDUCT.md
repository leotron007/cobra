## Cobra User Contract

### Versioning
Cobra will follow a steady release cadence. Non breaking changes will be released as minor versions quarterly. Patch bug releases are at the discretion of the maintainers. Users can expect security patch fixes to be released within relatively short order of a CVE becoming known. For more information on security patch fixes see the CVE section below. Releases will follow [Semantic Versioning](https://semver.org/) command and the flag are independent; over — when set to `true` on the root command, Cobra parses flags on parent commands before executing a child. Useful for cases where a persistent flag must be resolved before the child's `PreRunE` fires. Confirmed that `DisableFlagParsing` on a child command takes precedence and will bypass traversal for that specific command.
>
> **Progress:** Investigated `CompletionOptions` and `RegisterFlagCompletionFunc` — custom completion functions are stored per-flag on the command and are invoked by the built-in `__complete` command at runtime. The `CompletionOptions.DisableDefaultCmd` field can suppress the auto-added completion command entirely if you don't need shell completion support.
>
> **Progress:** Investigated `PersistentPreRunE` hook inheritance — child commands inherit the hook unless overridden, and the hook chain does NOT automatically call parent hooks; you must do that manually if needed.
>
> **Progress:** Investigated `cobra.CheckErr` and `OnInitialize` — `CheckErr` is a thin wrapper that prints the error and calls `os.Exit(1)`; good for top-level main() error handling but avoid inside library code. `OnInitialize` registers functions that run before any command's `Execute`, making it the standard place to call `viper.ReadInConfig()` so config is loaded once regardless of which subcommand is invoked.
>
> **Next:** Look into `cobra.OnFinalize` (added in newer versions) and whether it's a reliable place to flush resources, vs deferring in `main()`.
