package usecase

import (
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"
)

type Result struct {
	CharaId int
	Name    string
}

type GachaUseCase interface {
	Draw(times int) ([]*Result, error)
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

	rand.Seed(time.Now().Unix())

	r := make([]*Result, 0, times)
	for n := 0; n < times; n++ {
		p := rand.Float64() * 100
		total := 0.0
		for _, c := range chara {
			for _, g := range gacha {
				if g.Rarity == c.Rarity {
					total += g.Probability / float64(len(chara))
					break
				}
			}
			if total >= p {
				r = append(r, &Result{
					CharaId: c.CharaId,
					Name:    c.CharaName,
				})
				break
			}
		}
	}

	return r, nil
}
