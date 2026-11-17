package fortgo

import (
	"fmt"
	"net/http"
)

type authenticatedTransport struct {
	base http.RoundTripper
	src  TokenSource
}

func (t *authenticatedTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if r.Header.Get("Authorization") != "" {
		return t.base.RoundTrip(r)
	}

	token, err := t.src.Token(r.Context())
	if err != nil {
		return nil, err
	}

	r = r.Clone(r.Context()) // copy before applying so original req object doesnt get the header
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	return t.base.RoundTrip(r)
}
