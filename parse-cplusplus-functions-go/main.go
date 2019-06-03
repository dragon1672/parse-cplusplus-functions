package main

import (
	"fmt"
	"github.com/golang/glog"
)

type MethodSignature struct {
	fullMethodSignature string
}

func main() {

	expectedFunctions, err := getExpectedFunctions()
	if err != nil {
		glog.Fatal(err)
	}
	providedFunctions, err := getProvidedFunctions()
	if err != nil {
		glog.Fatal(err)
	}

	missingFunctions := getMissingFunctions(expectedFunctions, providedFunctions)

	fmt.Printf("Missing functions: %v", missingFunctions)
}

func getMissingFunctions(expectedFunctions []*MethodSignature, providedFunctions []*MethodSignature) []*MethodSignature {
	return nil // TODO
}

func getExpectedFunctions() ([]*MethodSignature, error) {
	// Load from path
	// convert to functions
	return nil, nil // TODO
}
func getProvidedFunctions() ([]*MethodSignature, error) {
	// Load from path
	// convert to functions
	return nil, nil // TODO
}
