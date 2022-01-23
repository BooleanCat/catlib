package catlib_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/BooleanCat/catlib"
	"github.com/BooleanCat/catlib/internal/assert"
)

func TestOkPrint(t *testing.T) {
	assert.Equal(t, fmt.Sprint(catlib.Ok("foo")), "Ok(foo)")
	assert.Equal(t, fmt.Sprint(catlib.Ok(42)), "Ok(42)")
}

func TestErrPrint(t *testing.T) {
	assert.Equal(t, fmt.Sprint(catlib.Err[string](errors.New("foo"))), "Err(foo)")
	assert.Equal(t, fmt.Sprint(catlib.Err[int](errors.New("blah"))), "Err(blah)")
}

func TestOkUnwrap(t *testing.T) {
	assert.Equal(t, catlib.Ok("foo").Unwrap(), "foo")
}

func TestErrUnwrap(t *testing.T) {
	defer func() { recover() }()
	catlib.Err[string](errors.New("foo")).Unwrap()
	t.Error("did not panic")
}

func TestOkUnwrapOr(t *testing.T) {
	assert.Equal(t, catlib.Ok(42).UnwrapOr(21), 42)
}

func TestErrUnwrapOr(t *testing.T) {
	assert.Equal(t, catlib.Err[int](errors.New("foo")).UnwrapOr(42), 42)
}

func TestOkUnwrapOrElse(t *testing.T) {
	assert.Equal(t, catlib.Ok(42).UnwrapOrElse(func() int {
		return 21
	}), 42)
}

func TestErrUnwrapOrElse(t *testing.T) {
	assert.Equal(t, catlib.Err[int](errors.New("foo")).UnwrapOrElse(func() int {
		return 21
	}), 21)
}

func TestOkUnwrapOrZero(t *testing.T) {
	assert.Equal(t, catlib.Ok(42).UnwrapOrZero(), 42)
}

func TestErrUnwrapOrZero(t *testing.T) {
	assert.Equal(t, catlib.Err[int](errors.New("foo")).UnwrapOrZero(), 0)
}

func TestOkValue(t *testing.T) {
	value, err := catlib.Ok(42).Value()
	assert.Equal(t, value, 42)
	assert.Nil(t, err)
}

func TestErrValue(t *testing.T) {
	_, err := catlib.Err[int](errors.New("foo")).Value()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "foo")
}

func TestIsOk(t *testing.T) {
	assert.True(t, catlib.Ok(42).IsOk())
	assert.False(t, catlib.Err[int](errors.New("foo")).IsOk())
}

func TestIsErr(t *testing.T) {
	assert.False(t, catlib.Ok(42).IsErr())
	assert.True(t, catlib.Err[int](errors.New("foo")).IsErr())
}

func TestExpectOk(t *testing.T) {
	assert.Equal(t, catlib.Ok(42).Expect("oops"), 42)
}

func TestExpectErr(t *testing.T) {
	defer func() { assert.Equal(t, fmt.Sprint(recover()), "oops") }()
	catlib.Err[int](errors.New("foo")).Expect("oops")
	t.Error("did not panic")
}
