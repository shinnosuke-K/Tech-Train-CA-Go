package usecase

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"
)

type Character struct {
	UserCharacterId string `json:"userCharacterID"`
	CharacterId     string `json:"characterID"`
	Name            string `json:"name"`
}

type CharacterUseCase interface {
	List(userId string) ([]*Character, error)
}

type characterUseCase struct {
	characterUseCase repository.CharacterRepository
}

func NewCharaUseCase(cg repository.CharacterRepository) CharacterUseCase {
	return &characterUseCase{
		characterUseCase: cg,
	}
}

func (c characterUseCase) List(userId string) ([]*Character, error) {

	possCharas, err := c.characterUseCase.GetPossession(userId)
	if err != nil {
		log.Println(err)
		return nil, errors.New("you don't have any characters")
	}

	// N+1問題
	charaList := make([]*Character, 0)
	for _, p := range possCharas {
		c, err := c.characterUseCase.GetCharacter(p.CharaId)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("not exits monster id = %s", p.CharaId)
		}

		charaList = append(charaList, &Character{
			UserCharacterId: p.PosseId,
			CharacterId:     c.CharaId,
			Name:            c.CharaName,
		})
	}

	return charaList, nil
}
