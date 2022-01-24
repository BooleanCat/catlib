package types_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/catlib/internal/assert"
	"github.com/BooleanCat/catlib/types"
)

func TestSomePrint(t *testing.T) {
	assert.Equal(t, fmt.Sprint(types.Some("foo")), "Some(foo)")
	assert.Equal(t, fmt.Sprint(types.Some(42)), "Some(42)")
}

func TestNonePrint(t *testing.T) {
	assert.Equal(t, fmt.Sprint(types.None[string]()), "None")
	assert.Equal(t, fmt.Sprint(types.None[int]()), "None")
}

func TestSomeUnwrap(t *testing.T) {
	assert.Equal(t, types.Some("foo").Unwrap(), "foo")
}

func TestNoneUnwrap(t *testing.T) {
	defer func() { recover() }()
	types.None[struct{}]().Unwrap()
	t.Error("did not panic")
}

func TestUnwrapOrSome(t *testing.T) {
	assert.Equal(t, types.Some(21).UnwrapOr(42), 21)
}

func TestUnwrapOrNone(t *testing.T) {
	assert.Equal(t, types.None[int]().UnwrapOr(42), 42)
}

func TestUnwrapOrElseSome(t *testing.T) {
	assert.Equal(t, types.Some(21).UnwrapOrElse(func() int {
		return 42
	}), 21)
}

func TestUnwrapOrElseNone(t *testing.T) {
	assert.Equal(t, types.None[int]().UnwrapOrElse(func() int {
		return 42
	}), 42)
}

func TestUnwrapOrZeroSome(t *testing.T) {
	assert.Equal(t, types.Some(21).UnwrapOrZero(), 21)
}

func TestUnwrapOrZeroNone(t *testing.T) {
	assert.Equal(t, types.None[int]().UnwrapOrZero(), 0)
}

func TestSomeValue(t *testing.T) {
	value, present := types.Some(42).Value()
	assert.Equal(t, value, 42)
	assert.True(t, present)
}

func TestNoneValue(t *testing.T) {
	_, present := types.None[int]().Value()
	assert.False(t, present)
}

func TestIsSome(t *testing.T) {
	assert.True(t, types.Some(42).IsSome())
	assert.False(t, types.None[int]().IsSome())
}

func TestIsNone(t *testing.T) {
	assert.False(t, types.Some(42).IsNone())
	assert.True(t, types.None[int]().IsNone())
}
