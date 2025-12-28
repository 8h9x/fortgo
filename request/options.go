package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"net/url"
	"strings"
)

func WithBody(body io.Reader, contentType string) Option {
	return func(cfg *requestConfig) error {
		cfg.body = body
		cfg.headers["Content-Type"] = contentType
		return nil
	}
}

func WithHeader(key, value string) Option {
	return func(cfg *requestConfig) error {
		cfg.headers[key] = value
		return nil
	}
}

func WithHeaders(headers map[string]string) Option {
	return func(cfg *requestConfig) error {
		maps.Copy(cfg.headers, headers)
		return nil
	}
}

func WithQueryParamaters(query url.Values) Option {
	return func(cfg *requestConfig) error {
		if cfg.query == nil {
			cfg.query = make(url.Values)
		}

		maps.Copy(cfg.query, query)

		return nil
	}
}

func WithFormBody(values url.Values) Option {
	return func(cfg *requestConfig) error {
		cfg.body = strings.NewReader(values.Encode())
		cfg.headers["Content-Type"] = "application/x-www-form-urlencoded"
		return nil
	}
}

func WithJSONBody(v any) Option {
	return func(cfg *requestConfig) error {
		data, err := json.Marshal(v)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}
		cfg.body = bytes.NewReader(data)
		cfg.headers["Content-Type"] = "application/json"
		return nil
	}
}

func WithTextBody(text string) Option {
	return func(cfg *requestConfig) error {
		cfg.body = strings.NewReader(text)
		cfg.headers["Content-Type"] = "text/plain"
		return nil
	}
}

func WithBearerToken(accessToken string) Option {
	return WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken))
}

func WithBasicToken(basicToken string) Option {
	return WithHeader("Authorization", fmt.Sprintf("Basic %s", basicToken))
}
