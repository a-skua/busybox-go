package option

import (
	"bytes"
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

func (o *Option[T]) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, []byte("null")) == 0 {
		o.Valid = false
		return nil
	}

	err := json.Unmarshal(data, &o.Value)
	if err != nil {
		return err
	}

	o.Valid = true
	return nil
}
