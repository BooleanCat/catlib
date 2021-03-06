package iter_test

import (
	"testing"

	"github.com/BooleanCat/catlib/internal/assert"
	"github.com/BooleanCat/catlib/iter"
)

func TestTake(t *testing.T) {
	capped := iter.Take[int](iter.Count(), 2)
	assert.Equal(t, capped.Next().Unwrap(), 0)
	assert.Equal(t, capped.Next().Unwrap(), 1)

	defer func() { recover() }()
	capped.Next().Unwrap()
	t.Error("did not panic")
}

func TestTakeZero(t *testing.T) {
	capped := iter.Take[int](iter.Count(), 0)

	defer func() { recover() }()
	capped.Next().Unwrap()
	t.Error("did not panic")
}
