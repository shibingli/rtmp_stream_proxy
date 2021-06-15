package utils

import (
	"os"
	"strings"
)

//Exist 文件是否存在
func Exist(path string) (bool, os.FileInfo, error) {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return true, fileInfo, nil
	}
	if os.IsNotExist(err) {
		return false, nil, nil
	}
	return false, nil, err
}

//Empty *
func Empty(value ...string) bool {
	for _, val := range value {
		if "" == strings.TrimSpace(val) {
			return true
		}
	}
	return false
}

//IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func PathJoin(path ...string) string {
	return strings.Join(path, string(os.PathSeparator))
}
