package file_parsing

import (
	"github.com/dlclark/regexp2"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"parse-cplusplus-functions-go/funcs"
	"path/filepath"
	"strings"
)

//                                 return type       functionname       arg tyoe            arg name
const functionRegex = `^(?!\*|//)([A-Za-z0-9_]+)\s+([A-Za-z0-9_:]+)\((([A-Za-z0-9_]+)\s+([A-Za-z0-9_]+)(,\s*)?)*\)`

func listFiles(dir string) ([]string, error) {
	files := []string{}

	visit := func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrapf(err, "error visiting path %s", path)
		}
		if !f.IsDir() {
			files = append(files, path)
		}
		return nil
	}

	err := filepath.Walk(dir, visit)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// ParseCPPAndHeaderFiles parses the
func ParseCPPAndHeaderFiles(dir string) ([]funcs.MethodSignature, []funcs.MethodSignature, error) {
	filePaths, err := listFiles(dir)
	if err != nil {
		return nil, nil, err
	}
	headerFiles := []funcs.MethodSignature{}
	cppFiles := []funcs.MethodSignature{}
	for _, filePath := range filePaths {
		// TODO optimize so no all files are being loaded
		fileData, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, nil, err
		}

		if strings.HasSuffix(strings.ToLower(filePath), strings.ToLower(".h")) {
			headerFiles = append(headerFiles, parseExpectedFunctions(string(fileData), filePath)...)
		} else if strings.HasSuffix(strings.ToLower(filePath), strings.ToLower(".cpp")) {
			cppFiles = append(cppFiles, parseExpectedFunctions(string(fileData), filePath)...)
		}
		// don't care about other files
	}
	return headerFiles, cppFiles, nil
}

func parseExpectedFunctions(fileData string, debug string) []funcs.MethodSignature {
	discoveredFunctions := []funcs.MethodSignature{}
	re := regexp2.MustCompile(functionRegex, regexp2.Multiline)
	for _, line := range strings.Split(fileData, "\n") {
		match, err := re.FindStringMatch(line)
		if err != nil {
			panic(err)
		}
		if match != nil {
			discoveredFunctions = append(discoveredFunctions, funcs.CreateSignature(match.Capture.String(), debug))
		}
	}
	return discoveredFunctions
}
