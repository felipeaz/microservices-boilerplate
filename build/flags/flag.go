package flags

import "flag"

type Flags struct {
	Debug *bool
}

func Build() Flags {
	var flags Flags

	flags.Debug = flag.Bool("debug", false, "enable/disable debug mode")
	flag.Parse()

	return flags
}
