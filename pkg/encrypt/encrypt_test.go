package encrypt

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	rsa := NewRSAEncrypt()
	//salt := rsa.MakeSalt()
	//log.Println(salt, "salt")
	//enpwd, err := rsa.EnPwdCode("123456", salt)
	//if err != nil {
	//	t.Errorf("err:%v", err)
	//}
	//log.Println(enpwd, "enpwd")
	//code, err := rsa.DePwdCode(enpwd, salt)
	//if err != nil {
	//	t.Errorf("err:%v", err)
	//}
	//t.Log(code, "code")

	//t.Log("打印aes方法")
	//rsa.LoadPrivateKey("private_key.pem")
	//rsa.LoadPublicKey("public_key.pem")
	//rsa.GenerateAESKey()
	////t.Log(rsa.privateKey)
	////t.Log(rsa.publicKey)
	//t.Log(rsa.aesKey)
	//
	//aes, i, err := rsa.EncryptAES([]byte("123456"))
	//if err != nil {
	//	t.Errorf("err:%v", err)
	//}
	//t.Log(aes, "aes")
	//t.Log(i, "i")
	//decryptAES, err := rsa.DecryptAES(aes, i)
	//if err != nil {
	//	t.Errorf("err:%v", err)
	//}
	//t.Log(string(decryptAES), "decryptAES")

	t.Log("打印rsa方法")
	p := []byte("123456")
	key, err := rsa.EncryptWithPublicKey(p)
	if err != nil {
		t.Errorf("err:%v", err)
	}
	t.Log(key, "key")
	privateKey, err := rsa.DecryptWithPrivateKey(key)
	if err != nil {
		t.Errorf("err:%v", err)
	}
	t.Log(privateKey, "privateKey")
	//rsa.EncryptAES()
	//
	//rsa.DecryptAES()
}
