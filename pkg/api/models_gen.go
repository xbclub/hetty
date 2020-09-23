// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package api

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type HTTPHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type HTTPRequestLog struct {
	ID        string           `json:"id"`
	URL       string           `json:"url"`
	Method    HTTPMethod       `json:"method"`
	Proto     string           `json:"proto"`
	Headers   []HTTPHeader     `json:"headers"`
	Body      *string          `json:"body"`
	Timestamp time.Time        `json:"timestamp"`
	Response  *HTTPResponseLog `json:"response"`
}

type HTTPResponseLog struct {
	RequestID  string       `json:"requestId"`
	Proto      string       `json:"proto"`
	Status     string       `json:"status"`
	StatusCode int          `json:"statusCode"`
	Body       *string      `json:"body"`
	Headers    []HTTPHeader `json:"headers"`
}

type HTTPMethod string

const (
	HTTPMethodGet     HTTPMethod = "GET"
	HTTPMethodHead    HTTPMethod = "HEAD"
	HTTPMethodPost    HTTPMethod = "POST"
	HTTPMethodPut     HTTPMethod = "PUT"
	HTTPMethodDelete  HTTPMethod = "DELETE"
	HTTPMethodConnect HTTPMethod = "CONNECT"
	HTTPMethodOptions HTTPMethod = "OPTIONS"
	HTTPMethodTrace   HTTPMethod = "TRACE"
	HTTPMethodPatch   HTTPMethod = "PATCH"
)

var AllHTTPMethod = []HTTPMethod{
	HTTPMethodGet,
	HTTPMethodHead,
	HTTPMethodPost,
	HTTPMethodPut,
	HTTPMethodDelete,
	HTTPMethodConnect,
	HTTPMethodOptions,
	HTTPMethodTrace,
	HTTPMethodPatch,
}

func (e HTTPMethod) IsValid() bool {
	switch e {
	case HTTPMethodGet, HTTPMethodHead, HTTPMethodPost, HTTPMethodPut, HTTPMethodDelete, HTTPMethodConnect, HTTPMethodOptions, HTTPMethodTrace, HTTPMethodPatch:
		return true
	}
	return false
}

func (e HTTPMethod) String() string {
	return string(e)
}

func (e *HTTPMethod) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HTTPMethod(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HttpMethod", str)
	}
	return nil
}

func (e HTTPMethod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
