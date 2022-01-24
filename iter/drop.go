package iter

import . "github.com/BooleanCat/catlib/types"

type DropIter[T any] struct {
	iter    Iterator[T]
	cap     int
	dropped bool
}

func Drop[T any](iter Iterator[T], cap int) *DropIter[T] {
	return &DropIter[T]{iter, cap, false}
}

func (iter *DropIter[T]) Next() Option[T] {
	if iter.dropped {
		return iter.iter.Next()
	}

	iter.dropped = true

	for i := 0; i < iter.cap; i++ {
		next := iter.iter.Next()
		if next.IsNone() {
			return next
		}
	}

	return iter.iter.Next()
}

var _ Iterator[struct{}] = new(DropIter[struct{}])
