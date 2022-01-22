package iter

import "github.com/BooleanCat/catlib"

type Iterator[T any] interface {
	Next() catlib.Option[T]
}

type CountIter struct {
	count int
}

func (c *CountIter) Next() catlib.Option[int] {
	next := c.count
	c.count += 1
	return catlib.Some(next)
}

func Count() *CountIter {
	return &CountIter{}
}

var _ Iterator[int] = new(CountIter)

func Fold[S, T any](iter Iterator[S], initial T, f func(S, T) T) T {
	for {
		if v, ok := iter.Next().Value(); !ok {
			return initial
		} else {
			initial = f(v, initial)
		}
	}
}
