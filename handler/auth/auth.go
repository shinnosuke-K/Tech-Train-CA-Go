package auth

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	Id  string
	Nbf string
	Iat string
}

func CreateJwt(claims jwt.MapClaims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func ParseToken(token string) (*Auth, error) {

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s", "Unexpected signing method")
		}

		keyData, err := ioutil.ReadFile(os.Getenv("KEY_PATH"))
		if err != nil {
			return nil, err
		}
		return keyData, nil
	})

	// 要修正
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse authorization")
	}

	t, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to parse authorization")
	}

	id, ok := t["user_id"].(string)
	if !ok {
		return nil, errors.New("failed to parse authorization")
	}

	nbf, ok := t["nbf"].(string)
	if !ok {
		return nil, errors.New("failed to parse authorization")
	}

	iat, ok := t["iat"].(string)
	if !ok {
		return nil, errors.New("failed to parse authorization")
	}

	return &Auth{
		Id:  id,
		Nbf: nbf,
		Iat: iat,
	}, nil
}
