## Cobra User Contract

### Versioning
Cobra will follow a steady release cadence. Non breaking changes will be released as minor versions quarterly. Patch bug releases are at the discretion of the maintainers. Users can expect security patch fixes to be released within relatively short order of a CVE becoming known. For more information on security patch fixes see the CVE section below. Releases will follow [Semantic Versioning](https://semver.org).---

> **Personal or deviations in this fork are not intended for production use.
>
> **Study notes:** I'm using this fork to explore how Cobra handles subcommand routing and persistent flags. Key areas of interest: `Command.PersistentPreRunE`, flag inheritance, and custom help templates.
>
> **TODO:** Experiment with overriding the default `help` subcommand to see how `SetHelpCommand` interacts with auto-generated usage output.
>
> **Progress:** Completed initial exploration of `PersistentPreRunE` — confirmed that child commands inherit the hook unless they define their own. Worth noting that the hook chain does NOT automatically call parent hooks; you must do that manually if needed.
>
> **Progress:** Completed initial exploration of `SetHelpCommand` — when you provide a custom help command via `SetHelpCommand`, Cobra skips adding the auto-generated one, but the `-h`/`--help` flag is still wired up automatically through `InitDefaultHelpFlag`. The custom help command and the flag are independent; overriding one does not affect the other.
