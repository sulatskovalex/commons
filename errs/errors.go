package errs

import "github.com/twitchtv/twirp"

type Error struct {
	ErrCode int    `json:"code"`
	Message string `json:"message"`
}

const (
	InternalCode = 0
	InternalMsg  = "internal"

	NotFoundCode = 1
	NotFoundMsg  = "not found"

	UnknownContentTypeCode = 2
	UnknownContentTypeMsg  = "unknown content type"

	ForbiddenCode = 3
	ForbiddenMsg  = "forbidden"

	UnauthorizedCode = 4
	UnauthorizedMsg  = "unauthorized"
)

var (
	InternalErr           = twirp.NewError(twirp.Internal, InternalMsg)
	NotFoundErr           = twirp.NewError(twirp.NotFound, NotFoundMsg)
	UnknownContentTypeErr = twirp.NewError(twirp.InvalidArgument, UnknownContentTypeMsg)
	ForbiddenErr          = twirp.NewError(twirp.PermissionDenied, ForbiddenMsg)
	UnauthorizedErr       = twirp.NewError(twirp.Unauthenticated, UnauthorizedMsg)
)

func ToErrorCode(message string) int {
	switch message {
	case NotFoundMsg:
		return NotFoundCode
	case UnknownContentTypeMsg:
		return UnknownContentTypeCode
	case UnauthorizedMsg:
		return UnauthorizedCode
	case ForbiddenMsg:
		return ForbiddenCode
	}
	return InternalCode
}
