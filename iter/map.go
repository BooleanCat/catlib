package iter

import "github.com/BooleanCat/catlib/types"

type MapIter[T, S any] struct {
	iter Iterator[T]
	f    func(T) S
}

func Map[T, S any](iter Iterator[T], f func(T) S) *MapIter[T, S] {
	return &MapIter[T, S]{iter, f}
}

func (iter *MapIter[T, S]) Next() types.Option[S] {
	if v, ok := iter.iter.Next().Value(); !ok {
		return types.None[S]()
	} else {
		return types.Some(iter.f(v))
	}
}

var _ Iterator[struct{}] = new(MapIter[struct{}, struct{}])
