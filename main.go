package main

import (
	"fmt"

	/** go.mod ファイルの module 名が、importパス(todo_app_mod) に影響する */
	"todo_app_mod/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)
	//
	// log.Println("test")

	fmt.Println(models.Db)
}
