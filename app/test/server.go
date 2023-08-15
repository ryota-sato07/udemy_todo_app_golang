package test

import (
	"fmt"
	"todo_app_mod/app/controllers"
	"todo_app_mod/app/models"
)

func Server() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
