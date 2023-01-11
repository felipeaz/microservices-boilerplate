package flags

import "flag"

const (
	flagName = "debug"
	usage    = "enable/disable debug mode"
)

type Flags struct {
	Debug *bool
}

func Build() Flags {
	var flags Flags

	flags.Debug = flag.Bool(flagName, false, usage)
	flag.Parse()

	return flags
}
