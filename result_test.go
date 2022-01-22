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
