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

var _ Iterator[struct{}] = new(TakeIter[struct{}])

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

var _ Iterator[struct{}] = new(DropIter[struct{}])

type MapIter[T, S any] struct {
	iter Iterator[T]
	f    func(T) S
}

func Map[T, S any](iter Iterator[T], f func(T) S) *MapIter[T, S] {
	return &MapIter[T, S]{iter, f}
}

func (iter *MapIter[T, S]) Next() catlib.Option[S] {
	next := iter.iter.Next()
	if !next.Present {
		return catlib.None[S]()
	}

	return catlib.Some(iter.f(next.Unwrap()))
}

var _ Iterator[struct{}] = new(MapIter[struct{}, struct{}])

func Fold[S, T any](iter Iterator[S], initial T, f func(S, T) T) T {
	for {
		next := iter.Next()
		if !next.Present {
			return initial
		}

		initial = f(next.Unwrap(), initial)
	}
}
