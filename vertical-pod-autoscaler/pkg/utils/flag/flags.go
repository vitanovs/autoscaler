package flags

import (
	"flag"
	"fmt"
	"strings"

	"k8s.io/utils/set"
)

var (
	_ flag.Value = &ChoiceFlagVar{}
)

// ChoiceFlagVar defines a command line flag with predefined allowed choices
type ChoiceFlagVar struct {
	value   string
	choices set.Set[string]
}

// NewChoiceFlagVar creates a new flags.ChoiceFlag
func NewChoiceFlagVar(choices set.Set[string]) ChoiceFlagVar {
	flag := ChoiceFlagVar{
		choices: choices,
	}
	return flag
}

// String returns the flag value
func (f *ChoiceFlagVar) String() string {
	return f.value
}

// Type returns the flag type
func (f *ChoiceFlagVar) Type() string {
	return "string"
}

// Set sets the flag value
func (f *ChoiceFlagVar) Set(value string) error {
	if !f.choices.Has(value) {
		choicesToStr := strings.Join(f.choices.UnsortedList(), ", ")
		return fmt.Errorf("expcted %s, but got %q", choicesToStr, value)
	}
	f.value = value
	return nil
}
