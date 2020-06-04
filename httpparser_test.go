package httpparser

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func mockCtxhttpDo(res *http.Response, errToReturn error) func(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	return func(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
		return res, errToReturn
	}
}

func marshall(v *map[string]interface{}, t interface{}) interface{} {
	tempJSON, _ := json.Marshal(v)
	json.Unmarshal(tempJSON, t)
	return t

}

func TestHttpParser_JSONParse_normal(t *testing.T) {
	type nameTest struct {
		Name string
	}
	body := []byte(`{"name": "james"}`)

	response := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer(body)),
	}

	jsonparser := NewHTTPParser(nil, nil)
	jsonparser.Do = mockCtxhttpDo(response, nil)
	resultInterface := &nameTest{} // placeholder for the result

	request := &http.Request{} // empty request
	response1, err := jsonparser.JSONParse(context.TODO(), request)
	if err != nil {
		t.Error(err)
	}
	responseStruct := &nameTest{}
	marshall(response1, responseStruct)
	json.Unmarshal(body, &resultInterface)
	assert.Equal(t, responseStruct.Name, resultInterface.Name)
}

func TestHTTPParser_JSONParse_BadJsonFailure(t *testing.T) {
	type nameTest struct {
		Name string
	}
	body := []byte(`{"name":}`)

	response := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer(body)),
	}
	jsonparser := NewHTTPParser(nil, nil)
	jsonparser.Do = mockCtxhttpDo(response, nil)

	resultInterface := &nameTest{} // placeholder for the result

	request := &http.Request{} // empty request
	response1, err := jsonparser.JSONParse(context.TODO(), request)
	responseStruct := &nameTest{}
	marshall(response1, responseStruct)

	json.Unmarshal(body, &resultInterface)
	assert.NotNil(t, err) // expecting Json parse failure, cannot be empty
	assert.EqualError(t, err, "json parse failure")
	assert.Equal(t, responseStruct.Name, resultInterface.Name)
}

func TestHTTPParser_JSONParse_BadHTTPCallFailure(t *testing.T) {
	type nameTest struct {
		Name string
	}
	body := []byte(`{}`)

	response := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer(body)),
	}

	jsonparser := NewHTTPParser(nil, nil)
	jsonparser.Do = mockCtxhttpDo(response, errors.New("mock http failure"))
	resultInterface := &nameTest{} // placeholder for the result

	request := &http.Request{} // empty request
	response1, err := jsonparser.JSONParse(context.TODO(), request)
	responseStruct := &nameTest{}
	marshall(response1, responseStruct)

	json.Unmarshal(body, &resultInterface)
	assert.NotNil(t, err) // expecting Json parse failure, cannot be empty
	assert.EqualError(t, err, "mock http failure")
	assert.Equal(t, responseStruct.Name, resultInterface.Name)
}

func TestHTTPParser_HTTPGet_normal(t *testing.T) {

	body := []byte(`{"name": "james"}`)

	response := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer(body)),
	}

	httpparser := NewHTTPParser(nil, nil)
	httpparser.Do = mockCtxhttpDo(response, nil)
	//httpparser := &HttpParser{client: &http.Client{}, Do: mockCtxDo} //mocked http.Client
	request := &http.Request{} // empty request
	result, err := httpparser.HTTPGet(context.TODO(), request)
	assert.Nil(t, err)
	assert.Equal(t, body, result)
}

func TestHTTPParser_HTTPGet_error(t *testing.T) {

	body := []byte(`{"name": "james"}`)

	response := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer(body)),
	}

	httpparser := NewHTTPParser(nil, nil)
	httpparser.Do = mockCtxhttpDo(response, errors.New("mock error"))
	request := &http.Request{} // empty request
	result, err := httpparser.HTTPGet(context.TODO(), request)
	assert.NotNil(t, err)
	assert.NotEqual(t, body, result)
	assert.Nil(t, result)
}
