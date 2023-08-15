package controllers

import (
	"net/http"
	"todo_app_mod/config"
)

/**
 * サーバー起動
 */
func StartMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
