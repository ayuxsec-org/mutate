package mutate

import "runtime/debug"

var Version string

func VersionString() string {
	// injected with goreleaser
	if Version != "" {
		return Version
	}

	// bulit using go toolchain
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "(devel)" {
			return info.Main.Version
		}
	}

	// default
	return "dev"
}
