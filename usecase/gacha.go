package usecase

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/logger"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/db"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/model"

	"github.com/google/uuid"

	"github.com/pkg/errors"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"
)

type Result struct {
	CharaId string
	Name    string
}

type GachaUseCase interface {
	Draw(times int) ([]*Result, error)
	Store(id string, results []*Result) error
}

type gachaUseCase struct {
	gachaRepository repository.GachaRepository
	transaction     db.Transaction
}

func NewGachaUseCase(ug repository.GachaRepository, tx db.Transaction) GachaUseCase {
	return &gachaUseCase{
		gachaRepository: ug,
		transaction:     tx,
	}
}

func (g gachaUseCase) Draw(times int) ([]*Result, error) {

	logger.Log.Info("[method:Draw] start")

	gacha, err := g.gachaRepository.GetRareRate()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get the Gacha record")
	}

	chara, err := g.gachaRepository.GetCharacter()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get the Character record")
	}

	countMap := make(map[int]int, 0)
	for _, c := range chara {
		countMap[c.Rarity] += 1
	}

	var totalWeights int
	for _, g := range gacha {
		totalWeights += countMap[g.Rarity] * g.Weights
	}

	r := make([]*Result, 0, times)
	for n := 0; n < times; n++ {
		p := rand.Intn(totalWeights)
		total := 0
		for _, c := range chara {
			for _, g := range gacha {
				if c.Rarity == g.Rarity {
					total += g.Weights
					break
				}
			}
			if p <= total {
				r = append(r, &Result{
					CharaId: c.ID,
					Name:    c.Name,
				})
				break
			}
		}
	}

	logger.Log.Info("[method:Draw] finished")
	return r, nil
}

func (g gachaUseCase) Store(id string, results []*Result) error {

	logger.Log.Info("[method:Store] start")

	err := g.transaction.DoInTx(func(tx *sql.Tx) error {
		for _, r := range results {
			posse := model.Possession{
				ID:      uuid.New().String(),
				UserID:  id,
				CharaID: r.CharaId,
				RegAt:   time.Now().Local(),
			}
			if err := g.gachaRepository.Store(tx, &posse); err != nil {
				return errors.Wrapf(err, "couldn't store character id=%s", r.CharaId)
			}
		}
		return nil
	})

	if err != nil {
		return errors.WithStack(err)
	}

	logger.Log.Info("[method:Store] finished")
	return nil
}
