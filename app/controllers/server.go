package controllers

import (
	"fmt"
	"lesson/config"
	"net/http"
	"text/template"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	template := template.Must(template.ParseFiles(files...))
	template.ExecuteTemplate(w, "layout", data)
}

func StartMainSever() error {
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
