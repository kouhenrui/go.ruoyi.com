package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	rsaRand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

type RSAEncrypt struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	aesKey     []byte
}
type RSAEncryptor interface {
	MakeSalt() string                                        //16位密码加密盐
	Rand6String() string                                     //随机6位密钥
	RandAllString() string                                   //随机16位密钥
	EnPwdCode(Pwd string, pwdKey string) (string, error)     //加密
	DePwdCode(pwd string, pwdKey string) (string, error)     //解密
	EncryptWithPublicKey(plaintext []byte) (string, error)   //公钥加密
	DecryptWithPrivateKey(ciphertext string) (string, error) //私钥解密
	EncryptAES(plaintext []byte) ([]byte, []byte, error)     //aes256加密
	DecryptAES(ciphertext []byte, iv []byte) (string, error) //aes256解密

}

var strByte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*-=+")
var strByteLen = len(strByte)

func NewRSAEncrypt() *RSAEncrypt {
	privateKey, _ := LoadPrivateKey("private_key.pem")
	publicKey, _ := LoadPublicKey("public_key.pem")
	aesKey, _ := GenerateAESKey()
	return &RSAEncrypt{
		privateKey: privateKey,
		publicKey:  publicKey,
		aesKey:     aesKey,
	}
}

// 生成12位随机字符串 加两位==
func (r *RSAEncrypt) RandAllString() string {

	bytes := make([]byte, 14)
	randString := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 14; i++ {
		bytes[i] = strByte[randString.Intn(strByteLen)]
	}

	return string(bytes) + "=="
	//str := strings.Builder{}
	//length := len(CHARS)
	//for i := 0; i < 14; i++ {
	//	l := CHARS[rand.Intn(length)]
	//	str.WriteString(l)
	//}
	//return str.String() + "=="
}

/**
 * @Author Khr
 * @Description redis缓存名称
 * @Date 14:08 2023/8/29
 * @Param
 * @return
 **/
func (r *RSAEncrypt) Rand6String() string {
	bytes := make([]byte, 4)
	randString := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 4; i++ {
		bytes[i] = strByte[randString.Intn(strByteLen)]
	}
	return string(bytes) + "=="
}

/**
 * @Author Khr
 * @Description 生成16位密钥
 * @Date 10:52 2023/8/29
 * @Param
 * @return
 **/
func (r *RSAEncrypt) MakeSalt() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 14)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b) + "=="

}

// PKCS7 填充模式
func (r *RSAEncrypt) pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 填充的反向操作，删除填充字符串
func (r *RSAEncrypt) pkcs7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}

// 实现加密
func (r *RSAEncrypt) aesEcrypt(origData []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err, "///////")
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = r.pkcs7Padding(origData, blockSize)
	//采用AES加密方法中CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	//执行加密
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 实现解密
func (r *RSAEncrypt) aesDeCrypt(cypted []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cypted))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(origData, cypted)
	//去除填充字符串
	origData, err = r.pkcs7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}

// 加密base64
func (r *RSAEncrypt) EnPwdCode(Pwd string, pwdKey string) (string, error) {
	pwd := []byte(Pwd)
	PwdKey := []byte(pwdKey)
	result, err := r.aesEcrypt(pwd, PwdKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), err
}

// 解密
func (r *RSAEncrypt) DePwdCode(pwd string, pwdKey string) (string, error) {
	PwdKey := []byte(pwdKey)
	//解密base64字符串
	pwdByte, _ := base64.StdEncoding.DecodeString(pwd)
	ecpwsd, err := r.aesDeCrypt(pwdByte, PwdKey)

	//fmt.Println("字符格式密码", ecpwsd)
	//fmt.Println("密码：", string(ecpwsd))
	if err != nil {
		return "密码解析错误", err
	}
	return string(ecpwsd), err
}

/**
 * @Author Khr
 * @Description //TODO 制作rsa加密和解密密钥文件
 * @Date 15:16 2024/5/28
 * @Param
 * @return
 **/
func (r *RSAEncrypt) MakeRSA() (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rsaRand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	// 保存私钥
	privateKeyFile, err := os.Create("private_key.pem")
	if err != nil {
		return nil, err
	}
	defer privateKeyFile.Close()
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	privateKeyFile.Write(privateKeyPEM)

	// 保存公钥
	publicKey := &privateKey.PublicKey
	publicKeyFile, err := os.Create("public_key.pem")
	if err != nil {
		return nil, err
	}
	defer publicKeyFile.Close()
	publicKeyPEM, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	pem.Encode(publicKeyFile, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyPEM,
	})

	return privateKey, nil
}

// LoadPrivateKey 从文件加载私钥
func LoadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// LoadPublicKey 从文件加载公钥
func LoadPublicKey(filename string) (*rsa.PublicKey, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not RSA public key")
	}
	return publicKey, nil
}

// DecryptWithPrivateKey 使用私钥解密数据
func (r *RSAEncrypt) DecryptWithPrivateKey(ciphertext string) ([]byte, error) {
	label := []byte("")
	hash := sha256.New()
	cipherBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	plaintext, err := rsa.DecryptOAEP(hash, rsaRand.Reader, r.privateKey, cipherBytes, label)
	if err != nil {
		return nil, err
	}
	//r.aesKey = plaintext
	return plaintext, nil
}

// EncryptWithPublicKey 使用公钥加密数据
func (r *RSAEncrypt) EncryptWithPublicKey(plaintext []byte) (string, error) {
	label := []byte("")
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rsaRand.Reader, r.publicKey, plaintext, label)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// 使用aes加密
func (r *RSAEncrypt) EncryptAES(plaintext []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher([]byte(r.aesKey))
	if err != nil {
		return nil, nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return nil, nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, iv, nil
}

// aes解密
func (r *RSAEncrypt) DecryptAES(ciphertext []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(r.aesKey)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	plaintext := make([]byte, len(ciphertext)-aes.BlockSize)
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])

	return string(plaintext), nil
}

// 生成aes秘钥
func GenerateAESKey() ([]byte, error) {
	key := make([]byte, 32) // 256位 = 32字节
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}
