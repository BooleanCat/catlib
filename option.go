package catlib

import "fmt"

type Option[T any] struct {
	t       T
	Present bool
}

func Some[T any](t T) Option[T] {
	return Option[T]{t: t, Present: true}
}

func None[T any]() Option[T] {
	return Option[T]{Present: false}
}

func (t Option[T]) String() string {
	if !t.Present {
		return "None"
	}

	return fmt.Sprintf("Some(%v)", t.t)
}

var _ fmt.Stringer = Option[struct{}]{}

func (t Option[T]) Unwrap() T {
	if t.Present {
		return t.t
	}

	panic(fmt.Sprintf(`unwrap "%s"`, t))
}
