package test

import "todo_app_mod/app/models"

func Todo() {
	/** Todo の作成 */
	user, _ := models.GetUser(2)
	user.CreatedTodo("First Todo")
}
