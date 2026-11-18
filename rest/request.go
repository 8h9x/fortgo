package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

func NewRequest(method, baseURL, pathname string, opts ...Option) (*http.Request, error) {
	if baseURL == "" {
		return nil, errors.New("base url is required")
	}

	cfg := &requestConfig{headers: make(map[string]string)}
	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, fmt.Errorf("failed to apply request option: %w", err)
		}
	}

	reqURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base url %q: %w", baseURL, err)
	}
	reqURL = reqURL.JoinPath(pathname)
	if len(cfg.query) > 0 {
		reqURL.RawQuery = cfg.query.Encode()
	}

	req, err := http.NewRequest(method, reqURL.String(), cfg.body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	for k, v := range cfg.headers {
		req.Header.Set(k, v)
	}
	return req, nil
}

func DecodeResponse[T any](resp *http.Response, out *T) error {
	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(out)
	if err != nil {
		return err
	}

	return nil
}
