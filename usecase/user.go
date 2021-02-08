package usecase

import (
	"fmt"
	"log"
	"time"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"
)

type UserUseCase interface {
	IsRecord(id string) bool
	Add(id, name string, regTime time.Time) error
	Get(id string) (*model.User, error)
	Update(id, name string) error
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
	return u.userRepository.IsRecord(id)
}

func (u userUseCase) Add(id, name string, regTime time.Time) error {

	user := model.User{
		ID:       id,
		Name:     name,
		RegAt:    regTime,
		UpdateAt: regTime,
	}

	if err := u.userRepository.Add(&user); err != nil {
		log.Println(err)
		return fmt.Errorf("couldn't create name=%s", name)
	}
	return nil
}

func (u userUseCase) Get(id string) (*model.User, error) {

	user, err := u.userRepository.Get(id)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("not found id=%s", id)
	}

	return user, nil
}

func (u userUseCase) Update(id, name string) error {

	user := model.User{
		ID:       id,
		Name:     name,
		UpdateAt: time.Now().Local(),
	}

	if err := u.userRepository.Update(&user); err != nil {
		log.Println(err)
		return fmt.Errorf("couldn't update user id=%s, name=%s", id, name)
	}
	return nil
}
