package usecase

import (
	"log"

	"github.com/pkg/errors"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"
	"github.com/shinnosuke-K/Tech-Train-CA-Go/infra/logger"
)

type Character struct {
	UserCharacterID string `json:"userCharacterID"`
	CharacterID     string `json:"characterID"`
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

	logger.Log.Info("[method:List] start")

	possCharas, err := c.characterUseCase.GetPossession(userId)
	if err != nil {
		log.Println(err)
		return nil, errors.New("you don't have any characters")
	}

	if len(possCharas) < 1 {
		return nil, errors.New("you don't have any characters")
	}

	charaIDs := make([]interface{}, 0, len(possCharas))
	for _, possChara := range possCharas {
		charaIDs = append(charaIDs, possChara.CharaID)
	}

	charaInfos, err := c.characterUseCase.GetCharacters(charaIDs)
	if err != nil {
		log.Println(err)
		return nil, errors.New("couldn't get characters")
	}

	charaList := make([]*Character, 0, len(possCharas))
	for _, chara := range charaInfos {
		for _, possChara := range possCharas {
			if possChara.CharaID == chara.ID {
				charaList = append(charaList, &Character{
					UserCharacterID: possChara.ID,
					CharacterID:     chara.ID,
					Name:            chara.Name,
				})
			}
		}
	}

	logger.Log.Info("[method:List] finished")
	return charaList, nil
}
