package request

import "net/http"

type Response[T any] struct {
	*http.Response
	Body T
}

type Error[T any] struct {
	StatusCode int
	Message    string
	Raw        T
	Err        error
}

func (e Error[T]) Error() string {
	return e.Message
}

func (e Error[T]) Unwrap() error {
	return e.Err
}

type EpicErrorResponse struct {
	ErrorCode          string `json:"errorCode"`
	ErrorMessage       string `json:"errorMessage"`
	NumericErrorCode   int    `json:"numericErrorCode"`
	OriginatingService string `json:"originatingService"`
	Intent             string `json:"intent"`
}
