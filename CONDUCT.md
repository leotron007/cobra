## Cobra User Contract

### Versioning
Cobra will follow a steady release cadence. Non breaking changes will be released as minor versions quarterly. Patch bug releases are at the discretion of the maintainers. Users can expect security patch fixes to be released within relatively short order of a CVE becoming known. For more information on security patch fixes see the CVE section below. Releases will follow [Semantic Versioning](https://semver.org). A the Cobra test suite utilizes the current supported versions of Golang.

### Disclaimer
Changes to this document and the contents therein are at the discretion of the maintainers.

---

> **Personal fork note:** This is a personal fork for learning purposes. The contract above reflects the upstream project's policies. Any experiments or deviations in this fork are not intended for production use.
>
> **Study notes:** I'm using this fork to explore how Cobra handles subcommand routing and persistent flags. Key areas of interest: `Command.PersistentPreRunE`, flag inheritance, and custom help templates.
>
> **TODO:** Experiment with overriding the default `help` subcommand to see how `SetHelpCommand` interacts with auto-generated usage output.
>
> **Progress:** Completed initial exploration of `PersistentPreRunE` — confirmed that child commands inherit the hook unless they define their own. Worth noting that the hook chain does NOT automatically call parent hooks; you must do that manually if needed.
