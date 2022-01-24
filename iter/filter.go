package iter

import "github.com/BooleanCat/catlib/types"

type FilterIter[T any] struct {
	iter Iterator[T]
	f    func(T) bool
}

func Filter[T any](iter Iterator[T], f func(T) bool) *FilterIter[T] {
	return &FilterIter[T]{iter, f}
}

func (iter *FilterIter[T]) Next() types.Option[T] {
	for {
		next := iter.iter.Next()
		if next.IsNone() {
			return next
		}

		if iter.f(next.Unwrap()) {
			return next
		}
	}
}

var _ Iterator[struct{}] = new(FilterIter[struct{}])

func Exclude[T any](iter Iterator[T], f func(T) bool) *FilterIter[T] {
	return &FilterIter[T]{iter, func(t T) bool {
		return !f(t)
	}}
}
