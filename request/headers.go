package request

import "net/http"

// DefaultContentType define default Content-Type Header for form body
var DefaultContentType = "application/x-www-form-urlencoded; charset=utf-8"

// DefaultJSONType define default Content-Type Header for json body
var DefaultJSONType = "application/json; charset=utf-8"

// DefaultUserAgent ...
var DefaultUserAgent = "go-request"

// DefaultHeaders ...
var DefaultHeaders = map[string]string{
	"Connection": "keep-alive",
	"Accept":     "*/*",
	// "Accept-Encoding": "gzip, deflate",
	"User-Agent": DefaultUserAgent,
}

// ApplyHeaders ...
// func applyHeaders(req *http.Request, contentType string, header map[string]string) {
// 	for k, v := range DefaultHeaders {
// 		req.Header.Add(k, v)
// 	}
// 	_, ok := req.Header["Content-Type"]
// 	if !ok {
// 		req.Header.Set("Content-Type", contentType)
// 	}
// 	for k, v := range header {
// 		req.Header.Set(k, v)
// 	}
// }

// ApplyHeaders ...
func applyHeaders(a *Args, req *http.Request, contentType string) {
	// apply default headers
	for k, v := range DefaultHeaders {
		_, ok := a.Headers[k]
		if !ok {
			req.Header.Set(k, v)
		}
	}

	// apply header in the Args
	for k, v := range a.Headers {
		req.Header.Set(k, v)
	}

	// apply "Content-Type" header
	_, ok := a.Headers["Content-Type"]
	if !ok {
		if a.Headers["Content-Type"] != "" {
			req.Header.Set("Content-Type", contentType)
		}
	}

}
