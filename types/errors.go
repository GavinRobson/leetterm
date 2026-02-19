package types

import "errors"

var Errors = struct {
	NoConfigFound error
}{
		NoConfigFound: errors.New("no config found"),
	}
