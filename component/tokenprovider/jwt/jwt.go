package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var SECERECT_KEY = "123456789_QWERTYUIIOP"

func GenerateJWT(id, role string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	var mySigningKey = []byte(SECERECT_KEY)
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetClaimsFromToken(rawToken string) map[string]string {

	jwtClaims := jwt.MapClaims{}
	jwt.ParseWithClaims(rawToken, jwtClaims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECERECT_KEY), nil
	})

	result := make(map[string]string)
	for key, value := range jwtClaims {
		valStr := fmt.Sprintf("%v", value)
		result[key] = valStr
	}
	return result
}
