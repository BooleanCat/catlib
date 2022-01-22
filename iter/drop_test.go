package iter_test

import (
	"testing"

	"github.com/BooleanCat/catlib/internal/assert"
	"github.com/BooleanCat/catlib/iter"
)

func TestDrop(t *testing.T) {
	skipped := iter.Drop[int](iter.Count(), 5)
	assert.Equal(t, skipped.Next().Unwrap(), 5)
}

func TestDropZero(t *testing.T) {
	skipped := iter.Drop[int](iter.Count(), 0)
	assert.Equal(t, skipped.Next().Unwrap(), 0)
}
