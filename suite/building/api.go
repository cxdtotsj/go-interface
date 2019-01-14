package building

import (
	"go-interface/config"
	"go-interface/request"
)

func createBuilding(jData interface{}, header map[string]string) (*request.Response, error) {
	url := config.URL("/building/create")
	req := request.NewRequest(nil)
	req.Json = jData
	req.Headers = header
	resp, err := req.Post(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
