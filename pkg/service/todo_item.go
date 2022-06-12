package service

import (
	todo_app "todo-app"
	"todo-app/pkg/repository"
)

type todoItemService struct {
	repo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *todoItemService{
	return &todoItemService{repo: repo, listRepo: listRepo}
}

func (s *todoItemService) Create(userId int, listId int, item todo_app.TodoItem) (int, error){
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil{
		return 0, err
	}
	return s.repo.Create(listId, item)
}

func (s *todoItemService) GetAll(userId, listId int)([]todo_app.TodoItem, error){
	return s.repo.GetAll(userId, listId)
}

func (s *todoItemService) GetById(userId, itemId int)(todo_app.TodoItem, error){
	return s.repo.GetById(userId, itemId)
}

func (s *todoItemService) Delete(userId, itemId int) error{
	return s.repo.Delete(userId, itemId)
}

func (s *todoItemService) Update(userId, itemId int, input todo_app.UpdateItemInput) error{
	return s.repo.Update(userId, itemId, input)
}