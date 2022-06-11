package service

import (
	todo_app "todo-app"
	"todo-app/pkg/repository"
)

type Authorization interface{
   CreateUser(user todo_app.User) (int, error)
   GenerateToken(username, password string) (string, error)
   ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo_app.TodoList) (int, error)
	GetAll(userId int) ([]todo_app.TodoList, error)
	GetById(UserId, listId int)(todo_app.TodoList, error)
	Delete(UserId, listId int) error
	Update(userId, id int, input todo_app.UpdateListInput) error
}

type TodoItem  interface{

}

type Service struct{
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
		TodoList: newTodoListService(repos.TodoList),
	}
}