package controllers

import (
	"fmt"
	"lesson/app/models"
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
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}

func session(w http.ResponseWriter, r *http.Request) (session models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		session = models.Session{UUID: cookie.Value}
		if ok, _ := session.CheckSession(); !ok {
			err = fmt.Errorf("invalid session")
		}
	}
	return session, err
}
