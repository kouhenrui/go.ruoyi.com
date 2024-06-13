package captcha

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"go.ruoyi.com/src/config"
	"go.ruoyi.com/src/config/dto"
	"go.ruoyi.com/src/config/redis"
	"image/color"
)

type Captcha struct {
	Id      string
	Content string
	Answer  string
	store   base64Captcha.Store
}

type redisStore struct {
	redisClient redis.NewRedis
}

var redisc redis.NewRedis

type Captchator interface {
	CreateCaptcha() error
	VerifyCaptcha(capt dto.Captcha) bool
}

func NewCaptcha() *Captcha {
	store := &redisStore{
		redisClient: redisc,
	}
	return &Captcha{
		store:   store,
		Id:      "",
		Content: "",
		Answer:  "",
	}
}

/*
 * @MethodName CreateCaptcha
 * @Description 生成图片验证
 * @Author khr
 * @Date 2023/5/8 10:44
 */

func (c *Captcha) CreateCaptcha() error {
	//创建一个字符串类型的验证码驱动DriverString, DriverChinese :中文驱动
	driver := base64Captcha.NewDriverString(
		80, 240, 0, base64Captcha.OptionShowHollowLine,
		6, "1234567890qwertyuioplkjhgfdsazxcvbnm",
		&color.RGBA{R: 3, G: 102, B: 214, A: 125},
		base64Captcha.DefaultEmbeddedFonts,
		[]string{"wqy-microhei.ttc"},
	)
	//生成验证码
	createCaptcha := base64Captcha.NewCaptcha(driver, c.store)
	id, content, answer, err := createCaptcha.Generate()
	c.Id = id
	c.Content = content
	c.Answer = answer
	//newCaptcha.Id = id
	//newCaptcha.Content = content
	if err != nil {
		fmt.Println("生成有错:", err)
		return err
	}
	return nil

}

/*
 * @MethodName VerifyCaptcha
 * @Description 验证图片验证码
 * @Author khr
 * @Date 2023/5/8 10:45
 */

func (c *Captcha) VerifyCaptcha() bool {
	// id 验证码id
	// answer 需要校验的内容
	// clear 校验完是否清除
	return c.store.Verify(c.Id, c.Content, true) //c.store(c.Id, c.Content, true)

}
func (r *redisStore) Set(id string, value string) error {
	key := config.Captcha.Prefix + id
	return r.redisClient.SetRedis(key, []byte(value), config.Captcha.Expired) //Redis.Set(ctx, key, value, global.CaptchaExp).Err()
}
func (r *redisStore) Get(id string, clear bool) string {
	key := config.Captcha.Prefix + id
	val := r.redisClient.GetRedis(key)
	err := r.redisClient.DelRedis(key)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return val
}
func (r *redisStore) Verify(id, answer string, clear bool) bool {
	v := r.Get(id, clear)
	return v == answer
}
