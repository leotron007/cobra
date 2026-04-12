## Cobra User Contract

### Versioning
Cobra will follow a steady release cadence. Non breaking changes will be released as minor versions quarterly. Patch bug releases are at the discretion of the maintainers. Users can expect security patch fixes to be released within relatively short order of a CVE becoming known. For more information on security patch fixes see the CVE section below. Releases will follow [Semantic Versioning](https://semver.org)** — Using explore how Cobra handles subcommand Key areas of interest: `Command. and how it interacts with auto-generated usage output inherit the hook unless overridden, and that the hook chain does NOT automatically call parent hooks; you must do that manually if needed.
>
> **Progress:** Completed initial exploration of `SetHelpCommand` — when you provide a custom help command via `SetHelpCommand`, Cobra skips adding the auto-generated one, but the `-h`/`--help` flag is still wired up automatically through `InitDefaultHelpFlag`. The custom help command and the flag are independent; overriding one does not affect the other.
>
> **Progress:** Investigated `TraverseChildren` — when set to `true` on the root command, Cobra parses flags on parent commands before executing a child. Useful for cases where a persistent flag must be resolved before the child's `PreRunE` fires. Confirmed that `DisableFlagParsing` on a child command takes precedence and will bypass traversal for that specific command.
>
> **Progress:** Investigated `CompletionOptions` and `RegisterFlagCompletionFunc` — custom completion functions are stored per-flag on the command and are invoked by the built-in `__complete` command at runtime. The `CompletionOptions.DisableDefaultCmd` field can suppress the auto-added completion command entirely if you don't need shell completion support.
>
> **Next:** Dig into `cobra.CheckErr` and the global `OnInitialize` hook — want to understand the intended pattern for wiring up config file loading (e.g. with viper) before any command runs.
