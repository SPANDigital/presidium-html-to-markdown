package collector

import (
	"github.com/stretchr/testify/assert"
	"htmltomarkdown/config"
	"htmltomarkdown/test"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestCollect(t *testing.T) {
	var dataPath = test.DataPath(t)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" {
			path = "index.html"
		}
		body := test.MustReadFile(t, filepath.Join(dataPath, path))
		w.Write(body)
	})

	dstDir := test.TempDir(t)
	defer os.RemoveAll(dstDir)

	server := httptest.NewServer(mux)
	err := Collect(server.URL, dstDir, config.Config{})
	assert.NoError(t, err)

	err = fstest.TestFS(os.DirFS(dstDir), "index.html", "contact.html", "images/sample.png")
	assert.NoError(t, err)
}
