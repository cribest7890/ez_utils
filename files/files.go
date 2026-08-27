package files

import (
	"os"
	"path/filepath"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func Write(path string, content string) error {
	dir := filepath.Dir(path)
	if dir != "." && !Exists(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return os.WriteFile(path, []byte(content), 0644)
}

func Append(path string, content string) error {
	dir := filepath.Dir(path)
	if dir != "." && !Exists(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	return err
}

func Read(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Clear(path string) error {
	return os.WriteFile(path, []byte(""), 0644)
}

func Delete(path string) error {
	if !Exists(path) {
		return nil
	}
	return os.Remove(path)
}

