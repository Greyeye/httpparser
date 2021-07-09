package httpparser

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPParser struct
// initialiser must define CtxClient...
// parser := &httpparser.HttpParser{
//			Client: httpClient,
//			CtxClient: ctxhttp.Do,
// }
// overriding ctxhttp package's Do call with "CtxClient"
// https://github.com/golang/net/blob/master/context/ctxhttp/ctxhttp.go#L23
// Do(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
type HTTPParser struct {
	client  *http.Client
	Do      func(req *http.Request) (*http.Response, error)
	timeout time.Duration
}

// HTTPParseriface is used to generate mock interface file under ./mock/
type HTTPParseriface interface {
	JSONParse(ctx context.Context, req *http.Request) (*map[string]interface{}, error)
	HTTPGet(ctx context.Context, req *http.Request) (result []byte, err error)
}

// JSONParse returns JSON payload
func (h *HTTPParser) JSONParse(ctx context.Context, req *http.Request) (*map[string]interface{}, error) {
	v := new(map[string]interface{})
	body, httpErr := h.HTTPGet(ctx, req)
	if httpErr != nil {
		return nil, httpErr
	}
	err := json.Unmarshal(body, &v)
	if err != nil || v == nil {
		return nil, errors.New("json parse failure")
	}
	return v, nil

}

// HTTPGet returns raw HTTP GET body from results.
func (h *HTTPParser) HTTPGet(ctx context.Context, req *http.Request) (result []byte, err error) {

	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()
	req = req.WithContext(ctx)
	res, getErr := h.Do(req)
	if getErr != nil {
		return nil, getErr
	}

	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	return body, nil

}

// NewHTTPParser initialise the client.
// If parameters are not given, it will initialise with the default values.
func NewHTTPParser(client *http.Client, timeout time.Duration) *HTTPParser {

	var c *http.Client

	if client == nil {
		c = http.DefaultClient
	} else {
		c = client
	}
	return &HTTPParser{c, c.Do, timeout}
}
