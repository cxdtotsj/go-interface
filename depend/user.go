package depend

import (
	"go-interface/info"
	"go-interface/request"

	simplejson "github.com/bitly/go-simplejson"
)

// UserLogin a method of login
func Login(email, mobile, password string) (*simplejson.Json, error) {
	api := "/user/login"
	url := info.BaseURL + api
	req := request.NewRequest(nil)
	req.Data = map[string]string{
		"email":    email,
		"mobile":   mobile,
		"password": password,
	}
	resp, _ := req.Post(url)
	defer resp.Body.Close()
	j, err := resp.Json()
	return j, err
}
