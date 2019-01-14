package corp

import (
	"go-interface/config"
	"go-interface/request"
)

func corpCreate(data, header map[string]string) (*request.Response, error) {
	url := config.URL("/corp/create")
	req := request.NewRequest(nil)
	req.Data = data
	req.Headers = header
	resp, err := req.Post(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
