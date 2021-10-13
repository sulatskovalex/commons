package errs

import "github.com/twitchtv/twirp"

const (
	NotFoundMsg           = "not found"
	UnknownContentTypeMsg = "unknown content type"
)

var (
	NotFoundErr           = twirp.NewError(twirp.NotFound, NotFoundMsg)
	UnknownContentTypeErr = twirp.NewError(twirp.InvalidArgument, UnknownContentTypeMsg)
)
