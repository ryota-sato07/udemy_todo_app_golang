package models

import (
	"log"
	"time"
)

/**
 * タスク情報
 */
type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

/**
 * タスクの作成メソッド
 */
func (u *User) CreatedTodo(content string) (err error) {
	cmd := `insert into todos (
content,
user_id,
created_at) values (?, ?, ?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

/**
 * タスクの取得関数
 */
func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, content, user_id, created_at
	from todos where id = ?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	return todo, err
}
