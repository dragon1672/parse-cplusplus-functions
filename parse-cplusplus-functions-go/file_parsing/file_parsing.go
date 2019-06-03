package file_parsing

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

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
