package pkg

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestCollect(t *testing.T) {
	var testdata = filepath.Join("../testdata", t.Name())

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" {
			path = "index.html"
		}
		body := mustReadFile(t, filepath.Join(testdata, path))
		w.Write(body)
	})

	dstDir := testDir(t)
	defer os.RemoveAll(dstDir)

	server := httptest.NewServer(mux)
	err := Collect(server.URL, dstDir)
	assert.NoError(t, err)

	err = fstest.TestFS(os.DirFS(dstDir), "index.html", "contact.html", "images/sample.png")
	assert.NoError(t, err)
}
