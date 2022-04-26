package restclient

import (
	"io"
	"io/ioutil"
	"net/http"
)

// HttpOptionFunc --
type HttpOptionFunc interface {
	Apply(req *http.Request)
}

// HttpBodyOptionFunc --
func HttpBodyOptionFunc(body io.Reader) HttpOptionFunc {
	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = ioutil.NopCloser(body)
	}

	return httpBodyOptionFunc{
		body: rc,
	}
}

type httpBodyOptionFunc struct {
	body io.ReadCloser
}

func (f httpBodyOptionFunc) Apply(req *http.Request) {
	req.Body = f.body
}

// HttpQueryOptionFunc --
func HttpQueryOptionFunc(key, value string) HttpOptionFunc {
	return httpQueryOptionFunc{
		key:   key,
		value: value,
	}
}

type httpQueryOptionFunc struct {
	key   string
	value string
}

func (f httpQueryOptionFunc) Apply(req *http.Request) {
	q := req.URL.Query()
	q.Set(f.key, f.value)
	req.URL.RawQuery = q.Encode()
}

// HttpHeaderOptionFunc --
func HttpHeaderOptionFunc(key, value string) HttpOptionFunc {
	return httpHeaderOptionFunc{
		key:   key,
		value: value,
	}
}

type httpHeaderOptionFunc struct {
	key   string
	value string
}

func (f httpHeaderOptionFunc) Apply(req *http.Request) {
	req.Header.Set(f.key, f.value)
}
