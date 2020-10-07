package main

import (
	"net/http"

	"github.com/VinceJnz/go-test/handler"
	"github.com/VinceJnz/go-test/page"
	"github.com/VinceJnz/go-test/test"
)

func main() {
	defer test.LogFile.Close()
	test.LogFile.Write("Test entry")
	http.HandleFunc("/view/", handler.Make(page.ViewHandler))
	http.HandleFunc("/edit/", handler.Make(page.EditHandler))
	http.HandleFunc("/save/", handler.Make(page.SaveHandler))
	http.ListenAndServe(":8085", nil)
}
