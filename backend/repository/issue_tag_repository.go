package repository

import (
	"context"
	"fmt"
	"server/ent"
	"server/ent/issuetag"
)

type IssueTagRepository struct {
	client *ent.Client
}

func CreateIssueTagRepository(client *ent.Client) *IssueTagRepository {
	return &IssueTagRepository{client}
}

func (repo *IssueTagRepository) GetAll(ctx context.Context) ([]*ent.IssueTag, error) {
	tags, err := repo.client.IssueTag.
		Query().
		Select(issuetag.FieldID).
		Select(issuetag.FieldName).All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed getting issue tags: %w", err)
	}

	return tags, nil
}
