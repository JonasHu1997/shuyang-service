package file

import (
	"os"
	"path/filepath"
)

func Move(src string, des string) (bool, error) {
	var err error
	f, err := os.Stat(src)
	if err != nil {
		return false, err
	}
	if f.IsDir() {
		_ = os.MkdirAll(des, 0755)
	} else {
		_ = os.MkdirAll(filepath.Dir(des), 0755)
	}
	err = os.Rename(src, des)
	if err != nil {
		return false, err
	}
	return true, nil
}
