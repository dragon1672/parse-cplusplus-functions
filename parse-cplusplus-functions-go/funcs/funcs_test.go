package funcs

import (
	"reflect"
	"testing"
)

func TestMethodSignature_GetBaseFunctionName(t *testing.T) {
	type fields struct {
		fullMethodSignature string
		namespace           string
		debug               string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Basic Function no args",
			fields: fields{fullMethodSignature: "void simpleFunc()"},
			want:   "simpleFunc",
		},
		{
			name:   "single arg",
			fields: fields{fullMethodSignature: "void simpleFunc(int simple)"},
			want:   "simpleFunc",
		},
		{
			name:   "multi arg",
			fields: fields{fullMethodSignature: "void simpleFunc(int one, int two)"},
			want:   "simpleFunc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MethodSignature{
				fullMethodSignature: tt.fields.fullMethodSignature,
				namespace:           tt.fields.namespace,
				debug:               tt.fields.debug,
			}
			if got := m.GetBaseFunctionName(); got != tt.want {
				t.Errorf("MethodSignature.GetBaseFunctionName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMethodSignature_Matches(t *testing.T) {
	type fields struct {
		fullMethodSignature string
		namespace           string
		debug               string
	}
	type args struct {
		them MethodSignature
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := MethodSignature{
				fullMethodSignature: tt.fields.fullMethodSignature,
				namespace:           tt.fields.namespace,
				debug:               tt.fields.debug,
			}
			if got := me.Matches(tt.args.them); got != tt.want {
				t.Errorf("MethodSignature.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSignature(t *testing.T) {
	type args struct {
		fullSignature string
		debug         string
	}
	tests := []struct {
		name string
		args args
		want MethodSignature
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateSignature(tt.args.fullSignature, tt.args.debug); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateSignatureWithNamespace(t *testing.T) {
	type args struct {
		fullSignature string
		namespace     string
		debug         string
	}
	tests := []struct {
		name string
		args args
		want MethodSignature
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateSignatureWithNamespace(tt.args.fullSignature, tt.args.namespace, tt.args.debug); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSignatureWithNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
