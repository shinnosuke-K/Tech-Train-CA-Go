package persistence

import (
	"context"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (u userPersistence) IsRecord(ctx context.Context, id int) bool {
	panic("implement me")
}

func (u userPersistence) Add(ctx context.Context) error {
	panic("implement me")
}

func (u userPersistence) Get(ctx context.Context, id int) (*model.User, error) {
	panic("implement me")
}

func (u userPersistence) Update(ctx context.Context) error {
	panic("implement me")
}
