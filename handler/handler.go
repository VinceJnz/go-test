package handler

import (
	"net/http"
	"regexp"

	"github.com/VinceJnz/go-test/page"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

//Make creates a handler for processing web requests
func Make(fn func(http.ResponseWriter, *http.Request, page.Page)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		var p page.Page
		err := p.Load(m[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, p)
	}
}
