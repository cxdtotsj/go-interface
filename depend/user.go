package depend

import (
	"go-interface/config"
	"go-interface/request"

	simplejson "github.com/bitly/go-simplejson"
)

// Login a method of login
func Login(email, mobile, password string) (*simplejson.Json, error) {
	url := config.URL("/user/login")
	req := request.NewRequest(nil)
	req.Data = map[string]string{
		"email":    email,
		"mobile":   mobile,
		"password": password,
	}
	resp, err := req.Post(url)
	if err != nil {
		return nil, err
	}
	j, err := resp.Json()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return j, nil
}
