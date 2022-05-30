package repository

import (
	"context"
	"fmt"
	"server/ent"
	"server/ent/issue"
)

type IssueRepository struct {
	client *ent.Client
}

func CreateIssueRepository(client *ent.Client) *IssueRepository {
	return &IssueRepository{client}
}

func (repo *IssueRepository) Save(ctx context.Context, creatorID int, Name string, Description string, Status issue.Status, tags []*ent.IssueTag) (*ent.Issue, error) {
	savedIssue, err := repo.client.Issue.
		Create().
		SetIssueCreatorID(creatorID).
		SetName(Name).
		SetDescription(Description).
		SetStatus(Status).
		AddAssignedTags(tags...).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed saving user: %w", err)
	}

	return savedIssue, nil
}

func (repo *IssueRepository) FindOne(ctx context.Context, issueId int) (*ent.Issue, error) {
	issues, err := repo.client.Issue.
		Query().
		Where(issue.ID(issueId)).
		Select(issue.FieldID).
		Select(issue.FieldName).
		Select(issue.FieldStatus).
		Select(issue.FieldOwnerID).
		Select(issue.FieldDescription).
		WithAssignedTags().
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed saving user: %w", err)
	}

	return issues, nil
}

func (repo *IssueRepository) Update(ctx context.Context, IssueID int, Status issue.Status, tags []*ent.IssueTag) error {
	_, err := repo.client.Issue.
		Update().
		SetStatus(Status).
		ClearAssignedTags().
		AddAssignedTags(tags...).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("failed saving user: %w", err)
	}

	return nil
}

func (repo *IssueRepository) GetAll(ctx context.Context) ([]*ent.Issue, error) {
	issues, err := repo.client.Issue.
		Query().
		Select(issue.FieldID).
		Select(issue.FieldName).
		Select(issue.FieldStatus).
		Select(issue.FieldOwnerID).
		WithAssignedTags().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed saving user: %w", err)
	}

	return issues, nil
}
