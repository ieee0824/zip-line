package tree

import (
	"os"
	"path/filepath"
)

func Tree(root string) ([]string, error) {
	ret := make([]string, 0, 1024)
	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		ret = append(ret, path)
		return nil
	}); err != nil {
		return nil, err
	}
	return ret, nil
}
