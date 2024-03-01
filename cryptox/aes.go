package cryptox

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/byteweap/pkg/convert"
)

// AES加密 - CBC模式
func EncryptAES(key, data string) (string, error) {
	k := convert.String2Bytes(key)
	plaintext := convert.String2Bytes(data)

	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	plaintext = pkcs5Padding(plaintext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	crypted := make([]byte, len(plaintext))
	blockMode.CryptBlocks(crypted, plaintext)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

// AES解密 - CBC模式
func DecryptAes(key string, data string) (string, error) {

	k := convert.String2Bytes(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	crypted, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pkcs5UnPadding(origData)
	return convert.Bytes2String(origData), nil
}

// 补码
func pkcs5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

// 去码
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
