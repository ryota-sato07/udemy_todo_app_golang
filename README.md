# udemy_todo_app_golang

**[【Go入門】Golang基礎入門 + 各種ライブラリ + 簡単なTodoWebアプリケーション開発(Go言語)](https://www.udemy.com/course/golang-webgosql/)** を写経して作成した、Todoアプリケーション

## 機能一覧

- ユーザー登録
- ログイン
- ログアウト
- Todo一覧
- Todo追加
- Todo編集
- Todo削除

## テーブル設計

### User: ユーザー情報

| フィールド | 型 | Extra | 説明 |
| --- | --- | --- | --- |
| id | int | AUTO_INC | ユーザーID |
| uuid | string |  | 一意の文字列 |
| name | string |  | ユーザー名 |
| email | string |  | メールアドレス |
| password | string |  | パスワード |
| created_at | time.Time |  | 作成日 |
| todos | []Todo |  | Todo一覧 |

### Session: セッション情報

| フィールド | 型 | Extra | 説明 |
| --- | --- | --- | --- |
| id | int | AUTO_INC | セッションID |
| uuid | string |  | 一意の文字列 |
| email | string |  | メールアドレス |
| user_id | int |  | ユーザーID |
| created_at | time.Time |  | 作成日 |

### Todo: タスク情報

| フィールド | 型 | Extra | 説明 |
| --- | --- | --- | --- |
| id | int | AUTO_INC | タスクID |
| content | string |  | 内容 |
| user_id | int |  | ユーザーID |
| created_at | time.Time |  | 作成日 |

## 開発方法

```
# ルートディレクトリにいることを確認
$ pwd
xxxxxx/go/src/udemy_todo_app_golang

# 開発サーバーを起動(http://127.0.0.1:8080/)
$ go run ./main.go
```
