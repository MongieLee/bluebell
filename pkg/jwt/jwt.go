package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

var jwtSecret = []byte("bluebell")

type CustomClaims struct {
	jwt.StandardClaims
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
}

// GenToken 生成JWT
func GenToken(userId int64, username string) (string, error) {
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Second).Unix(),
			Issuer:    string(jwtSecret),
		},
		UserId:   userId,
		Username: username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析Token
func ParseToken(token string) (*CustomClaims, error) {
	cc := new(CustomClaims)
	t, err := jwt.ParseWithClaims(token, cc, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if t.Valid {
		return cc, nil
	}
	return nil, errors.New("invalid token")
}
