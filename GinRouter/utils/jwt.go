package utils

import (
	"GinRouter/pkg/setting"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"time"
)


var JwtSecret = []byte(setting.AppSetting.JwtSecret)


type Claims struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
	Password string `json:"password"`
	jwt.StandardClaims
}


//生成Token字符串
func GenerateToken(id int, username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	fmt.Print(id)
	claims := Claims{
		username,
		id,
		EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "https://go-admin/",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	return token, err
}


//解析Token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}


//校验Token
func CerifyToken(token string) bool{
	claims, err := ParseToken(token)
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			fmt.Println("Token超时")
			return false
		default:
			return false
		}
	}
	fmt.Println("校验成功", claims)
	return true
}


//通过ID获取用户信息
func GetIdFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)

			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return "0"
}