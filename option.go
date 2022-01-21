package catlib

import "fmt"

type Option[T any] struct {
	t       T
	present bool
}

func Some[T any](t T) Option[T] {
	return Option[T]{t: t, present: true}
}

func None[T any]() Option[T] {
	return Option[T]{present: false}
}

func (t Option[T]) String() string {
	if !t.present {
		return "None"
	}

	return fmt.Sprintf("Some(%v)", t.t)
}

var _ fmt.Stringer = Option[struct{}]{}

func (t Option[T]) Unwrap() T {
	if !t.present {
		panic(fmt.Sprintf(`unwrap "%s"`, t))
	}

	return t.t
}
