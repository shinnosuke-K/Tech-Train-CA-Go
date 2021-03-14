package auth

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func CreateJwtToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	keyData, err := ioutil.ReadFile(os.Getenv("KEY_PATH"))
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(keyData)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func rsaPublicKyeFunc() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}

		keyData, err := ioutil.ReadFile(os.Getenv("KEY_PATH"))
		if err != nil {
			return nil, err
		}
		return keyData, nil
	}
}

func Validate(accessToken string) error {

	token, err := jwt.Parse(accessToken, rsaPublicKyeFunc())
	if err != nil {
		return errors.Wrap(err, "failed to parse authorization")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("cannot parse claims")
	}

	if err := validateId(claims); err != nil {
		return errors.WithStack(err)
	}

	if err := validateIat(claims); err != nil {
		return errors.WithStack(err)
	}

	if err := validateNbf(claims); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func validateId(claims jwt.MapClaims) error {
	tokenId, ok := claims["user_id"]
	if !ok {
		return errors.New("cannot parse user_id")
	}

	_, ok = tokenId.(string)
	if !ok {
		return errors.New("cannot parse user_id")
	}
	return nil
}

func validateNbf(claims jwt.MapClaims) error {
	tokenNbf, ok := claims["nbf"]
	if !ok {
		return errors.New("cannot parse nbf")
	}

	_, ok = tokenNbf.(string)
	if !ok {
		return errors.New("cannot parse nbf")
	}
	return nil
}

func validateIat(claims jwt.MapClaims) error {
	tokenIat, ok := claims["iat"]
	if !ok {
		return errors.New("cannot parse iat")
	}

	_, ok = tokenIat.(string)
	if !ok {
		return errors.New("cannot parse iat")
	}
	return nil
}

func Get(header http.Header, key string) (string, error) {
	token, err := jwt.Parse(header.Get("x-token"), rsaPublicKyeFunc())
	if err != nil {
		return "", errors.WithStack(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to convert jwt.MapClaims")
	}

	value, ok := claims[key].(string)
	if !ok {
		return "", errors.Errorf("not exist key= %s in claims", key)
	}

	return value, nil
}
