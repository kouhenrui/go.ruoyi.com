package token

import (
	"errors"
	"github.com/go-pay/gopay/pkg/jwt"
	"go.ruoyi.com/pkg/msg"
	"go.ruoyi.com/src/config"
	"go.ruoyi.com/src/config/dto"
	"time"
)

var hotkey = []byte(config.JWTKey)

type AllClaims struct {
	jwt.StandardClaims
	User dto.TokenClaims
}
type Token struct{}
type Tokentor interface {
	//AnalysyToken(c *gin.Context) bool
	ParseToken(tokenString string) dto.TokenClaims
	SignToken(infoClaims dto.TokenClaims, day time.Duration) (*dto.TokenAndExp, error)
}

func NewToken() *Token {
	return &Token{}
}

// 颁发token admin
func (t *Token) SignToken(infoClaims dto.TokenClaims, day time.Duration) (*dto.TokenAndExp, error) {
	expireTime := time.Now().Add(day) //7天过期时间
	claims := &AllClaims{
		User: infoClaims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "khr",  // 签名颁发者
			Subject:   "sign", //签名主题
		},
	}
	tp := &dto.TokenAndExp{}
	var err error
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tp.Token, err = token.SignedString(hotkey)
	if err != nil {
		return nil, errors.New(msg.MakeTokenError)
	}
	tp.ExpTime = expireTime.Format("2006-01-02 15:04:05")
	return tp, nil
}

//// 验证token
//func (t *Token) AnalysyToken(c *gin.Context) bool {
//	tokenString := c.GetHeader("Authorization")
//	if tokenString == "" {
//		return false
//	}
//	return true
//}

// 解析Token
func (t *Token) ParseToken(tokenString string) dto.TokenClaims {
	//解析token
	token, _ := jwt.ParseWithClaims(tokenString, &AllClaims{}, func(token *jwt.Token) (interface{}, error) {
		return hotkey, nil
	})
	user, _ := token.Claims.(*AllClaims)
	return user.User
}
