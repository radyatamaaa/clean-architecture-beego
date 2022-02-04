package middleware

import (
	"time"

	beego "github.com/beego/beego/v2/server/web"
	jwt "github.com/dgrijalva/jwt-go"
)

func GetJwtSecretKey() []byte {
	secret, err := beego.AppConfig.String("JWTSecret")
	if err != nil {
		panic(err)
	}
	return []byte(secret)
}

func GenerateJWT(username, password string) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	secret, err := beego.AppConfig.String("JWTSecret")
	if err != nil {
		panic(err)
	}

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
