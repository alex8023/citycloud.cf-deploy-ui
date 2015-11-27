package utils

import (
	"os"
	"path/filepath"
)

//返回文件路径，和文件名
func Detechfile(absfile string) (string, string, bool, error) {

	dir := filepath.Dir(absfile)

	base := filepath.Base(absfile)

	file, err := os.Open(absfile)

	if err != nil {
		return "", "", false, err
	}

	info, errs := file.Stat()

	if errs != nil {
		return "", "", false, errs
	}

	isDir := info.IsDir()

	return dir, base, isDir, nil
}
