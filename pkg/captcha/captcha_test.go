package captcha

import "testing"

func TestCaptcha(t *testing.T) {

	capcha := NewCaptcha()
	capcha.CreateCaptcha()

	t.Fatal(capcha.Answer, capcha.Id)

	r := capcha.VerifyCaptcha()
	t.Fatal(r, "r")
}
