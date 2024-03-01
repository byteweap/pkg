package cryptox

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"

	"github.com/byteweap/pkg/convert"
)

// 加密
func EncryptRSA(publicKey, data string) (string, error) {
	public := convert.String2Bytes(publicKey)
	plaintext := convert.String2Bytes(data)
	//解密pem格式的公钥
	block, _ := pem.Decode(public)
	if block == nil {
		return "", errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	bs, err := rsa.EncryptPKCS1v15(rand.Reader, pub, plaintext)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bs), err
}

// 解密
func DecryptRSA(privateKey, data string) (string, error) {

	crypted, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	private := convert.String2Bytes(privateKey)
	//解密
	block, _ := pem.Decode(private)
	if block == nil {
		return "", errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	p, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	priv := p.(*rsa.PrivateKey)
	// 解密
	rs, err := rsa.DecryptPKCS1v15(rand.Reader, priv, crypted)
	if err != nil {
		return "", err
	}
	return convert.Bytes2String(rs), nil
}
