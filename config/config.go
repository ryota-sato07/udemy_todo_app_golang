package config

import (
	"log"

	"todo_app_mod/utils"

	"gopkg.in/ini.v1"
)

/**
 * 構成リスト
 */
type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

/**
 * 構成リスト： 変数定義
 */
var Config ConfigList

/**
 * 構成リスト： 実行
 */
func init() {
	LoadConfg()
	utils.LoggingSettings(Config.LogFile)
}

/**
 * 構成リスト： 設定
 */
func LoadConfg() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(),
	}
}
