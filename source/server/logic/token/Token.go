package token

import (
	"github.com/dgrijalva/jwt-go"
	"server/component/config"
)

type Token struct {
	Uid                string `json:"uid"`
	Nickname           string `json:"nickname"`
	jwt.StandardClaims        // jwt标准字段 有效期等
}

/**
 * @Description: 生成token
 */
func Sign(token *Token) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, token).SignedString([]byte(config.Config.GetString("jwt.key")))
}

// Parse 解析token
func Parse(sign string) (*Token, error) {
	token := &Token{}
	_, err := jwt.ParseWithClaims(sign, token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.GetString("jwt.key")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
