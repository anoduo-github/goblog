package middleware

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//自定义一个字符串
var jwtkey = []byte("lee's blog")

//信息
type Claims struct {
	Role string //权限
	jwt.StandardClaims
}

//保存token
func SaveToken(role string) (string, error) {
	expireTime := time.Now().Add(2 * time.Hour)
	claims := &Claims{
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "your father", // 签名颁发者
			Subject:   "user token",  //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	return tokenString, err
}

//得到token信息
func GetTokenInfo(tokenString string) (string, error) {
	token, claims, err := ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("token无效")
	}
	return claims.Role, nil
}

//ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
