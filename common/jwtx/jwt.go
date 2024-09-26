package jwtx

import (
	"github.com/golang-jwt/jwt"
)

func GetToken(secretKey string, iat int64, seconds int64, uid string, TokenType, nickName string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	claims["tokenType"] = TokenType
	claims["nickName"] = nickName
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
