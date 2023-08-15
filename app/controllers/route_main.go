package controllers

import (
	"html/template"
	"log"
	"net/http"
)

/**
 * トップページ
 */
func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, "Hello")
}
