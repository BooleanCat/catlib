package iter

import (
	"bufio"
	"fmt"
	"io"

	"github.com/BooleanCat/catlib"
)

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

func Collect[T any](iter Iterator[T]) []T {
	items := make([]T, 0)

	for {
		if v, ok := iter.Next().Value(); ok {
			items = append(items, v)
		} else {
			return items
		}
	}
}

type LineIter struct {
	r        *bufio.Reader
	finished bool
}

func Lines(r io.Reader) *LineIter {
	return &LineIter{bufio.NewReader(r), false}
}

func (iter *LineIter) Next() catlib.Option[catlib.Result[[]byte]] {
	if iter.finished {
		return catlib.None[catlib.Result[[]byte]]()
	}

	content, err := iter.r.ReadBytes('\n')

	if err == io.EOF {
		iter.finished = true
		return catlib.Some(catlib.Ok(content))
	}

	if err != nil {
		iter.finished = true
		return catlib.Some(catlib.Err[[]byte](fmt.Errorf(`read line: %w`, err)))
	}

	return catlib.Some(catlib.Ok(content[:len(content)-1]))
}
