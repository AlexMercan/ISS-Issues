package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"server/repository"
)

type IssueTagController struct {
	repo *repository.IssueTagRepository
}

func CreateIssueTagController(repo *repository.IssueTagRepository) *IssueTagController {
	return &IssueTagController{repo}
}

func (IssueTagController *IssueTagController) GetAll(w http.ResponseWriter, r *http.Request) {
	tags, err := IssueTagController.repo.GetAll(r.Context())
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		log.Print(err)
		repoError := ErrorResponse{Type: "internal error", Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(repoError)
		return
	}
	_ = json.NewEncoder(w).Encode(tags)
}
