package funcs

type MethodSignature struct {
	fullMethodSignature string
}


func CreateSignature(fullSignature string) *MethodSignature {
	return &MethodSignature{
		fullMethodSignature:fullSignature,
	}
}
