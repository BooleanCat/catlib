package assert

import "testing"

type Ordered interface {
	~string | ~int
}

func Equal[S Ordered](t *testing.T, a, b S) {
	if a != b {
		t.Errorf(`expected "%v" to equal "%v"`, a, b)
	}
}

func True(t *testing.T, b bool) {
	if !b {
		t.Error(`expected "false" to be "true"`)
	}
}

func False(t *testing.T, b bool) {
	if b {
		t.Error(`expected "true" to be "false"`)
	}
}

func Nil(t *testing.T, v interface{}) {
	if v != nil {
		t.Errorf(`expected "%v" to be nil`, v)
	}
}

func NotNil(t *testing.T, v interface{}) {
	if v == nil {
		t.Errorf(`expected "%v" not to be nil`, v)
	}
}
