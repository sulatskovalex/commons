package errs

import "github.com/twitchtv/twirp"

const (
	NotFoundMsg = "not found"
)

var (
	NotFoundErr = twirp.NewError(twirp.NotFound, NotFoundMsg)
)
