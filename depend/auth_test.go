package depend

import (
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
	cid, token := CorpAuth()
	assert.Equal(t, want.cid, cid)
	assert.NotEqual(t, want.token, token)
}

func TestSuperToken(t *testing.T) {
	token := SuperAuth()
	assert.NotEqual(t, map[string]string{}, token)
}
