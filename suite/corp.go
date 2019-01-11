package suite

import (
	"go-interface/depend"
	"go-interface/info"
	"go-interface/request"
)

var _, corpToken, _ = depend.CorpToken()
var superToken, _ = depend.SuperToken()

func corpCreate(data, header map[string]string) (*request.Response, error) {
	url := info.URL("/corp/create")
	req := request.NewRequest(nil)
	req.Data = data
	req.Headers = header
	resp, err := req.Post(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
