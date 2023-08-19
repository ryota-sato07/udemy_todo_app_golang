package debug

import (
	"todo_app_mod/app/models"
)

func Todo() {
	/** Todo の作成 */
	// user, _ := models.GetUser(3)
	// user.CreatedTodo("Third Todo")

	/** Todo の単数取得 */
	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	/** Todo の複数取得 */
	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	/** User から Todo を取得 */
	// user2, _ := models.GetUser(2)
	// todos, _ := user2.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	/** Todo の更新 */
	// t, _ := models.GetTodo(1)
	// t.Content = "Update Todo"
	// t.UpdateTodo()

	/** Todo の更新 */
	t, _ := models.GetTodo(3)
	t.DeleteTodo()
}
