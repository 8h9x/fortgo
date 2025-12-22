package fortgo

//import (
//	"github.com/8h9x/fortgo/request"
//	"net/http"
//)

//func (c *Client) Request(method string, serviceURL string, endpoint string, header http.Header, body string) (*http.Response, error) {
//	for key, value := range c.Header {
//		header.Set(key, value[0])
//	}
//
//	req, err := request.MakeRequest(method, serviceURL, endpoint, request.WithHeaders(header), )
//
//	return request.Request(c.HTTPClient, method, url, header, body)
//}