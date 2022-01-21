package iter_test

import (
	"strconv"
	"testing"

	"github.com/BooleanCat/catlib/internal/assert"
	"github.com/BooleanCat/catlib/iter"
)

func TestCount(t *testing.T) {
	count := iter.Count()
	assert.Equal(t, count.Next().Unwrap(), 0)
	assert.Equal(t, count.Next().Unwrap(), 1)
	assert.Equal(t, count.Next().Unwrap(), 2)
}

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

func TestDrop(t *testing.T) {
	skipped := iter.Drop[int](iter.Count(), 5)
	assert.Equal(t, skipped.Next().Unwrap(), 5)
}

func TestDropZero(t *testing.T) {
	skipped := iter.Drop[int](iter.Count(), 0)
	assert.Equal(t, skipped.Next().Unwrap(), 0)
}

func TestMap(t *testing.T) {
	mapped := iter.Map[int](iter.Count(), strconv.Itoa)
	assert.Equal(t, mapped.Next().Unwrap(), "0")
	assert.Equal(t, mapped.Next().Unwrap(), "1")
}
