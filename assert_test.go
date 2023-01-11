package strcase

import (
	"fmt"
	"testing"
)

type fakeT struct {
	fail bool
	log  string
}

func (t *fakeT) Fail() {
	t.fail = true
}
func (t *fakeT) Logf(format string, args ...interface{}) {
	t.log = fmt.Sprintf(format, args...)
}

func TestAssertTrue(t *testing.T) {
	{
		f := &fakeT{}
		assertTrue(f, true)
		if f.fail == true {
			t.Fail()
		}
	}
	{
		f := &fakeT{}
		assertTrue(f, false)
		if f.fail != true {
			t.Fail()
		}
	}
}

func TestAssertEqual(t *testing.T) {
	{
		f := &fakeT{}
		assertEqual(f, "foo", "foo")
		if f.fail == true {
			t.Fail()
		}
	}
	{
		f := &fakeT{}
		assertEqual(f, "foo", "bar")
		if f.fail != true {
			t.Fail()
		}
		if f.log != "Expected: foo Actual: bar" {
			t.Fail()
		}
	}
}
