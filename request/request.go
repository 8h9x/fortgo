package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func ResponseParser[T any](resp *http.Response) (Response[T], error) {
	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		epicError := EpicErrorResponse{}

		err := json.NewDecoder(resp.Body).Decode(&epicError)
		if err != nil {
			return Response[T]{}, Error[T]{
				StatusCode: resp.StatusCode,
				Message:    "failed to decode response error body",
				Err:        err,
			}
		}

		return Response[T]{}, Error[EpicErrorResponse]{
			StatusCode: resp.StatusCode,
			Message:    epicError.ErrorMessage,
			Err:        errors.New(epicError.ErrorMessage),
		}
	}

	var body T
	err := json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return Response[T]{}, Error[T]{
			StatusCode: resp.StatusCode,
			Message:    "failed to decode response body",
			Err:        err,
		}
	}

	return Response[T]{
		Response: resp,
		Body:     body,
	}, nil
}

func Request(httpClient *http.Client, method string, url string, header http.Header, body string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	for key, value := range header {
		req.Header.Set(key, value[0])
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		if resp.Body != nil {
			defer resp.Body.Close()

			var res EpicErrorResponse
			err = json.NewDecoder(resp.Body).Decode(&res)
			if err != nil {
				return nil, err
			}

			if res.ErrorMessage != "" {
				return nil, &Error[EpicErrorResponse]{
					StatusCode: resp.StatusCode,
					Message:    fmt.Sprintf("%s request to %s failed with error message: %s", method, url, res.ErrorMessage),
					Raw:        res,
				}
			}
		}

		return nil, fmt.Errorf("%s request to %s failed with status code %d", method, url, resp.StatusCode)
	}

	return resp, nil
}