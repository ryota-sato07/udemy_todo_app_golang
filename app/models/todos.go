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
 * 単数タスクの取得関数
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

/**
 * 複数タスクの取得関数
 */
func GetTodos() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

/**
 * ユーザーIDからタスクを取得するメソッド
 */
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos where user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

/**
 * タスクの更新メソッド
 */
func (t *Todo) UpdateTodo() error {
	cmd := `update todos set content = ?, user_id = ? where id = ?`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
