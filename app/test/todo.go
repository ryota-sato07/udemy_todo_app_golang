package test

import (
	"fmt"
	"todo_app_mod/app/models"
)

func Todo() {
	/** Todo の作成 */
	user, _ := models.GetUser(2)
	user.CreatedTodo("Second Todo")

	/** Todo の単数取得 */
	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	/** Todo の複数取得 */
	todos, _ := models.GetTodos()
	for _, v := range todos {
		fmt.Println(v)
	}
}
