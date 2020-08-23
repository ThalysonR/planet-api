package httpclient

import "io"

// IHTTPClient é uma interface que contém métodos relacionados a um cliente http
type IHTTPClient interface {
	HTTPRequest(method string, URL string, path string, body io.Reader) ([]byte, error)
}
