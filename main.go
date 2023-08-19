package main

import (
	"todo_app_mod/app/debug"
)

func main() {
	// fmt.Println(models.Db)

	/** ini の検証 */
	// debug.Init()

	/** User の検証 */
	// debug.User()

	/** Todo の検証 */
	// debug.Todo()

	/** Server の検証 */
	debug.Server()
}
