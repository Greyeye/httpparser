package httpparser

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"golang.org/x/net/context/ctxhttp"
	"io/ioutil"
	"net/http"
	"time"
)

// initialiser must define CtxClient...
// parser := &httpparser.HttpParser{
//			Client: httpClient,
//			CtxClient: ctxhttp.Do,
// }

// overriding ctxhttp package's Do call with "CtxClient"
// https://github.com/golang/net/blob/master/context/ctxhttp/ctxhttp.go#L23
// Do(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
type HttpParser struct {
	client  *http.Client
	Do      func(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error)
	timeout *time.Duration
}

type HttpParseriface interface {
	JSONParse(ctx context.Context, req *http.Request) (*map[string]interface{}, error)
	HTTPGet(ctx context.Context, req *http.Request) (result []byte, err error)
	NewHttpParser(ctx context.Context, timeout time.Duration, req *http.Request)
}

func (h *HttpParser) JSONParse(ctx context.Context, req *http.Request) (*map[string]interface{}, error) {
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

func (h *HttpParser) HTTPGet(ctx context.Context, req *http.Request) (result []byte, err error) {

	ctx, cancel := context.WithTimeout(ctx, *h.timeout)
	defer cancel()
	// if req is nil, http.DefaultClient is used.
	// using HttpParser.Do()
	res, getErr := h.Do(ctx, h.client, req)
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

func NewHttpParser(client *http.Client, timeout *time.Duration) *HttpParser {
	var t time.Duration
	var c *http.Client
	if timeout == nil {
		t = 30 * time.Second
	}

	if client == nil {
		c = http.DefaultClient
	}
	hp := &HttpParser{c, ctxhttp.Do, &t}
	return hp
}
