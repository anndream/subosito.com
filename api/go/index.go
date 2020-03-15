package handler

import (
	"net/http"

	"subosito.com/go/tiffany"
)

var repos = []string{"gotenv", "gozaru", "tiffany", "image64"}

func has(name string) bool {
	for i := range repos {
		if repos[i] == name {
			return true
		}
	}

	return false
}

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if !has(name) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tiffany.Render(w, tiffany.Option{
		CanonicalURL: "subosito.com/go/" + name,
		RepoURL:      "https://github.com/subosito/" + name,
	})
}
