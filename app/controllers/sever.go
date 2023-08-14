package controllers

import (
	"net/http"
	"todo_app_mod/config"
)

func StartMainServer() error {
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
