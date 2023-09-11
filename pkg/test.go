package pkg

import (
	"encoding/json"
	"os"
	"testing"
)

func mustReadFile(t *testing.T, path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		t.Error(err)
		return nil
	}
	return content
}

func mustReadFileAsString(t *testing.T, path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		t.Error(err)
		return ""
	}
	return string(content)
}

func mustUnmarshal(t *testing.T, path string, value interface{}) {
	content, err := os.ReadFile(path)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(content, value)
	if err != nil {
		t.Error(err)
	}
}

func testDir(t *testing.T) string {
	path, err := os.MkdirTemp(".", "_test")
	if err != nil {
		t.Error(err)
	}
	return path
}
