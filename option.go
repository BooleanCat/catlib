package catlib

import "fmt"

// Option provides a wrapper for a value that can indicate the presence or the
// absence of a value. Option should be instantiated using the Some(value) and
// None() variants.
type Option[T any] struct {
	t       T
	present bool
}

// Some creates an option holding a value.
func Some[T any](t T) Option[T] {
	return Option[T]{t, true}
}

// None creates an option holding no value.
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

// Unwrap returns the underlying value of an Option, or panics if no value is
// present.
func (t Option[T]) Unwrap() T {
	if t.present {
		return t.t
	}

	panic(fmt.Sprintf(`unwrap "%s"`, t))
}

// UnwrapOr returns the underlying value of an Option, or the provided value if
// the Option contains no value.
func (t Option[T]) UnwrapOr(s T) T {
	if t.present {
		return t.t
	}

	return s
}

// UnwrapOrElse returns the underlying value of an Option, or the result of
// calling the provided function if the Option contains no value.
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
