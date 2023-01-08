package command

import (
	"fmt"
)

type idOptions struct {
	id int32
}

var idOpts idOptions

func (opts idOptions) validate() error {
	if opts.id <= 0 {
		return fmt.Errorf("invalid pull request ID")
	}

	return nil
}

