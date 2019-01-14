package log

import "testing"

func TestDebugf(t *testing.T) {
	a := "测试"
	Debugf("%s", a)
}
