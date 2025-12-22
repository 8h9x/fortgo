package request

import (
	"fmt"
	"io"
	"net/http"
)

type Option func(*requestConfig) error

type requestConfig struct {
	body    io.Reader
	headers map[string]string
}

type Response[T any] struct {
	Data       T
	StatusCode int
	Headers    http.Header
}

type Error struct {
	Message    string
	ErrorCode  string
	Raw        any
	Err        error
}

func (e Error) Error() string {
	if e.ErrorCode != "" {
		return fmt.Sprintf("%s (code: %s)", e.Message, e.ErrorCode)
	}
	return e.Message
}

func (e Error) Unwrap() error {
	return e.Err
}

type EpicErrorResponse struct {
	ErrorCode          string `json:"errorCode"`
	ErrorMessage       string `json:"errorMessage"`
	NumericErrorCode   int    `json:"numericErrorCode"`
	OriginatingService string `json:"originatingService"`
	MessageVars        []string `json:"messageVars,omitempty"`
	Intent             string `json:"intent"`
	Continuation       string `json:"continuation,omitempty"`
	ContinuationURL    string `json:"continuationUrl,omitempty"`
	CorrectiveAction   string `json:"correctiveAction,omitempty"`
}
