package utils

import (
	"boarderbackend/pkgs/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtSecret = []byte(setting.App.JwtSecret)

type Claims struct {
	Username string `json:username`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Minute * 100)

	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer: "boarder",
			ExpiresAt: expireTime.Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)

	return token, err
}

func UpdateToken(token string) (string, error){
	claims, err := ParseToken(token)
	if err != nil {
		return "", err
	}
	return GenerateToken(claims.Username)
}

func ParseToken(token string) (*Claims, error){
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if  tokenClaims != nil  {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

