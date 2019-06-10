package funcs

import "github.com/dlclark/regexp2"

type MethodSignature struct {
	fullMethodSignature string
	namespace           string
	debug               string
}

//                                 return type       functionname       arg tyoe            arg name
const functionRegex = `^(?!\*|//)([A-Za-z0-9_]+)\s+([A-Za-z0-9_:]+)\((([A-Za-z0-9_]+)\s+([A-Za-z0-9_]+)(,\s*)?)*\)`

func (m MethodSignature) GetBaseFunctionName() string {
	re := regexp2.MustCompile(functionRegex, regexp2.Multiline)
	match, err := re.FindStringMatch(m.fullMethodSignature)
	if err != nil {
		panic(err)
	}
	if match == nil {
		panic("ahhhhhhh")
	}
	return match.GroupByNumber(2).String()
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
