package controller

import (
	"net/http"
	"server/repository"
)

type UserController struct {
	repo repository.IUserRepository
}

func CreateUserController(repo repository.IUserRepository) *UserController {
	return &UserController{repo: repo}
}

func (userController *UserController) GetProjects(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("hello"))
}
