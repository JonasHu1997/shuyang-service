package file

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func Read(filename string) ([]byte, error) {
	ok, err := Exist(filename)
	if !ok {
		return []byte{}, errors.New(fmt.Sprintf("[%s] is not exist", filename))
	} else if err != nil {
		return []byte{}, err
	}

	fileObj, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err == nil {
		defer func() {
			_ = fileObj.Close()
		}()
		buffer, err := ioutil.ReadAll(fileObj)
		//result := strings.Replace(string(buffer), "\n", "", 1)
		//return []byte(result), err
		return buffer, err
	}
	return []byte{}, nil
}

func Exist(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Size(filename string) int64 {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return -1
	}
	fileSize := fileInfo.Size()
	return fileSize
}
