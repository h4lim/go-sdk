package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(uniqueID string, username string, timeout int) *string {

	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(timeout)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["client"] = username
	sign.Claims = claims

	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		log.Errorf(uniqueID, "Error occurred ", err.Error())
		return nil
	}

	return &token
}
