package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func Write(content []byte, filename string) (int64, error) {
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		return 0, err
	}

	_, err = os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return 0, err
		}
	}
	err = ioutil.WriteFile(filename, content, 0755)
	if err == nil {
		return Size(filename), nil
	}
	return 0, err

}

func Append(buffer []byte, filename string) (int, error) {
	fd, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	n, err := fd.Write(buffer)
	defer func() {
		_ = fd.Close()
	}()
	return n, err
}
