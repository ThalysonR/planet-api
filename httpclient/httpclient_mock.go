package httpclient

import "io"

type HTTPClientMock struct {
	HTTPRequestMock func(method string, URL string, path string, body io.Reader) ([]byte, error)
}

func (hm *HTTPClientMock) HTTPRequest(method string, URL string, path string, body io.Reader) ([]byte, error) {
	return hm.HTTPRequestMock(method, URL, path, body)
}
