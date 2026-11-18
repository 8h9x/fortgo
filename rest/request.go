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

// func test() {
// 	req, err := NewRequest(
// 		http.MethodPost,
// 		consts.FortniteMCPService,
// 		fmt.Sprintf("fortnite/api/game/v2/profile/%s/client/%s?profileId=%s&rvn=-1", "accountID", "operation", "profileID"),
// 		WithTextBody(""),
// 		WithHeader("Content-Type", "application/json"), // TODO: fix, this is yucky
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	var t account.Token
// 	err = DecodeResponse(res, &t)
// 	if err != nil {
// 		panic(err)
// 	}
// }
