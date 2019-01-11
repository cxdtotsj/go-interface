package request

import (
	"compress/gzip"
	"compress/zlib"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	simplejson "github.com/bitly/go-simplejson"
)

// Response ...
type Response struct {
	*http.Response
	// multiple calls resp.body
	content []byte
}

// Json return json format
func (resp *Response) Json() (*simplejson.Json, error) {
	b, err := resp.Content()
	if err != nil {
		return nil, err
	}
	return simplejson.NewJson(b)
}

// JSON return json.Map() of map[string]interface{}
func (resp *Response) JSON() (map[string]interface{}, error) {
	b, err := resp.Content()
	if err != nil {
		return nil, err
	}
	j, err := simplejson.NewJson(b)
	if err != nil {
		return nil, err
	}
	return j.Map()
}

// Content return []byte
func (resp *Response) Content() (b []byte, err error) {
	if resp.content != nil {
		return resp.content, nil
	}
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		if reader, err = gzip.NewReader(resp.Body); err != nil {
			return nil, err
		}
	case "deflate":
		if reader, err = zlib.NewReader(resp.Body); err != nil {
			return nil, err
		}
	default:
		reader = resp.Body
	}
	defer reader.Close()

	if b, err = ioutil.ReadAll(reader); err != nil {
		return nil, err
	}
	// save resp.body to Response
	resp.content = b
	return b, err
}

// Text return string format
func (resp *Response) Text() (string, error) {
	b, err := resp.Content()
	return string(b), err
}

// OK return true if StatusCode is 200
func (resp *Response) OK() bool {
	return resp.StatusCode == 200
}

// URL return finally request url
func (resp *Response) URL() (*url.URL, error) {
	u := resp.Request.URL
	switch resp.StatusCode {
	case http.StatusMovedPermanently, http.StatusFound,
		http.StatusSeeOther, http.StatusTemporaryRedirect:
		location, err := resp.Location()
		if err != nil {
			return nil, err
		}
		u = u.ResolveReference(location)
	}
	return u, nil
}
