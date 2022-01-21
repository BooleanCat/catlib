package catlib_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/BooleanCat/catlib"
	"github.com/BooleanCat/catlib/internal/assert"
)

func TestOkPrint(t *testing.T) {
	o := catlib.Ok("foo")
	p := catlib.Ok(42)
	assert.Equal(t, fmt.Sprint(o), "Ok(foo)")
	assert.Equal(t, fmt.Sprint(p), "Ok(42)")
}

func TestErrPrint(t *testing.T) {
	o := catlib.Err[string](errors.New("foo"))
	p := catlib.Err[int](errors.New("blah"))
	assert.Equal(t, fmt.Sprint(o), "Err(foo)")
	assert.Equal(t, fmt.Sprint(p), "Err(blah)")
}

func TestOkUnwrap(t *testing.T) {
	o := catlib.Ok("foo")
	assert.Equal(t, o.Unwrap(), "foo")
}

func TestErrUnwrap(t *testing.T) {
	o := catlib.Err[string](errors.New("foo"))
	defer func() { recover() }()
	o.Unwrap()

	t.Error("did not panic")
}
