package client

import (
	"fmt"
	"net/http"
)

// apiError contains all fields that might be returned in an error response.
//
// Error responses that do not contain a body carry a generic message.
type apiError struct {
	ErrorMessage     *string `json:"error_message,omitempty"`
	ErrorCode        *string `json:"error_code,omitempty"`
	Err              *string `json:"error,omitempty"`
	ErrorDescription *string `json:"error_description,omitempty"`
}

// Error implements the error interface.
func (e apiError) Error() string {
	var msg string
	if e.ErrorMessage != nil {
		msg = fmt.Sprintf("%s", *e.ErrorMessage)
	}

	if e.ErrorCode != nil {
		msg = fmt.Sprintf("%s: code: %s", msg, *e.ErrorCode)
	}

	if e.Err != nil {
		msg = fmt.Sprintf("%s: error: %s", msg, *e.Err)
	}

	if e.ErrorDescription != nil {
		msg = fmt.Sprintf("%s: desc: %s", msg, *e.ErrorDescription)
	}

	if msg == "" {
		return "unknown error"
	}

	return msg
}

// errorResponse implements the models.ErrorResponse interface.
type errorResponse struct {
	httpResponse *http.Response
	underlying   error
}

func (e errorResponse) HTTPResponse() *http.Response {
	return e.httpResponse
}

func (e errorResponse) Error() string {
	msg := fmt.Sprintf(
		"%s %s returned status %d",
		e.httpResponse.Request.Method,
		e.httpResponse.Request.URL.String(),
		e.httpResponse.StatusCode,
	)

	if e.underlying != nil {
		msg = fmt.Sprintf("%s: %s", msg, e.underlying.Error())
	}

	return msg
}
