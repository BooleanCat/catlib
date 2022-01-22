package iter_test

import (
	"testing"

	"github.com/BooleanCat/catlib/internal/assert"
	"github.com/BooleanCat/catlib/iter"
)

func TestFilter(t *testing.T) {
	isEven := func(i int) bool { return i%2 == 0 }
	filtered := iter.Filter[int](iter.Count(), isEven)
	assert.Equal(t, filtered.Next().Unwrap(), 0)
	assert.Equal(t, filtered.Next().Unwrap(), 2)
	assert.Equal(t, filtered.Next().Unwrap(), 4)
}

func TestFilterZero(t *testing.T) {
	all := func(i int) bool { return true }
	filtered := iter.Filter[int](iter.Take[int](iter.Count(), 0), all)
	assert.True(t, filtered.Next().IsNone())
}

func TestExclude(t *testing.T) {
	isEven := func(i int) bool { return i%2 == 0 }
	filtered := iter.Exclude[int](iter.Count(), isEven)
	assert.Equal(t, filtered.Next().Unwrap(), 1)
	assert.Equal(t, filtered.Next().Unwrap(), 3)
	assert.Equal(t, filtered.Next().Unwrap(), 5)
}

func TestExcludeZero(t *testing.T) {
	all := func(i int) bool { return true }
	filtered := iter.Exclude[int](iter.Take[int](iter.Count(), 0), all)
	assert.True(t, filtered.Next().IsNone())
}
