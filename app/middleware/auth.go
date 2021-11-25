package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JWTCustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) GenerateTokenJWT(id uint) (string, error) {
	claims := JWTCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.SecretJWT))
	return token, nil
}

func GetClaimUser (c echo.Context) *JWTCustomClaims {
	user := c.Get("user").(*jwt.Token)
	return user.Claims.(*JWTCustomClaims)
}