package globals

import (
	"os"
	"path/filepath"

	"github.com/arduino/arduino101load/version"
)

var (
	// VersionInfo contains all info injected during build
	VersionInfo = version.NewInfo(filepath.Base(os.Args[0]))
)
