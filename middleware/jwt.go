package middleware

import (
	"encoding/json"
	"fmt"

	// "time"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	jwt "github.com/dgrijalva/jwt-go"
)

func GetJwtSecretKey() []byte {
	secret, err := beego.AppConfig.String("JWTSecret")
	if err != nil {
		panic(err)
	}
	return []byte(secret)
}

// func GenerateJWT(username, password string) (string, error) {
// 	token := jwt.New(jwt.SigningMethodES256)

// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["authorized"] = true
// 	claims["username"] = username
// 	claims["password"] = password
// 	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

// 	secret, err := beego.AppConfig.String("JWTSecret")
// 	if err != nil {
// 		panic(err)
// 	}

// 	tokenString, err := token.SignedString(secret)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

func JWTAuth(ctx *context.Context) {
	var uri string = ctx.Input.URI()
	if uri == "/api/v1/login" {
		return
	}

	if ctx.Input.Header("Authorization") == "" {
		ctx.Output.SetStatus(403)
		resBody, err := json.Marshal(map[string]string{"code": "403", "message": "Not Allowed"})
		ctx.Output.Body(resBody)
		if err != nil {
			panic(err)
		}
	}

	var tokenString string = ctx.Input.Header("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		secret := GetJwtSecretKey()

		return secret, nil
	})

	if err != nil {
		ctx.Output.SetStatus(403)
		var responseBody = map[string]string{"code": "403", "message": err.Error()}
		resBytes, err := json.Marshal(responseBody)
		ctx.Output.Body(resBytes)
		if err != nil {
			panic(err)
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && claims != nil {
		return
	} else {
		ctx.Output.SetStatus(403)
		resBody, err := json.Marshal(map[string]string{"code": "403", "message": ctx.Input.Header("Authorization")})
		ctx.Output.Body(resBody)
		if err != nil {
			panic(err)
		}
	}
}
