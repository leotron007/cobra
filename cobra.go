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
var EnableCommandSorting = true

// EnableCaseInsensitive allows case-insensitive commands names. (case sensitive by default)
var EnableCaseInsensitive = false

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
var MousetrapDisplayDuration = 5000000000 // 5 seconds

// AddTemplateFunc adds a template function that's available to Usage and Help
// template generation.
func AddTemplateFunc(name string, tmplFunc interface{}) {
	templateFunc[name] = tmplFunc
}

// AddTemplateFuncs adds multiple template functions that are available to Usage and
// Help template generation.
func AddTemplateFuncs(tmplFuncs map[string]interface{}) {
	for k, v := range tmplFuncs {
		templateFunc[k] = v
	}
}

// OnInitialize sets the passed functions to be run when each command's
// Execute method is called.
func OnInitialize(y ...func()) {
	initializers = append(initializers, y...)
}

// OnFinalize sets the passed functions to be run when each command's
// Execute method is completed.
func OnFinalize(y ...func()) {
	finalizers = append(finalizers, y...)
}

// FIXME Workaround for https://github.com/spf13/cobra/issues/671
func stringContainsSeparator(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '_' && r != '-' {
			return true
		}
	}
	return false
}

func valueToString(v interface{}) string {
	switch vv := v.(type) {
	case string:
		return vv
	case bool:
		return strconv.FormatBool(vv)
	case int:
		return strconv.Itoa(vv)
	case int64:
		return strconv.FormatInt(vv, 10)
	case float64:
		return strconv.FormatFloat(vv, 'f', -1, 64)
	case []string:
		return strings.Join(vv, ", ")
	default:
		return fmt.Sprintf("%v", reflect.ValueOf(v))
	}
}
