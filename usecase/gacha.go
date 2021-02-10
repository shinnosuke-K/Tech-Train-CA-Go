package usecase

import (
	"math/rand"
	"time"

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
}

func NewGachaUseCase(ug repository.GachaRepository) GachaUseCase {
	return &gachaUseCase{
		gachaRepository: ug,
	}
}

func (g gachaUseCase) Draw(times int) ([]*Result, error) {

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

	rand.Seed(time.Now().Unix())

	r := make([]*Result, 0, times)
	for n := 0; n < times; n++ {
		p := rand.Float64() * 100
		total := 0.0
		for _, c := range chara {
			for _, g := range gacha {
				if g.Rarity == c.Rarity {
					total += g.Probability / float64(countMap[c.Rarity])
					break
				}
			}
			if total >= p {
				r = append(r, &Result{
					CharaId: c.Id,
					Name:    c.Name,
				})
				break
			}
		}
	}

	return r, nil
}

func (g gachaUseCase) Store(id string, results []*Result) error {

	for _, r := range results {
		posse := model.Possession{
			Id:      uuid.New().String(),
			UserId:  id,
			CharaId: r.CharaId,
			RegAt:   time.Now().Local(),
		}

		if err := g.gachaRepository.Store(&posse); err != nil {
			return errors.Wrap(err, "couldn't create")
		}
	}

	return nil
}
