package restclient

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/wizk3y/go-jenkins-sdk/logger"
)

// WrapWithBasicAuthRoundTripper --
func WrapWithBasicAuthRoundTripper(r http.RoundTripper, username, password string) http.RoundTripper {
	return &basicAuthRoundTripper{
		r: r,

		username: username,
		password: password,
	}
}

type basicAuthRoundTripper struct {
	r http.RoundTripper

	username string
	password string
}

// RoundTrip -- implement for interface http.RoundTripper
func (t *basicAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.username, t.password)

	return t.r.RoundTrip(req)
}

// WrapWithLogRoundTripper --
func WrapWithLogRoundTripper(r http.RoundTripper, l logger.LoggerInterface) http.RoundTripper {
	return &logRoundTripper{
		r: r,
		l: l,
	}
}

type logRoundTripper struct {
	r http.RoundTripper
	l logger.LoggerInterface
}

// RoundTrip -- implement for interface http.RoundTripper
func (t *logRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// do normal when no logger
	if t.l == nil {
		return t.r.RoundTrip(req)
	}

	// log before do request
	url := req.URL.String()
	t.l.Infof("Start do request with url %v", url)
	if req.Body != nil {
		reqBuffer, _ := ioutil.ReadAll(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBuffer))
		t.l.Debugf("Request body: %v", string(reqBuffer))
	}

	if len(req.Header) > 0 {
		t.l.Debug("Request header:")
		for k, v := range req.Header {
			t.l.Debugf("    %v: %v", k, v)
		}
	}

	resp, err := t.r.RoundTrip(req)

	// log after do request
	if err == nil {
		if resp.Body != nil {
			respBuffer, _ := ioutil.ReadAll(resp.Body)
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(respBuffer))
			t.l.Debugf("Response body: %v", string(respBuffer))
		}

		if len(resp.Header) > 0 {
			t.l.Debug("Response header:")
			for k, v := range resp.Header {
				t.l.Debugf("    %v: %v", k, v)
			}
		}
	}
	t.l.Infof("Done do request with url %v. Status code: %v", url, resp.StatusCode)

	return resp, err
}
