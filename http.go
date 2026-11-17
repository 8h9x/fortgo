package fortgo

import (
	"net/http"
	"time"
)

func NewHTTPClient(src TokenSource) *http.Client {
	return &http.Client{
		Timeout:   30 * time.Second,
		Transport: &authenticatedTransport{src: src, base: http.DefaultTransport},
	}
}
