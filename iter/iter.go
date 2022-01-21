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
	if !next.Present {
		return catlib.None[T]()
	}

	iter.taken += 1

	if iter.taken <= iter.cap {
		return next
	}

	return catlib.None[T]()
}

type DropIter[T any] struct {
	iter    Iterator[T]
	cap     int
	dropped bool
}

func Drop[T any](iter Iterator[T], cap int) *DropIter[T] {
	return &DropIter[T]{iter, cap, false}
}

func (iter *DropIter[T]) Next() catlib.Option[T] {
	if iter.dropped {
		return iter.iter.Next()
	}

	iter.dropped = true

	for i := 0; i < iter.cap; i++ {
		next := iter.iter.Next()
		if !next.Present {
			return catlib.None[T]()
		}
	}

	return iter.iter.Next()
}
