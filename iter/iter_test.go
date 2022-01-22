package iter_test

import (
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

func TestFold(t *testing.T) {
	add := func(a, b int) int { return a + b }
	sum := iter.Fold[int](iter.Take[int](iter.Count(), 11), 0, add)
	assert.Equal(t, sum, 55)
}
