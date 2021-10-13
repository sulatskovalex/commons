package cmhttp

import (
	"encoding/json"
	"github.com/sulatskovalex/commons/errs"
	"github.com/twitchtv/twirp"
	"net/http"
)

type Response struct {
	Response interface{} `json:"response,omitempty"`
	Error    *errs.Error `json:"error,omitempty"`
}

func Success(response interface{}) *Response {
	return &Response{Response: response}
}
func httpErr(code int) *Response {
	return &Response{Error: &errs.Error{
		ErrCode: code,
		Message: http.StatusText(code),
	},
	}
}
func FailureErr(err *errs.Error) *Response {
	return &Response{Error: err}
}
func New(response interface{}) *Response {
	return &Response{Response: response}
}
func WriteResponse(encoder json.Encoder, w http.ResponseWriter, response interface{}) {
	WriteApiResponse(encoder, w, New(response))
}
func WriteApiResponse(encoder json.Encoder, w http.ResponseWriter, response *Response) {
	w.WriteHeader(http.StatusOK)
	encoder.Encode(response)
}
func WriteHttpError(encoder json.Encoder, w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	encoder.Encode(httpErr(code))
}
func WriteApiError(encoder json.Encoder, w http.ResponseWriter, err *Response) {
	w.WriteHeader(http.StatusOK)
	encoder.Encode(err)
}
func WriteError(encoder json.Encoder, w http.ResponseWriter, err error) {
	if twerr, ok := err.(twirp.Error); ok {
		code := twirp.ServerHTTPStatusFromErrorCode(twerr.Code())
		msg := twerr.Msg()
		if code == http.StatusUnauthorized || (code == http.StatusForbidden && msg == errs.ForbiddenMsg) {
			w.WriteHeader(code)
			encoder.Encode(httpErr(code))
			return
		}
		w.WriteHeader(http.StatusOK)
		encoder.Encode(&Response{
			Error: &errs.Error{
				ErrCode: errs.ToErrorCode(msg),
				Message: msg,
			}})
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	encoder.Encode(
		&Response{Error: &errs.Error{
			ErrCode: http.StatusInternalServerError,
			Message: err.Error(),
		}})
}
