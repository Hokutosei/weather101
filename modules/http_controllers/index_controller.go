package http_controllers

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("index rendered...")
	indexTemplate := "index.html"
	t := template.New(indexTemplate).Delims("{{%", "%}}")
	// indexVars := IndexVars{}

	parsed_template_str := fmt.Sprintf("public/%s", indexTemplate)
	t, _ = t.ParseFiles(parsed_template_str)
	t.Execute(w, nil)
}
