package utils

import (
	"os"
	"path/filepath"
)

func FileName() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	fileName := filepath.Base(execPath)

	return fileName, nil
}
