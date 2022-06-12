package repository

import (
	"github.com/jmoiron/sqlx"
	todo_app "todo-app"
)

type Authorization interface{
	CreateUser(user todo_app.User) (int, error)
	GetUser(username, password string) (todo_app.User, error)
}

type TodoList interface {
	Create(userId int, item todo_app.TodoList) (int, error)
	GetAll(userId int) ([]todo_app.TodoList, error)
	GetById(UserId, listId int)(todo_app.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, id int, input todo_app.UpdateListInput) error
}

type TodoItem  interface{
	Create(listId int, item todo_app.TodoItem) (int, error)
	GetAll(UserId, listId int)([]todo_app.TodoItem, error)
	GetById(UserId, itemId int)(todo_app.TodoItem, error)
	Delete(UserId, itemId int) error
	Update(userId, itemId int, input todo_app.UpdateItemInput) error
}

type Repository struct{
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
		TodoItem: newTodoItemPostgres(db),
	}
}
