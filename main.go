package main

import (
	"todo_app_mod/app/test"
)

func main() {
	// fmt.Println(models.Db)

	/** ini の検証 */
	// test.Init()

	/** User の検証 */
	// test.User()

	/** Todo の検証 */
	test.Todo()

	/** Server の検証 */
	test.Server()
}
