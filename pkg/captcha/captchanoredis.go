package captcha

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
	"sync"
)

type Captchar struct {
	Id      string
	Content string
	Answer  string
	store   base64Captcha.Store
}

type memoryStore struct {
	mu    sync.RWMutex
	store map[string]string
}

func newMemoryStore() *memoryStore {
	return &memoryStore{
		store: make(map[string]string),
	}
}

func (m *memoryStore) Set(id string, value string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.store[id] = value
	return nil
}

func (m *memoryStore) Get(id string, clear bool) string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, exists := m.store[id]
	if !exists {
		return ""
	}
	if clear {
		delete(m.store, id)
	}
	return value
}

func (m *memoryStore) Verify(id, answer string, clear bool) bool {
	v := m.Get(id, clear)
	return v == answer
}

func NewCaptchar() *Captchar {
	return &Captchar{
		store: newMemoryStore(),
	}
}

func (c *Captchar) CreateCaptchar() error {
	// 定义一个字符串类型的验证码生成器
	driver := base64Captcha.NewDriverString(
		80, 240, 0, base64Captcha.OptionShowHollowLine, 6, "1234567890qwertyuioplkjhgfdsazxcvbnm",
		&color.RGBA{R: 3, G: 102, B: 214, A: 125},
		base64Captcha.DefaultEmbeddedFonts, nil,
	)

	// 创建验证码
	captcha := base64Captcha.NewCaptcha(driver, c.store)
	id, content, answer, err := captcha.Generate()
	if err != nil {
		return err
	}

	c.Id = id
	c.Content = content
	c.Answer = answer
	return nil
}

func (c *Captchar) VerifyCaptchar(id, answer string) bool {
	return c.store.Verify(id, answer, true)
}
