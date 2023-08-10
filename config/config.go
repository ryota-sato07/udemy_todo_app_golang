package config

import "log"

/**
 * 構成リスト
 */
type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

var Config ConfigList

/**
 * 構成リスト： 関数の実行
 */
func init() {
	LoadConfg()
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
		Port:       cfg.Seqtion("web").Key("port").MustString("8000"),
		-SQLDriver: cfg.Seqtion("db").Key("driver").String(),
		-DbName:    cfg.Seqtion("db").Key("name").String(),
		-LogFile:   cfg.Seqtion("web").Key("logfile").String(),
	}
}
