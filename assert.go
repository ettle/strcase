package strcase

import "testing"

// assertTrue is a lightweight replacement for testify/assert to reduce
// dependencies
func assertTrue(t *testing.T, value bool, msg ...interface{}) {
	if !value {
		if len(msg) > 0 {
			t.Log(msg...)
		}
		t.Fail()
	}
}
