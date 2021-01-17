package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/google/uuid"
)

func CreateUserId() string {
	userId := uuid.Must(uuid.NewRandom())
	return strings.ReplaceAll(userId.String(), "-", "")
}

func GetJSTTime() time.Time {
	timeUTC := time.Now().UTC().Local()
	return timeUTC
}

func ParsedJWTToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s", "Unexpected signing method")

		} else {
			keyData, err := ioutil.ReadFile(os.Getenv("KEY_PATH"))
			if err != nil {
				return nil, err
			}
			return keyData, nil
		}
	})
}
