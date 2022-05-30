package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"server/ent"
	"server/ent/issue"
	"server/repository"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type IssueController struct {
	repo *repository.IssueRepository
}

func CreateIssueController(repo *repository.IssueRepository) *IssueController {
	return &IssueController{repo}
}

func (issueController *IssueController) GetIssues(w http.ResponseWriter, r *http.Request) {
	issues, err := issueController.repo.GetAll(r.Context())
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		log.Print(err)
		repoError := ErrorResponse{Type: "internal error", Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(repoError)
		return
	}
	_ = json.NewEncoder(w).Encode(issues)
}

func (issueController *IssueController) GetIssue(w http.ResponseWriter, r *http.Request) {
	issueId, err := strconv.Atoi(chi.URLParam(r, "issueId"))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	foundIssue, err := issueController.repo.FindOne(r.Context(), issueId)
	if err != nil {
		log.Print(err)
		repoError := ErrorResponse{Type: "not found", Message: err.Error()}
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(repoError)
		return
	}
	_ = json.NewEncoder(w).Encode(foundIssue)
}

func (issueController *IssueController) UpdateIssue(w http.ResponseWriter, r *http.Request) {
	var issue ent.Issue
	issueId, err := strconv.Atoi(chi.URLParam(r, "issueId"))

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&issue)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = issueController.repo.Update(r.Context(), issueId, issue.Status, issue.Edges.AssignedTags)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (issueController *IssueController) SaveIssue(w http.ResponseWriter, r *http.Request) {
	var decodedIssue ent.Issue
	err := json.NewDecoder(r.Body).Decode(&decodedIssue)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	claims := r.Context().Value("props").(*Claims)
	creatorId := claims.Id
	_, err = issueController.repo.Save(r.Context(), creatorId, decodedIssue.Name, decodedIssue.Description, issue.StatusOpen, decodedIssue.Edges.AssignedTags)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
