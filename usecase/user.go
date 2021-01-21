package usecase

import (
	"time"

	"github.com/pkg/errors"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"
)

type UserUseCase interface {
	IsRecord(id string) bool
	Add(id, name string, regTime time.Time) error
	Get(id string) (*model.User, error)
	Update() error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (u userUseCase) IsRecord(id string) bool {
	return u.IsRecord(id)
}

func (u userUseCase) Add(id, name string, regTime time.Time) error {

	user := model.User{
		UserId:   id,
		UserName: name,
		RegAt:    regTime,
		UpdateAt: regTime,
	}

	if err := u.userRepository.Add(&user); err != nil {
		return errors.Wrap(err, "user table couldn't create")
	}
	return nil
}

func (u userUseCase) Get(id string) (*model.User, error) {
	panic("implement me")
}

func (u userUseCase) Update() error {
	panic("implement me")
}
