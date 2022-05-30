package repository

import (
	"context"
	"fmt"
	"server/ent"
	"server/ent/user"
)

type UserRepository struct {
	client *ent.Client
}

func CreateUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{client}
}

func (repo *UserRepository) FindOne(ctx context.Context, ID int) (*ent.User, error) {
	foundUser, err := repo.client.User.
		Query().
		Select(user.FieldID).
		Select(user.FieldUsername).
		Select(user.FieldRole).
		Where(user.ID(ID)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed saving user: %w", err)
	}

	return foundUser, nil
}
