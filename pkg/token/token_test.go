package token

import (
	"go.ruoyi.com/src/config/dto"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	newtoken := NewToken()
	tc := &dto.TokenClaims{
		Id:    0,
		Name:  "张三",
		Phone: "123456",
		Email: "123@163.com",
		Role:  10,
	}
	token, err := newtoken.SignToken(*tc, time.Duration(10))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
	tcs := newtoken.ParseToken(token.Token)
	t.Log(tcs)
}
