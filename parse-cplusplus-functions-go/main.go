package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"parse-cplusplus-functions-go/file_parsing"
	"parse-cplusplus-functions-go/funcs"
	"parse-cplusplus-functions-go/master_functions"
)

func keyFunctionsByName(signatures []funcs.MethodSignature) map[string][]funcs.MethodSignature {
	m := make(map[string][]funcs.MethodSignature)
	for _, f := range signatures {
		baseName := f.GetBaseFunctionName()
		if _, ok := m[baseName]; !ok {
			// Init the array in the map
			m[baseName] = []funcs.MethodSignature{}
		}
		m[baseName] = append(m[baseName], f)
	}
	return m
}

func functionInList(f funcs.MethodSignature, l []funcs.MethodSignature) bool {
	for _, toCheck := range l {
		if f.Matches(toCheck) {
			return true
		}
	}
	return false
}

func getMissingFunctions(expectedFunctions []funcs.MethodSignature, providedFunctions []funcs.MethodSignature) []funcs.MethodSignature {
	expectedMap := keyFunctionsByName(expectedFunctions)
	providedMap := keyFunctionsByName(providedFunctions)
	missingFunctions := []funcs.MethodSignature{}
	for expectedFunctionName, expectedFunctions := range expectedMap {
		for _, expectedFunction := range expectedFunctions {
			//check if this method is in provided methods
			if providedFunctions, ok := providedMap[expectedFunctionName]; ok {
				if !functionInList(expectedFunction, providedFunctions) {
					missingFunctions = append(missingFunctions, expectedFunction)
				}
			} else {
				missingFunctions = append(missingFunctions, expectedFunction)
			}
		}
	}

	return missingFunctions
}

func getProvidedFunctions(dir string) ([]funcs.MethodSignature, error) {
	headerFunctions, cppFunctions, err := file_parsing.ParseCPPAndHeaderFiles(dir)
	if err != nil {
		return nil, err
	}
	discoveredFunctions := []funcs.MethodSignature{}
	glog.Infof("Discovered %d header functions", len(headerFunctions))
	//discoveredFunctions = append(discoveredFunctions, headerFunctions...)

	glog.Infof("Discovered %d cpp functions", len(cppFunctions))
	discoveredFunctions = append(discoveredFunctions, cppFunctions...)
	return discoveredFunctions, nil
}

var (
	codePath       = flag.String("codePath", "", "directory containing cpp and header files")
	masterDictPath = flag.String("masterDictPath ", "", "filepath of csv file with expected paths")
)

func main() {

	expectedFunctions, err := master_functions.GetExpectedFunctions(*masterDictPath)
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
