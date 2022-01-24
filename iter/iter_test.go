package iter_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/BooleanCat/catlib"
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

func TestLines(t *testing.T) {
	file, err := os.Open("fixtures/lines.txt")
	assert.Nil(t, err)
	defer file.Close()

	lines := iter.Collect[catlib.Result[[]byte]](iter.Lines(file))

	assert.Equal(t, len(lines), 5)
	assert.DeepEqual(t, lines[0].Unwrap(), []byte("This is"))
	assert.DeepEqual(t, lines[1].Unwrap(), []byte("a file"))
	assert.DeepEqual(t, lines[2].Unwrap(), []byte("with"))
	assert.DeepEqual(t, lines[3].Unwrap(), []byte("a trailing newline"))
	assert.DeepEqual(t, lines[4].Unwrap(), []byte(""))
}

func TestLinesEmpty(t *testing.T) {
	lines := iter.Collect[catlib.Result[[]byte]](iter.Lines(new(bytes.Buffer)))

	assert.Equal(t, len(lines), 1)
	assert.DeepEqual(t, lines[0].Unwrap(), []byte(""))
}
