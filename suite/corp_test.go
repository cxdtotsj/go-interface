package suite

import (
	"fmt"
	"go-interface/common"
	"go-interface/tools"
	"testing"

	"github.com/stretchr/testify/assert"
)

var st, _ = common.SuperToken()
var cid, ct, _ = common.CorpToken()

func TestCorpCreate(t *testing.T) {
	var tests = []struct {
		header     map[string]string
		data       map[string]string
		statusCode int
		id         string
		code       float64
	}{
		{common.Auth(st), map[string]string{"name": "随机测试" + tools.RandInt()}, 200, "", 0},
		{common.Auth(ct), map[string]string{"name": "gotest" + tools.RandInt()}, 403, "", 1402},
	}
	for _, test := range tests {
		resp, err := corpCreate(test.data, test.header)
		j, _ := resp.Json()
		text, _ := resp.Text()
		defer resp.Body.Close()
		fmt.Println("tezt:", text)
		if err != nil {
			t.Errorf("response error: %v", err)
		}
		assert.Equal(t, test.statusCode, resp.StatusCode, text)
		if resp.OK() {
			assert.NotEqual(t, 0, j["id"].(string), text)
		} else {
			assert.Equal(t, test.code, j["code"].(float64), text)
		}
	}
}
