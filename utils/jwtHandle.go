package utils

import (
	"github.com/dgrijalva/jwt-go"
	"pcrweb/config"
	"pcrweb/model"
	"time"
)

func GenerateToken(uuid string, account string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.JwtClaims{
		Uuid:    uuid,
		Account: account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*24*3, // 天有效期
		},
	})
	return token.SignedString([]byte(config.SecretKey))
}

func VerifyToken(tokenStr string, SecretKey []byte) (uuid string, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claim := token.Claims.(jwt.MapClaims)
	uuid = claim["Uuid"].(string)
	return uuid, err
}
