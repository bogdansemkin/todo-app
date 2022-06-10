package service

import (
	"crypto/sha1"
	"fmt"
	todo_app "todo-app"
	"todo-app/pkg/repository"
)

const solt = "efgsnjrtldgjriesot"

type AuthService struct  {
	repo repository.Authorization
}

func newAuthService (repo repository.Authorization) *AuthService{
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo_app.User)(int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string{
	hash :=  sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(solt)))
}
