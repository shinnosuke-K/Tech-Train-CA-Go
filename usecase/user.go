package usecase

import (
	"context"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"
)

type UserUseCase interface {
	IsRecord(ctx context.Context, id int) bool
	Add(user *model.User) error
	Get(ctx context.Context, id int) (*model.User, error)
	Update(ctx context.Context) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (u userUseCase) IsRecord(ctx context.Context, id int) bool {
	panic("implement me")
}

func (u userUseCase) Add(user *model.User) error {
	panic("implement me")
}

func (u userUseCase) Get(ctx context.Context, id int) (*model.User, error) {
	panic("implement me")
}

func (u userUseCase) Update(ctx context.Context) error {
	panic("implement me")
}
