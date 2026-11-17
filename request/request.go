package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func New(method, baseURL, pathname string, opts ...Option) (*http.Request, error) {
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

func DecodeResponse[T any](resp *http.Response) (Response[T], error) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response[T]{}, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var epicError EpicErrorResponse
		if err := json.Unmarshal(body, &epicError); err == nil && epicError.ErrorCode != "" {
			return Response[T]{
				StatusCode: resp.StatusCode,
				Headers:    resp.Header,
			}, Error{
				Message:   epicError.ErrorMessage,
				ErrorCode: epicError.ErrorCode,
				Raw:       epicError,
				Err:       fmt.Errorf("epic games error: %s", epicError.ErrorCode),
			}
		}

		return Response[T]{
			StatusCode: resp.StatusCode,
			Headers:    resp.Header,
		}, Error{
			Message: fmt.Sprintf("request failed with status %d: %s", resp.StatusCode, string(body)),
			Raw:     string(body),
			Err:     fmt.Errorf("http error %d", resp.StatusCode),
		}
	}

	var data T
	if len(body) > 0 {
		if err := json.Unmarshal(body, &data); err != nil {
			return Response[T]{
				StatusCode: resp.StatusCode,
				Headers:    resp.Header,
			}, Error{
				Message: "failed to decode response body",
				Raw:     string(body),
				Err:     err,
			}
		}
	}

	return Response[T]{
		Data:       data,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
	}, nil
}
