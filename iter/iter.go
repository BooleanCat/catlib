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
	return new(CountIter)
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

type SliceIter[T any] struct {
	slice []T
	index int
}

func Lift[T any](slice []T) *SliceIter[T] {
	return &SliceIter[T]{slice: slice}
}

func (iter *SliceIter[T]) Next() catlib.Option[T] {
	if iter.index == len(iter.slice) {
		return catlib.None[T]()
	}

	iter.index += 1
	return catlib.Some(iter.slice[iter.index-1])
}

var _ Iterator[struct{}] = &SliceIter[struct{}]{}
