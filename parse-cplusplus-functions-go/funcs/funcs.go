package funcs

type MethodSignature struct {
	fullMethodSignature string
	namespace           string
	debug               string
}

func (m MethodSignature) GetBaseFunctionName() string {
	// TODO extract only the base method
	return m.fullMethodSignature
}

func (me MethodSignature) Matches(them MethodSignature) bool {
	// TODO match based off function params too
	return me.GetBaseFunctionName() == them.GetBaseFunctionName()
}

func CreateSignature(fullSignature string, debug string) MethodSignature {
	return MethodSignature{
		fullMethodSignature: fullSignature,
		debug:               debug,
	}
}

func CreateSignatureWithNamespace(fullSignature string, namespace string, debug string) MethodSignature {
	return MethodSignature{
		fullMethodSignature: fullSignature,
		namespace:           namespace,
		debug:               debug,
	}
}
