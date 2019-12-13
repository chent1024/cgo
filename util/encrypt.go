package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"errors"
)

// 3DES加密
func DesEncryptStr(src []byte, key, iv string) (string, error) {
	block, err := des.NewTripleDESCipher([]byte(key))
	if err != nil {
		return "", err
	}

	src = pkcs7Padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	encode := make([]byte, len(src))
	blockMode.CryptBlocks(encode, src)
	dst := base64.StdEncoding.EncodeToString(encode)
	return dst, nil
}

// 3DES解密
func DesDecryptStr(src, key, iv string) ([]byte, error) {
	block, err := des.NewTripleDESCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	srcByte, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	decode := make([]byte, len(srcByte))
	blockMode.CryptBlocks(decode, srcByte)
	decode = pkcs7UnPadding(decode)
	return decode, nil
}

// 解密微信加密数据
func DecryptEncryptedData(encryptedData, sessionKey, iv string) ([]byte, error) {
	if len(sessionKey) != 24 {
		return nil, errors.New("sessionKey length is error")
	}

	aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}

	if len(iv) != 24 {
		return nil, err
	}

	aesIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	aesCipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}

	aesPlantText := make([]byte, len(aesCipherText))
	aesBlock, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(aesBlock, aesIV)
	mode.CryptBlocks(aesPlantText, aesCipherText)
	aesPlantText = pkcs7UnPadding(aesPlantText)
	return aesPlantText, nil
}

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding return unpadding []Byte plantText
func pkcs7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	if length > 0 {
		unPadding := int(plantText[length-1])
		if length < unPadding {
			return nil
		}

		return plantText[:(length - unPadding)]
	}
	return plantText
}
