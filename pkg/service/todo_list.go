package service

import (
	todo_app "todo-app"
	"todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func newTodoListService(repo repository.TodoList) *TodoListService{
	return &TodoListService{repo: repo}
}

func (s * TodoListService) Create(userId int, list todo_app.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo_app.TodoList, error){
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int)(todo_app.TodoList, error){
	return s.repo.GetById(userId, listId)
}

