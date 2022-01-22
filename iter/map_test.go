package iter_test

import (
	"strconv"
	"testing"

	"github.com/BooleanCat/catlib/internal/assert"
	"github.com/BooleanCat/catlib/iter"
)

func TestMap(t *testing.T) {
	mapped := iter.Map[int](iter.Count(), strconv.Itoa)
	assert.Equal(t, mapped.Next().Unwrap(), "0")
	assert.Equal(t, mapped.Next().Unwrap(), "1")
}

func TestMapZero(t *testing.T) {
	mapped := iter.Map[int](iter.Take[int](iter.Count(), 0), strconv.Itoa)
	assert.True(t, mapped.Next().IsNone())
}
