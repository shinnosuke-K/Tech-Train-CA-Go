package repository

import (
	"context"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"
)

type UserRepository interface {
	IsRecord(ctx context.Context, id int) bool
	Add(user *model.User) error
	Get(ctx context.Context, id int) (*model.User, error)
	Update(ctx context.Context) error
}
