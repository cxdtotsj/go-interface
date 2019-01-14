package depend

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorpToken(t *testing.T) {
	var want = struct {
		cid   string
		token map[string]string
	}{
		"109777231634510875", map[string]string{},
	}
	cid, token, _ := CorpAuth()
	assert.Equal(t, want.cid, cid)
	assert.NotEqual(t, want.token, token)
}

func TestSuperToken(t *testing.T) {
	token, _ := SuperAuth()
	assert.NotEqual(t, map[string]string{}, token)
}

func TestCommonAuth(t *testing.T) {
	auth, _ := CommonAuth()
	fmt.Println(auth)
}

func TestOtherCorpAuth(t *testing.T) {
	auth, _ := OtherAuth()
	fmt.Println(auth)
}
