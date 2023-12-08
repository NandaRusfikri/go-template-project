package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"go-template-project/dto"
	"os"
	"strings"
	"time"
)

type MetaToken struct {
	ID            string
	Email         string
	ExpiredAt     time.Time
	Authorization bool
}

type AccessToken struct {
	Claims MetaToken
}

func Sign(claims dto.Claims, JwtSecretKey string) (string, error) {

	jwtSecretKey := JwtSecretKey

	//claims := jwt.MapClaims{}
	//claims["exp"] = expiredAt
	//claims["authorization"] = true
	//
	//for i, v := range Data {
	//	claims[i] = v
	//}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(jwtSecretKey))

	if err != nil {
		logrus.Error(err.Error())
		return accessToken, err
	}

	return accessToken, nil
}

func VerifyTokenHeader(ctx *gin.Context) (*jwt.Token, error) {
	tokenHeader := ctx.GetHeader("Authorization")
	accessToken := strings.SplitAfter(tokenHeader, "Bearer")[1]
	jwtSecretKey := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func Auth(ctx *gin.Context) {
	tokenHeader := ctx.GetHeader("Authorization")
	if tokenHeader == "" {
		ctx.JSON(401, gin.H{
			"status_code": 401,
			"message":     "Unauthorized",
		})
		ctx.Abort()
		return
	}

	Secret := os.Getenv("JWT_SECRET")
	_, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(Secret), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		ctx.JSON(401, gin.H{
			"status_code": 401,
			"message":     err.Error(),
		})
		ctx.Abort()
		return
	}

}

func VerifyToken(accessToken, SecretPublicKeyEnvName string) (*jwt.Token, error) {
	jwtSecretKey := os.Getenv(SecretPublicKeyEnvName)

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}
