package page

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

//Page defines a web page structure
type Page struct {
	Title string
	Body  []byte
}

//Load for loading a page from a file
func (p *Page) Load(title string) error {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	p.Title = title
	p.Body = body
	return nil
}

//ViewHandler handler for pages
func ViewHandler(w http.ResponseWriter, r *http.Request, p Page) {
	p.Render(w, "view")
}

//EditHandler handler for pages
func EditHandler(w http.ResponseWriter, r *http.Request, p Page) {
	p.Render(w, "edit")
}

//SaveHandler handler for pages
func SaveHandler(w http.ResponseWriter, r *http.Request, p Page) {
	body := r.FormValue("body")
	p.Body = []byte(body)
	filename := p.Title + ".txt"
	err := ioutil.WriteFile(filename, p.Body, 0600)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+p.Title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

//Render for rendering a template
func (p *Page) Render(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
