package main

import (
	"github.com/codegangsta/negroni"
	"html/template"
	"log"
	"net/http"
	"os"
)

var fm = template.FuncMap{
	"plus": func(a, b int) int {
		return a + b
	},
	"minus": func(a, b int) int {
		return a - b
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.
		New("").
		Funcs(fm).
		ParseFiles(
		"tmpl/index.html",
		"tmpl/partials.html",
	))

	dummy := make([]int, 6)

	err := templates.ExecuteTemplate(w, "index.html", struct {
		Iter []int
		Size int
	}{
		dummy,
		len(dummy),
	})
	if err != nil {
		log.Fatalf("error executing tmpl: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":" + os.Getenv("PORT"))
}
