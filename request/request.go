package request

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DefaultClient ...
var DefaultClient = new(http.Client)

// FileField for upload file
type FileField struct {
	FieldName string
	FileName  string
	File      io.Reader
}

// Args for requests args
type Args struct {
	Client  *http.Client
	Headers map[string]string
	Params  map[string]string
	// Data is map[string]string or map[string][]string
	Data    interface{}
	Json    interface{}
	Files   []FileField
	Body    io.Reader
	Cookies map[string]string
}

// Request is alias of Args
type Request struct {
	*Args
}

// NewArgs return *Args
func NewArgs(client *http.Client) *Args {
	if client == nil {
		client = DefaultClient
	}
	// set default headers
	headers := map[string]string{}
	for k, v := range DefaultHeaders {
		headers[k] = v
	}
	return &Args{
		Client:  client,
		Headers: headers,
		Params:  nil,
		Data:    nil,
		Json:    nil,
		Body:    nil,
		Cookies: nil,
	}
}

// NewRequest return *Request
func NewRequest(client *http.Client) *Request {
	return &Request{NewArgs(client)}
}

func newURL(u string, params map[string]string) string {
	if params == nil {
		return u
	}

	p := url.Values{}
	for k, v := range params {
		p.Set(k, v)
	}
	if strings.Contains(u, "?") {
		return u + "&" + p.Encode()
	}
	return u + "?" + p.Encode()
}

func newBody(a *Args) (body io.Reader, contentType string, err error) {
	if a.Body != nil {
		return a.Body, "", nil
	}

	if a.Data == nil && a.Files == nil && a.Json == nil {
		return nil, "", nil
	}
	if a.Files != nil {
		return newMultipartBody(a, nil)
	} else if a.Json != nil {
		return newJsonBody(a)
	}
	return newFormBody(a)
}

func buildHTTPRequest(method string, url string, a *Args) (req *http.Request, err error) {
	body, contentType, err := newBody(a)
	if err != nil {
		return nil, err
	}

	u := newURL(url, a.Params)
	req, err = http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}

	applyHeaders(a, req, contentType)
	applyCookies(a, req)
	return
}

func newRequest(method string, url string, a *Args) (resp *Response, err error) {
	req, err := buildHTTPRequest(method, url, a)
	if err != nil {
		return nil, err
	}

	s, err := a.Client.Do(req)

	resp = &Response{s, nil}
	return
}

// Get method
func Get(url string, a *Args) (*Response, error) {
	resp, err := newRequest("GET", url, a)
	return resp, err
}

// Get of Request
func (req *Request) Get(url string) (*Response, error) {
	resp, err := Get(url, req.Args)
	return resp, err
}

// Post method
func Post(url string, a *Args) (*Response, error) {
	resp, err := newRequest("POST", url, a)
	return resp, err
}

// Post of Request
func (req *Request) Post(url string) (*Response, error) {
	resp, err := Post(url, req.Args)
	return resp, err
}

// PostFrom method

// Put method
func Put(url string, a *Args) (*Response, error) {
	resp, err := newRequest("PUT", url, a)
	return resp, err
}

// Delete method
func Delete(url string, a *Args) (*Response, error) {
	resp, err := newRequest("DELETE", url, a)
	return resp, err
}

//Options method
func Options(url string, a *Args) (*Response, error) {
	resp, err := newRequest("OPTIONS", url, a)
	return resp, err
}
