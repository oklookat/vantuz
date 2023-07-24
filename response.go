package vantuz

import (
	"net/http"
)

func newResponse(req *Request, resp *http.Response) *Response {
	return &Response{
		Response: resp,
		req:      req,
	}
}

type Response struct {
	*http.Response

	req *Request
}

// status code >= 400.
func (r Response) IsError() bool {
	return isHttpError(r.StatusCode)
}

// status code >= 200 and <= 299.
func (r Response) IsSuccess() bool {
	return isHttpSuccess(r.StatusCode)
}

// Object from Request.SetError().
func (r Response) Error() any {
	return r.req.err
}
