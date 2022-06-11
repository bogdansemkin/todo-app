package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo_app "todo-app"
)

type TodoListPostgres struct{
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todo_app.TodoList) (int, error){
	//start transaction
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title,description) VALUES($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err2 := row.Scan(&id); err2 != nil{
		tx.Rollback()
		return 0, err2
	}
	createUserListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUserListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
