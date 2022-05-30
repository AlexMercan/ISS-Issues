package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/repository"
)

type UserController struct {
	repo *repository.UserRepository
}

func CreateUserController(repo *repository.UserRepository) *UserController {
	return &UserController{repo}
}

func (userController *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("props").(*Claims)
	fmt.Printf("Claims: %v", claims)
	user, err := userController.repo.FindOne(r.Context(), claims.Id)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		log.Print(err)
		repoError := ErrorResponse{Type: "internal error", Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(repoError)
		return
	}
	_ = json.NewEncoder(w).Encode(user)
}
