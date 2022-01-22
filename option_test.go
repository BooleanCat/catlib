package catlib_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/catlib"
	"github.com/BooleanCat/catlib/internal/assert"
)

func TestSomePrint(t *testing.T) {
	assert.Equal(t, fmt.Sprint(catlib.Some("foo")), "Some(foo)")
	assert.Equal(t, fmt.Sprint(catlib.Some(42)), "Some(42)")
}

func TestNonePrint(t *testing.T) {
	assert.Equal(t, fmt.Sprint(catlib.None[string]()), "None")
	assert.Equal(t, fmt.Sprint(catlib.None[int]()), "None")
}

func TestSomeUnwrap(t *testing.T) {
	assert.Equal(t, catlib.Some("foo").Unwrap(), "foo")
}

func TestNoneUnwrap(t *testing.T) {
	defer func() { recover() }()
	catlib.None[struct{}]().Unwrap()
	t.Error("did not panic")
}

func TestUnwrapOrSome(t *testing.T) {
	assert.Equal(t, catlib.Some(21).UnwrapOr(42), 21)
}

func TestUnwrapOrNone(t *testing.T) {
	assert.Equal(t, catlib.None[int]().UnwrapOr(42), 42)
}

func TestUnwrapOrElseSome(t *testing.T) {
	assert.Equal(t, catlib.Some(21).UnwrapOrElse(func() int {
		return 42
	}), 21)
}

func TestUnwrapOrElseNone(t *testing.T) {
	assert.Equal(t, catlib.None[int]().UnwrapOrElse(func() int {
		return 42
	}), 42)
}

func TestUnwrapOrZeroSome(t *testing.T) {
	assert.Equal(t, catlib.Some(21).UnwrapOrZero(), 21)
}

func TestUnwrapOrZeroNone(t *testing.T) {
	assert.Equal(t, catlib.None[int]().UnwrapOrZero(), 0)
}
