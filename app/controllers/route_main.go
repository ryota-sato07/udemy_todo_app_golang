package controllers

import (
	"net/http"
)

/**
 * トップページ
 */
func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "public_navbar", "top")
}
