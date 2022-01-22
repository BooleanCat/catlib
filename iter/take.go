package iter

import "github.com/BooleanCat/catlib"

type TakeIter[T any] struct {
	iter  Iterator[T]
	cap   int
	taken int
}

func Take[T any](iter Iterator[T], cap int) *TakeIter[T] {
	return &TakeIter[T]{iter, cap, 0}
}

func (iter *TakeIter[T]) Next() catlib.Option[T] {
	next := iter.iter.Next()
	if next.IsNone() {
		return next
	}

	iter.taken += 1

	if iter.taken <= iter.cap {
		return next
	}

	return catlib.None[T]()
}

var _ Iterator[struct{}] = new(TakeIter[struct{}])
