package option

import (
	"encoding/json"
	"fmt"
)

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

func (o Option[T]) String() string {
	if o.Valid {
		return fmt.Sprintf("Some(%v)", o.Value)
	}

	return "None"
}

func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.Valid {
		return json.Marshal(o.Value)
	}

	return []byte("null"), nil
}
