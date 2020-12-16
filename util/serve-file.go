package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func serveFile(dir string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(dir, r.URL.Path)
		fmt.Println(path)

		http.ServeFile(w, r, path)
	}

	return http.HandlerFunc(fn)
}

func main() {
	static := "../util"

	http.Handle("/", serveFile(static))
	http.ListenAndServe(":3001", nil)

}
