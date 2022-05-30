package repository

import (
	"context"
	"fmt"
	"log"
	"server/ent/user"

	"server/ent"
)

type UserAuthRepository struct {
	client *ent.Client
}

func CreateUserAuthRepository(client *ent.Client) *UserAuthRepository {
	return &UserAuthRepository{client}
}

func (repo *UserAuthRepository) FindOneById(ctx context.Context, ID uint) (*ent.User, error) {
	user, err := repo.client.User.
		Query().
		Where(user.ID(int(ID))).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", user)
	return user, nil
}

func (repo *UserAuthRepository) FindOneByUsername(ctx context.Context, Username string) (*ent.User, error) {
	user, err := repo.client.User.
		Query().
		Where(user.Username(Username)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", user)
	return user, nil
}

func (repo *UserAuthRepository) Save(ctx context.Context, Username string, Password string, Role user.Role) (*ent.User, error) {
	user, err := repo.client.User.
		Create().
		SetPassword(Password).
		SetUsername(Username).
		SetRole(Role).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed saving user: %w", err)
	}
	return user, nil
}
