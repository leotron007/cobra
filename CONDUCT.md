## Cobra User Contract

### Versioning
Cobra will follow a steady release cadence. Non breaking changes will be released as minor versions quarterly. Patch bug releases are at the discretion of the maintainers. Users can expect security patch fixes to be released within relatively short order of a CVE becoming known. For more information on security patch fixes see the CVE section below. Releases will follow [Semantic Versioning](https://semver.org). Users should not expect unpredict. A low severity CVE will use GitHub issues and the [#cobra slack of Cobra.

There may be times, however, when master is closed for breaking changes. This is likely to happen as we near the release of a new version.

Breaking changes are not allowed in release branches, as these represent minor versions that have already been released. These version have consumers who expect the APIs, behaviors, etc, to remain stable during the lifetime of the patch stream for the minor release.

Examples of breaking changes include:
- Removing or renaming exported constant, variable, type, or function.
- Updating the version of critical libraries such as `spf13/pflag`, `spf13/viper` etc...
  - Some version updates may be acceptable for picking up bug fixes, but maintainers must exercise caution when reviewing.

There may, at times, need to be exceptions where breaking changes are allowed in release branches. These are at the discretion of the project's maintainers, and must be carefully considered before merging.

### CI Testing
Maintainers will ensure the Cobra test suite utilizes the current supported versions of Golang.

### Disclaimer
Changes to this document and the contents therein are at the discretion of the maintainers.

---

> **Personal fork note:** This is a personal fork for learning purposes. The contract above reflects the upstream project's policies. Any experiments or deviations in this fork are not intended for production use.
>
> **Study notes:** I'm using this fork to explore how Cobra handles subcommand routing and persistent flags. Key areas of interest: `Command.PersistentPreRunE`, flag inheritance, and custom help templates.
