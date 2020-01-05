package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("GFDHTR&$%$%^#$GGFG5545gdgdf-90345")

// JWT 参考：http://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html
// 自定义 Payload（负载）
type Claims struct {
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

// 创建token
func GenerateToken(nickname string, expiration time.Duration) (string, error) {
	fmt.Println(nickname)
	fmt.Println(expiration)
	return "GFDHTR&$%$%^#$GGFG5545gdgdf-90345", nil
}

//
//// 解析 token
//func ParseToken(token string) (*Claims, error) {
//	// 解析token
//	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, keyFunc)
//	if tokenClaims != nil {
//		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
//			return claims, nil
//		}
//	}
//	return nil, err
//}
//
//// 返回key
//func keyFunc(token *jwt.Token) (interface{}, error) {
//	return jwtSecret, nil
//}
