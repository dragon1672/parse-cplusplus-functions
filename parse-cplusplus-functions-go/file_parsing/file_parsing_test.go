package file_parsing

import (
	"parse-cplusplus-functions-go/funcs"
	"reflect"
	"testing"
)

func Test_ParseExpectedFunctions(t *testing.T) {
	tcs := []struct {
		description string
		input       string
		want        []funcs.MethodSignature
	}{
		{
			description: "dead simple func",
			input:       `void maFunction()`,
			want: []funcs.MethodSignature{
				funcs.CreateSignature("void maFunction()", "testing"),
			},
		},
		{
			description: "simple hello world",
			input: `
void maFunction() {
 std::cout << "hello world"
}
`,
			want: []funcs.MethodSignature{
				funcs.CreateSignature("void maFunction()", "testing"),
			},
		},
		{
			description: "class prefixed functions",
			input: `
void someClass::maFunction() {
 std::cout << "hello world"
}
`,
			want: []funcs.MethodSignature{
				funcs.CreateSignature("void someClass::maFunction()", "testing"),
			},
		},
		{
			description: "avoids comments //",
			input:       `// void commentedFunction() {}`,
			want:        []funcs.MethodSignature{},
		},
		{
			description: "avoids while",
			input:       `while (true) {`,
			want:        []funcs.MethodSignature{},
		},
		{
			description: "avoids for",
			input:       `for true {`,
			want:        []funcs.MethodSignature{},
		},
		{
			description: "avoids if",
			input:       `if true {`,
			want:        []funcs.MethodSignature{},
		},
		{
			description: "avoids ML_LOG",
			input:       `ML_LOG true {`,
			want:        []funcs.MethodSignature{},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			got := parseExpectedFunctions(tc.input, "testing")
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("want %v, got %v", tc.want, got)
			}
		})
	}
}
