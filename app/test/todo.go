package test

import (
	"fmt"
	"todo_app_mod/app/models"
)

func Todo() {
	/** Todo の作成 */
	// user, _ := models.GetUser(2)
	// user.CreatedTodo("First Todo")

	/** Todo の取得 */
	t, _ := models.GetTodo(1)
	fmt.Println(t)
}
