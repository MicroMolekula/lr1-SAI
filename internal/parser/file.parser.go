package parser

import (
	"errors"
	"os"
	"strings"
)

func GetFileData(fileName string) (string, error) {
	f, errOpen := os.Open(fileName)
	if errOpen == nil {
		fileInfo, errInfo := f.Stat()
		if errInfo != nil {
			return "", errors.New(errInfo.Error())
		}
		data := make([]byte, fileInfo.Size())
		_, errRead := f.Read(data)
		if errRead != nil {
			return "", errors.New(errRead.Error())
		}
		return strings.ToLower(string(data)), nil
	}
	defer f.Close()
	return "", errors.New("ошибка при открытии файла")
}
