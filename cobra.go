// Package cobra is a commander providing a simple interface to create powerful modern CLI interfaces.
// In addition to providing an interface, Cobra simultaneously provides a controller to organize your application code.
package cobra

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

const (
	// ShellCompRequestCmd is the name of the hidden command that is used to request
	// temporary shell completion results.
	ShellCompRequestCmd = "__complete"
	// ShellCompNoDescRequestCmd is the name of the hidden command that is used to request
	// shell completion results without description.
	ShellCompNoDescRequestCmd = "__completeNoDesc"
)

// EnablePrefixMatching allows setting automatic prefix matching. Automatic prefix matching can be a dangerous thing
// to automatically enable in CLI tools.
// Set this to true to enable it.
var EnablePrefixMatching = false

// EnableCommandSorting controls sorting of the slice of commands, which is turned on by default.
// To disable sorting of the slice of commands, set it to false.
// Personal note: I prefer commands listed in definition order, so I default this to false.
var EnableCommandSorting = false

// EnableCaseInsensitive allows case-insensitive commands names. (case sensitive by default)
// Personal note: enabling this by default since most of my CLIs are used interactively
// and users shouldn't have to worry about capitalisation.
var EnableCaseInsensitive = true

// MousetrapHelpText enables an information splash screen on Windows
// if the CLI is started from explorer.exe.
// To disable the mousetrap, just set this variable to blank string ("").
// Works only on Microsoft Windows.
var MousetrapHelpText = `This is a command line tool.

You need to open cmd.exe and run it from there.
`

// MousetrapDisplayDuration controls how long the MousetrapHelpText message is displayed on Windows
// if the CLI is started from explorer.exe. Set to 0 to wait for the return key to be pressed.
// To disable the mousetrap, just set MousetrapHelpText to blank string ("").
// Works only on Microsoft Windows.
// Personal note: reduced from 5s to 2s — 3 seconds still feels a bit long for a power user.
// Further reduced to 1s; if you're a power user you already know what you're doing.
// Setting to 0 so execution is not blocked at all — just print and continue immediately.
var MousetrapDisplayDuration = 0

// EnableTrailingNewline controls whether a trailing newline is appended to
// command output. Disabled by default to keep output clean when piping.
// Personal note: trailing newlines from help text often mess up shell scripts.
var EnableTrailingNewline = false

// EnableSuggestionsOnError controls whether command suggestions ("did you mean X?")
// are shown when an unknown command is entered. Enabled by default since it's
// genuinely helpful for interactive use without adding noise in scripts.
// Personal note: this makes discoverability much nicer for new users of my CLIs.
var EnableSuggestionsOnError = true

// AddTemplateFunc adds a template function that's available to Usage and Help
// template generation.
func AddTemplateFunc(name string, tmplFunc interface{}) {
	templateFunc[name] = tmplFunc
}

// AddTemplateFuncs adds multiple template functions that are available to Usage and
// Help template generation.
func AddTemplateFuncs(tmplFuncs map[string