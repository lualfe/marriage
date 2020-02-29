package responsewriter

import (
	"net/http"

	"github.com/labstack/echo"
)

// Response model
type Response struct {
	Code     int
	Message  string
	Internal error
}

// JSON method treats the json object that will be formed
func (r Response) JSON(c echo.Context, body interface{}) {
	switch r.Code {
	case http.StatusOK:
		if body == nil {
			c.String(r.Code, "")
			return
		}
		c.JSON(r.Code, body)
		return
	default:
		c.JSON(r.Code, r)
		return
	}
}

// UnexpectedError returns an response struct from unexpected errors
func UnexpectedError(err error) Response {
	return Response{
		Code:     http.StatusInternalServerError,
		Message:  err.Error(),
		Internal: err,
	}
}

// Error returns an error from a request
func Error(message string, code int) Response {
	return Response{
		Code:     code,
		Message:  message,
		Internal: nil,
	}
}

// Redirect handles responses when redirects are needed
func Redirect(message string, code int) Response {
	if code < 300 || code >= 400 {
		return Response{}
	}
	return Response{
		Code:     code,
		Message:  message,
		Internal: nil,
	}
}

// Success returns an successful response
func Success() Response {
	return Response{
		Code:     http.StatusOK,
		Message:  "",
		Internal: nil,
	}
}
