package corp

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

func TestCorpCreate(t *testing.T) {
	var tests = []struct {
		header     map[string]string
		data       map[string]string
		statusCode int
		id         string
		code       float64
	}{
		{superAuth, map[string]string{"name": ""}, 400, "", 1400},         // 无名称
		{nil, map[string]string{"name": corpName("随机测试")}, 401, "", 1401}, //无token
		{superAuth, map[string]string{"name": corpName("随机测试")}, 200, "", 0},
		{corpAuth, map[string]string{"name": corpName("gotest")}, 403, "", 1403},
		{commonAuth, map[string]string{"name": corpName("随机测试")}, 403, "", 1403}, // 普通用户
		{otherAuth, map[string]string{"name": corpName("随机测试")}, 403, "", 1403},  // 其他组织管理员
	}
	for _, test := range tests {
		resp, err := corpCreate(test.data, test.header)
		if err != nil {
			t.Errorf("response error: %v", err)
		}
		j, _ := resp.Json()
		defer resp.Body.Close()

		assert.Equal(t, test.statusCode, resp.StatusCode, config.AssertInfo(test.data, resp))
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
