package main

import (
	"fmt"

	/** go.mod ファイルの module 名が、importパス(todo_app_mod) に影響する */
	"todo_app_mod/app/models"
	"todo_app_mod/app/test"
)

func main() {
	/** ini の検証 */
	// test.Init()

	/** User の検証 */
	test.User()

	fmt.Println(models.Db)

}
