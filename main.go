package main

import (
	"fmt"

	/** go.mod ファイルの module 名が、importパス(todo_app_mod) に影響する */
	"todo_app_mod/app/models"
)

func main() {
	/** ini ファイルの取得 */
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)
	//
	// log.Println("test")

	fmt.Println(models.Db)

	/** ユーザーの作成 */
	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)
	//
	// u.CreateUser()

	/** ユーザーの取得 */
	u, _ := models.GetUser(1)
	fmt.Println(u)

	/** ユーザーの更新 */
	fmt.Println("========== [ ユーザーの更新 ] ==========")
	u.Name = "Test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
	fmt.Println("========== [ ユーザーの更新 ] ==========")

	/** ユーザーの削除 */
	u.DeleteUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}
