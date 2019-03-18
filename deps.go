package fileroller

import (
	"github.com/spf13/afero"
)

// DepFS is an external dependancy variable
// set to an instance of afero.FS.
var DepFS = afero.NewOsFs()
