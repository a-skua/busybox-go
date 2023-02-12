package option

import (
	"reflect"
	"testing"
)

func TestSome(t *testing.T) {
	type test[T any] struct {
		name  string
		value T
		want  Option[T]
	}

	do := func(tt *test[int]) {
		t.Run(tt.name, func(t *testing.T) {
			got := Some(tt.value)
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test[int]{
		{
			name:  "valid int",
			value: 0,
			want: Option[int]{
				Value: 0,
				Valid: true,
			},
		},
		{
			name:  "valid int",
			value: 1,
			want: Option[int]{
				Value: 1,
				Valid: true,
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestNone(t *testing.T) {
	type test[T any] struct {
		name string
		want Option[T]
	}

	do := func(tt *test[int]) {
		t.Run(tt.name, func(t *testing.T) {
			got := None[int]()
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v", tt.want, got)
			}
		})
	}

	tests := []*test[int]{
		{
			name: "invalid int",
			want: Option[int]{
				Value: 0,
				Valid: false,
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
