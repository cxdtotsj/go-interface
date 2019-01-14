package building

import (
	"go-interface/config"
	"go-interface/depend"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _, corpAuth, _ = depend.CorpAuth()
var superAuth, _ = depend.SuperAuth()
var commonAuth, _ = depend.CommonAuth()
var otherAuth, _ = depend.OtherAuth()

func TestCreateBuilding(t *testing.T) {
	var tests = []struct {
		header     map[string]string
		jdata      interface{}
		statusCode int
		id         string
		code       float64
	}{
		{corpAuth, *newCreate(bName, nil, 0, 0, 0, nil), 401, "", 1400},
		// {corpAuth, newCreate(bName, nil, 0, 0, 0, nil), 400, "", 1400},
	}
	for _, test := range tests {
		resp, err := createBuilding(test.jdata, test.header)
		if err != nil {
			t.Errorf("response error: %v", err)
		}
		j, _ := resp.Json()
		defer resp.Body.Close()

		assert.Equal(t, test.statusCode, resp.StatusCode, config.AssertInfo(test.jdata, resp))
		// 正常流程
		if test.statusCode == 200 {
			assert.NotEqual(t, 0, j.Get("id").Interface().(string), config.ErrInfo(resp))
			// 异常流程
		} else {
			code, _ := j.Get("code").Float64()
			assert.Equal(t, test.code, code, config.ErrInfo(resp))
		}
	}
}
