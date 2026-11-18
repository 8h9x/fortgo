package rest

import (
	"io"
	"net/url"
)

type Option func(*requestConfig) error

type requestConfig struct {
	body    io.Reader
	headers map[string]string
	query   url.Values
}
