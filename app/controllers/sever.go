package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"todo_app_mod/app/models"
	"todo_app_mod/config"
)

/**
 * 共通レイアウト・各ページのHTMLを読み込む
 */
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

/**
 * ログイン済みユーザのCookieを取得
 */
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

/**
 * サーバー起動
 */
func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	/** ログイン前のみアクセス可能 */
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)

	/** ログイン後のみアクセス可能 */
	http.HandleFunc("/todos", index)
	http.HandleFunc("/logout", logout)

	/** 認証 */
	http.HandleFunc("/authenticate", authenticate)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
