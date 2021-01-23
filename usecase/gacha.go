package usecase

import "github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"

type Result struct {
	Results []struct {
		CharaId int
		Name    string
	}
}

type GachaUseCase interface {
	Draw(times int) (Result, error)
}

type gachaUseCase struct {
	gachaRepository repository.GachaRepository
}

func NewGachaUseCase(ug repository.GachaRepository) GachaUseCase {
	return &gachaUseCase{
		gachaRepository: ug,
	}
}

func (g gachaUseCase) Draw(times int) (Result, error) {
	panic("implement me")
}
