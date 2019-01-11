package request

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
	"strings"
)

// newMultipartBody for file upload
func newMultipartBody(a *Args, data url.Values) (body io.Reader, contentType string, err error) {
	files := a.Files
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//set files
	for _, file := range files {
		fileWriter, err := bodyWriter.CreateFormFile(file.FieldName, file.FileName)
		if err != nil {
			return nil, "", err
		}
		// iocppy
		_, err = io.Copy(fileWriter, file.File)
		if err != nil {
			return nil, "", err
		}
	}

	// data in Args
	if a.Data != nil {
		switch a.Data.(type) {
		case map[string]string:
			for k, v := range a.Data.(map[string]string) {
				bodyWriter.WriteField(k, v)
			}
		case map[string][]string:
			for k, v := range a.Data.(map[string][]string) {
				for n := range v {
					bodyWriter.WriteField(k, v[n])
				}
			}
		}
	}
	contentType = bodyWriter.FormDataContentType()
	defer bodyWriter.Close()

	body = bodyBuf
	return body, contentType, nil
}

// newJsonBody for the request of PostJson
func newJsonBody(a *Args) (body io.Reader, contentType string, err error) {
	b, err := json.Marshal(a.Json)
	if err != nil {
		return nil, "", err
	}
	return bytes.NewBuffer(b), DefaultJSONType, nil
}

// data can be map[string]string or map[string][]string
func newFormBody(a *Args) (body io.Reader, contentType string, err error) {
	vs := url.Values{}
	switch a.Data.(type) {
	case map[string]string:
		for k, v := range a.Data.(map[string]string) {
			vs.Set(k, v)
		}
	case map[string][]string:
		for k, v := range a.Data.(map[string][]string) {
			for n := range v {
				vs.Add(k, v[n])
			}
		}
	}
	if a.Files != nil {
		return newMultipartBody(a, vs)
	}
	return strings.NewReader(vs.Encode()), DefaultContentType, nil
}
