package captcha

import "testing"

func TestCaptchar(t *testing.T) {
	captchar := NewCaptchar()
	err := captchar.CreateCaptchar()
	if err != nil {
		return
	}
	t.Log(captchar.Id, captchar.Answer)

	captchar.VerifyCaptchar(captchar.Id, captchar.Answer)
}
