package test

import (
	"fmt"
	"log"
	"todo_app_mod/config"
)

func Init() {
	/** ini ファイルの取得 */
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")

}
