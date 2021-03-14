package usecase

import (
	"database/sql"
	"time"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/logger"

	"github.com/pkg/errors"

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
	transaction    repository.Transaction
}

func NewUserUseCase(ur repository.UserRepository, tx repository.Transaction) UserUseCase {
	return &userUseCase{
		userRepository: ur,
		transaction:    tx,
	}
}

func (u userUseCase) IsRecord(id string) bool {
	return u.userRepository.IsRecord(id)
}

func (u userUseCase) Add(id, name string, regTime time.Time) error {

	err := u.transaction.DoInTx(func(tx *sql.Tx) error {
		user := model.User{
			ID:       id,
			Name:     name,
			RegAt:    regTime,
			UpdateAt: regTime,
		}

		if err := u.userRepository.Add(tx, &user); err != nil {
			return errors.Errorf("couldn't create name=%s", name)
		}
		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	logger.Log.Info("[method:Add] finished adding")
	return nil
}

func (u userUseCase) Get(id string) (*model.User, error) {

	user, err := u.userRepository.Get(id)
	if err != nil {
		return nil, errors.Errorf("not found id=%s", id)
	}

	return user, nil
}

func (u userUseCase) Update(id, name string) error {

	err := u.transaction.DoInTx(func(tx *sql.Tx) error {
		user := model.User{
			ID:       id,
			Name:     name,
			UpdateAt: time.Now().Local(),
		}

		if err := u.userRepository.Update(tx, &user); err != nil {
			return errors.Errorf("couldn't update user id=%s, name=%s", id, name)
		}
		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	logger.Log.Info(" [method:Update] finished updating")
	return nil
}
