package master_functions

import (
	"encoding/csv"
	"fmt"
	"os"
	"parse-cplusplus-functions-go/funcs"
)

type csvLine struct {
	Namespace   string
	FunctionSig string
}

func loadCSV(filePath string) ([]*csvLine, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	csvData := []*csvLine{}
	for _, line := range lines {
		csvData = append(csvData,
			&csvLine{
				Namespace:   line[0],
				FunctionSig: line[1],
			})
	}
	return csvData, nil
}

func parseFunctions(d []*csvLine, debugString string) []funcs.MethodSignature {
	ret := []funcs.MethodSignature{}
	for _, line := range d {
		ret = append(ret, funcs.CreateSignatureWithNamespace(line.FunctionSig, line.Namespace, debugString))
	}
	return ret
}

func GetExpectedFunctions(filePath string) ([]funcs.MethodSignature, error) {
	csvData, err := loadCSV(filePath)
	if err != nil {
		return nil, err
	}

	functions := parseFunctions(csvData, fmt.Sprintf("%s %s", "config_file:", filePath))

	return functions, nil
}
