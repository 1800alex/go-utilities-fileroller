package fileroller

import (
	"github.com/spf13/afero"
)

var StubDepFS = afero.NewMemMapFs()
var oldDepFS = DepFS

func Stubs() {
	DepFS = StubDepFS
}

func StubsRestore() {
	StubDepFS = afero.NewMemMapFs() /* Get a new copy */
	DepFS = oldDepFS
}
