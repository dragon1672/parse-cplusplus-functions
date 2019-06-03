package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"parse-cplusplus-functions-go/file_parsing"
	"parse-cplusplus-functions-go/funcs"
)

func getMissingFunctions(expectedFunctions []*funcs.MethodSignature, providedFunctions []*funcs.MethodSignature) []*funcs.MethodSignature {
	return nil // TODO
}

func getExpectedFunctions(filePath string) ([]*funcs.MethodSignature, error) {
	// Load from path
	// convert to functions
	return nil, nil // TODO
}

func getProvidedFunctions(dir string) ([]*funcs.MethodSignature, error) {
	headerFunctions, cppFunctions, err := file_parsing.ParseCPPAndHeaderFiles(dir)
	if err != nil {
		return nil, err
	}
	discoveredFunctions := []*funcs.MethodSignature{}
	// TODO remove duplicates
	discoveredFunctions = append(discoveredFunctions, headerFunctions...)
	discoveredFunctions = append(discoveredFunctions, cppFunctions...)
	return discoveredFunctions, nil
}

var (
	codePath = flag.String("codePath", "", "directory containing cpp and header files")
	masterDictPath = flag.String("masterDictPath ", "", "filepath of csv file with expected paths")
)


func main() {

	expectedFunctions, err := getExpectedFunctions(*masterDictPath)
	if err != nil {
		glog.Fatal(err)
	}
	providedFunctions, err := getProvidedFunctions(*codePath)
	if err != nil {
		glog.Fatal(err)
	}

	missingFunctions := getMissingFunctions(expectedFunctions, providedFunctions)

	fmt.Printf("Missing functions: %v", missingFunctions)
}
