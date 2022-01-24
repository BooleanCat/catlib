package types_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/BooleanCat/catlib/internal/assert"
	"github.com/BooleanCat/catlib/types"
)

func TestOkPrint(t *testing.T) {
	assert.Equal(t, fmt.Sprint(types.Ok("foo")), "Ok(foo)")
	assert.Equal(t, fmt.Sprint(types.Ok(42)), "Ok(42)")
}

func TestErrPrint(t *testing.T) {
	assert.Equal(t, fmt.Sprint(types.Err[string](errors.New("foo"))), "Err(foo)")
	assert.Equal(t, fmt.Sprint(types.Err[int](errors.New("blah"))), "Err(blah)")
}

func TestOkUnwrap(t *testing.T) {
	assert.Equal(t, types.Ok("foo").Unwrap(), "foo")
}

func TestErrUnwrap(t *testing.T) {
	defer func() { recover() }()
	types.Err[string](errors.New("foo")).Unwrap()
	t.Error("did not panic")
}

func TestOkUnwrapOr(t *testing.T) {
	assert.Equal(t, types.Ok(42).UnwrapOr(21), 42)
}

func TestErrUnwrapOr(t *testing.T) {
	assert.Equal(t, types.Err[int](errors.New("foo")).UnwrapOr(42), 42)
}

func TestOkUnwrapOrElse(t *testing.T) {
	assert.Equal(t, types.Ok(42).UnwrapOrElse(func() int {
		return 21
	}), 42)
}

func TestErrUnwrapOrElse(t *testing.T) {
	assert.Equal(t, types.Err[int](errors.New("foo")).UnwrapOrElse(func() int {
		return 21
	}), 21)
}

func TestOkUnwrapOrZero(t *testing.T) {
	assert.Equal(t, types.Ok(42).UnwrapOrZero(), 42)
}

func TestErrUnwrapOrZero(t *testing.T) {
	assert.Equal(t, types.Err[int](errors.New("foo")).UnwrapOrZero(), 0)
}

func TestOkValue(t *testing.T) {
	value, err := types.Ok(42).Value()
	assert.Equal(t, value, 42)
	assert.Nil(t, err)
}

func TestErrValue(t *testing.T) {
	_, err := types.Err[int](errors.New("foo")).Value()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "foo")
}

func TestIsOk(t *testing.T) {
	assert.True(t, types.Ok(42).IsOk())
	assert.False(t, types.Err[int](errors.New("foo")).IsOk())
}

func TestIsErr(t *testing.T) {
	assert.False(t, types.Ok(42).IsErr())
	assert.True(t, types.Err[int](errors.New("foo")).IsErr())
}

func TestExpectOk(t *testing.T) {
	assert.Equal(t, types.Ok(42).Expect("oops"), 42)
}

func TestExpectErr(t *testing.T) {
	defer func() { assert.Equal(t, fmt.Sprint(recover()), "oops") }()
	types.Err[int](errors.New("foo")).Expect("oops")
	t.Error("did not panic")
}
