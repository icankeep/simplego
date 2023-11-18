package utils

import (
	"os"
	"path/filepath"
)

func ReadStringFromFile(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	bytes, err := os.ReadFile(absPath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func MustReadStringFromFile(path string) string {
	s, err := ReadStringFromFile(path)
	if err != nil {
		panic(err)
	}
	return string(s)
}

func WriteStringToFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

func MustWriteStringToFile(path string, content string) {
	err := WriteStringToFile(path, content)
	if err != nil {
		panic(err)
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
