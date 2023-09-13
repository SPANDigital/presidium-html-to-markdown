package test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func DataPath(t *testing.T) string {
	return filepath.Join("../test/data", t.Name())
}

func MustReadFile(t *testing.T, path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		t.Error(err)
		return nil
	}
	return content
}

func MustReadFileAsString(t *testing.T, path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		t.Error(err)
		return ""
	}
	return string(content)
}

func MustUnmarshal(t *testing.T, path string, value interface{}) {
	content, err := os.ReadFile(path)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(content, value)
	if err != nil {
		t.Error(err)
	}
}

func TempDir(t *testing.T) string {
	path, err := os.MkdirTemp(".", "_test")
	if err != nil {
		t.Error(err)
	}
	return path
}
