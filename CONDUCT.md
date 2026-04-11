## Cobra User Contract

### Versioning
Cobra will follow a steady release cadence. Non breaking changes will be released as minor versions quarterly. Patch bug releases are at the discretion of the maintainers. Users can expect security patch fixes to be released within relatively short order of a CVE becoming known. For more information on security patch fixes see the CVE section below. Releases will follow [Semantic Versioning](https://semver.org).
 this fork to explore how Cobra handles subcommand routing and persistent flags. Key areas of interest: `Command.PersistentPreRunE` andacts with auto-generated usage output.
>
> **Progress:** Completed initial exploration of `PersistentPreRunE` — confirmed that child commands inherit the hook unless that the hook chain does NOT automatically call parent hooks; you must do that manually if needed.
>
> **Progress:** Completed initial exploration of `SetHelpCommand` — when you provide a custom help command via `SetHelpCommand`, Cobra skips adding the auto-generated one, but the `-h`/`--help` flag is still wired up automatically through `InitDefaultHelpFlag`. The custom help command and the flag are independent; overriding one does not affect the other.
>
> **Progress:** Investigated `TraverseChildren` — when set to `true` on the root command, Cobra parses flags on parent commands before executing a child. Useful for cases where a persistent flag must be resolved before the child's `PreRunE` fires. Confirmed that `DisableFlagParsing` on a child command takes precedence and will bypass traversal for that specific command.
>
> **Next:** Look into `CompletionOptions` and how Cobra generates shell completion scripts. Specifically want to understand how custom completion functions registered via `RegisterFlagCompletionFunc` interact with the built-in `__complete` command.
