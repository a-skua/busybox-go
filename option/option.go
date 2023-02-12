package option

type Option[T any] struct {
	Value T
	Valid bool
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		Value: value,
		Valid: true,
	}
}

func None[T any]() Option[T] {
	return Option[T]{}
}
