package catlib_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/catlib"
	"github.com/BooleanCat/catlib/internal/assert"
)

func TestSomePrint(t *testing.T) {
	o := catlib.Some("foo")
	p := catlib.Some(42)
	assert.Equal(t, fmt.Sprint(o), "Some(foo)")
	assert.Equal(t, fmt.Sprint(p), "Some(42)")
}

func TestNonePrint(t *testing.T) {
	o := catlib.None[string]()
	p := catlib.None[int]()
	assert.Equal(t, fmt.Sprint(o), "None")
	assert.Equal(t, fmt.Sprint(p), "None")
}

func TestSomeUnwrap(t *testing.T) {
	o := catlib.Some("foo")
	assert.Equal(t, o.Unwrap(), "foo")
}

func TestNoneUnwrap(t *testing.T) {
	o := catlib.None[string]()
	defer func() { recover() }()
	o.Unwrap()

	t.Error("did not panic")
}
