package jwtToken

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	applicationName      = os.Getenv("APP_NAME")
	jwtSignatureKey      = []byte(os.Getenv("APP_SECRET_KEY_JWT"))
	jwtSigningMethod     = jwt.SigningMethodHS256
	loginExpiredDuration = time.Duration(60*2) * time.Minute // 2 hours
)

type MyClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

// GenerateToken - generate token
func GenerateToken(userID int) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    applicationName,
			ExpiresAt: time.Now().Add(loginExpiredDuration).Unix(),
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	stringToken, err := token.SignedString(jwtSignatureKey)
	if err != nil {
		return "", errors.New("Generate Token Jwt err = " + err.Error())
	}

	return stringToken, nil
}
