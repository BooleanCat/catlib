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
