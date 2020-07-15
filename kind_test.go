package reflecthelper

import (
	"reflect"
	"testing"
)

func TestGetKind(t *testing.T) {
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantRes reflect.Kind
	}{
		{
			name: "invalid kind",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantRes: reflect.Invalid,
		},
		{
			name: "normal kind",
			args: args{
				val: reflect.ValueOf(int(5)),
			},
			wantRes: reflect.Int,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetKind(tt.args.val); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetKind() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestGetElemKind(t *testing.T) {
	var valInt *int
	test := 5
	valInt = &test

	var valNilPtr **int
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantRes reflect.Kind
	}{
		{
			name: "invalid kind",
			args: args{
				val: reflect.ValueOf(nil),
			},
			wantRes: reflect.Invalid,
		},
		{
			name: "normal kind",
			args: args{
				val: reflect.ValueOf(int(5)),
			},
			wantRes: reflect.Int,
		},
		{
			name: "normal pointer kind",
			args: args{
				val: reflect.ValueOf(valInt),
			},
			wantRes: reflect.Int,
		},
		{
			name: "nil pointer kind",
			args: args{
				val: reflect.ValueOf(valNilPtr),
			},
			wantRes: reflect.Ptr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := GetElemKind(tt.args.val); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetElemKind() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}