package iter

import "github.com/BooleanCat/catlib"

type MapIter[T, S any] struct {
	iter Iterator[T]
	f    func(T) S
}

func Map[T, S any](iter Iterator[T], f func(T) S) *MapIter[T, S] {
	return &MapIter[T, S]{iter, f}
}

func (iter *MapIter[T, S]) Next() catlib.Option[S] {
	if v, ok := iter.iter.Next().Value(); !ok {
		return catlib.None[S]()
	} else {
		return catlib.Some(iter.f(v))
	}
}

var _ Iterator[struct{}] = new(MapIter[struct{}, struct{}])
