package catlib

import "fmt"

type Option[T any] struct {
	t       T
	present bool
}

func Some[T any](t T) Option[T] {
	return Option[T]{t, true}
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
	if t.present {
		return t.t
	}

	panic(fmt.Sprintf(`unwrap "%s"`, t))
}

func (t Option[T]) UnwrapOr(s T) T {
	if t.present {
		return t.t
	}

	return s
}

func (t Option[T]) UnwrapOrElse(f func() T) T {
	if t.present {
		return t.t
	}

	return f()
}

func (t Option[T]) UnwrapOrZero() T {
	if t.present {
		return t.t
	}

	var s T
	return s
}

func (t Option[T]) Value() (T, bool) {
	return t.t, t.present
}

func (t Option[T]) IsSome() bool {
	return t.present
}

func (t Option[T]) IsNone() bool {
	return !t.present
}
