package str

import "testing"

func TestFixLen(t *testing.T) {
	var ss = map[string]int{
		"":                    10,
		"hello":               10,
		"hello world":         10,
		"你好world":             5,
		"hello哇，饭已OK，下来mixi吧": 9,
	}
	for s, n := range ss {
		fixStr, fixed := FixLenWithDesc(s, n)
		t.Logf("%q[:%d] = %q %v", s, n, fixStr, fixed)
	}
}
