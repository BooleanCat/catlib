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

func TestLift(t *testing.T) {
	slicesIter := iter.Lift([]string{"foo", "bar"})
	assert.Equal(t, slicesIter.Next().Unwrap(), "foo")
	assert.Equal(t, slicesIter.Next().Unwrap(), "bar")
	assert.True(t, slicesIter.Next().IsNone())
}

func TestLiftEmpty(t *testing.T) {
	assert.True(t, iter.Lift([]string{}).Next().IsNone())
}

func TestCollect(t *testing.T) {
	numbers := iter.Collect[int](iter.Take[int](iter.Count(), 3))
	assert.DeepEqual(t, numbers, []int{0, 1, 2})
}

func TestEmpty(t *testing.T) {
	numbers := iter.Collect[int](iter.Take[int](iter.Count(), 0))
	assert.DeepEqual(t, numbers, []int{})
}
