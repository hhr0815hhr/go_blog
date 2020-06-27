package common

import (
	"blog/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("my secret key:Eassy")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user *model.User) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "eassy",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey, nil
	})
	return token, claims, err
}
