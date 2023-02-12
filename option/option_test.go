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

func TestOption_String(t *testing.T) {
	type test struct {
		name   string
		option Option[int]
		want   string
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.option.String()
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name:   "Some",
			option: Some(10),
			want:   "Some(10)",
		},
		{
			name:   "None",
			option: None[int](),
			want:   "None",
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestOption_MarshalJSON(t *testing.T) {
	type test struct {
		name    string
		option  Option[int]
		want    []byte
		wantErr bool
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.option.MarshalJSON()
			if tt.wantErr != (err != nil) {
				t.Fatalf("want-error=%v, error=%v.", tt.wantErr, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%s, got=%s.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name:    "Some",
			option:  Some(10),
			want:    []byte("10"),
			wantErr: false,
		},
		{
			name:    "None",
			option:  None[int](),
			want:    []byte("null"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}
