package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//GetENV *
func GetENV(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}

//GetENVToBool *
func GetENVToBool(key string) bool {
	envStr := strings.TrimSpace(os.Getenv(key))
	boo, err := StringUtils(envStr).Bool()
	if nil != err {
		boo = false
	}
	return boo
}

//GetENVToInt *
func GetENVToInt(key string) (int, error) {
	envStr := strings.TrimSpace(os.Getenv(key))
	return StringUtils(envStr).Int()
}

//GetENVToInt64 *
func GetENVToInt64(key string) (int64, error) {
	envStr := strings.TrimSpace(os.Getenv(key))
	return StringUtils(envStr).Int64()
}

//GetBinPath 获取当前程序完整路径
func GetBinPath() (string, error) {
	file, _ := exec.LookPath(os.Args[0])
	absBinPath, _ := filepath.Abs(file)
	binPath, err := filepath.EvalSymlinks(absBinPath)
	if nil != err {
		return "", err
	}
	return binPath, nil
}

//GetBinDir 获取当前程序所在目录
func GetBinDir() string {
	binPath, _ := GetBinPath()

	binDir := filepath.Dir(binPath)
	binDir = strings.TrimSpace(binDir)
	binDir = strings.TrimRight(binDir, string(os.PathSeparator))

	return binDir
}
