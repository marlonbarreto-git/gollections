package collection_test

import (
	"github.com/marlonbarreto-git/gollections/collection"
	"github.com/marlonbarreto-git/gollections/list"
	"reflect"
	"slices"
	"testing"
)

func TestList_MarshalJSON(t *testing.T) {
	type testCase[T any] struct {
		name    string
		list    collection.List[T]
		want    []byte
		wantErr bool
	}
	tests := []testCase[any]{
		{
			name:    "Given a numbers list Then return a correct json array",
			list:    list.Of[any](1, 2, 3),
			want:    []byte(`[1,2,3]`),
			wantErr: false,
		},
		{
			name:    "Given a strings list Then returns a correct json array",
			list:    list.Of[any]("1", "2", "3"),
			want:    []byte(`["1","2","3"]`),
			wantErr: false,
		},
		{
			name:    "Given a nil list Then returns a correct null",
			list:    nil,
			want:    []byte(`null`),
			wantErr: false,
		},
		{
			name:    "Given a empty list Then returns a correct empty json array",
			list:    collection.List[any]{},
			want:    []byte(`[]`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.list.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_UnmarshalJSON(t *testing.T) {
	type args struct {
		bytes []byte
	}
	type testCase[T any] struct {
		name         string
		expectedList collection.List[T]
		bytes        []byte
		wantErr      bool
	}
	tests := []testCase[any]{
		{
			name:         "Given a numbers json array Then returns a list.List",
			expectedList: list.Of[any](1.0, 2.0, 3.0),
			bytes:        []byte(`[1,2,3]`),
			wantErr:      false,
		},
		{
			name:         "Given a strings json array Then returns a list.List",
			expectedList: list.Of[any]("1.0", "2.0", "3.0"),
			bytes:        []byte(`["1.0","2.0","3.0"]`),
			wantErr:      false,
		},
		{
			name:         "Given a bad json array Then returns a error",
			expectedList: nil,
			bytes:        []byte(`["1.0","2.0,"3.0]`),
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testList := collection.List[any]{}

			if err := testList.UnmarshalJSON(tt.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !slices.Equal(testList.ToArray(), tt.expectedList.ToArray()) {
				t.Errorf("UnmarshalJSON() error = are not equals")
			}
		})
	}
}
