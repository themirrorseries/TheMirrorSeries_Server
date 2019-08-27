package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

const (
	ivDefValue = "0102030405060708"
)

var mykey []byte = []byte("1111222233334444")
var mytext []byte = []byte("aaaabbbbccccddddd")

func main() {
	fmt.Println(mytext)
	en, _ := AesEncrypt(mytext, mykey)
	fmt.Println(en)
	fmt.Println(AesDecrypt(en, mykey))

}

func AesEncrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("invalid decrypt key")
	}
	blockSize := block.BlockSize()
	plaintext = PKCS5Padding(plaintext, blockSize)
	iv := []byte(ivDefValue)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	blockMode.CryptBlocks(ciphertext, plaintext)
	fmt.Println("aestext", plaintext)
	return ciphertext, nil
}

func AesDecrypt(ciphertext []byte, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("invalid decrypt key")
	}

	blockSize := block.BlockSize()

	if len(ciphertext) < blockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := []byte(ivDefValue)
	if len(ciphertext)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	blockModel := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plaintext, ciphertext)
	plaintext = PKCS5UnPadding(plaintext)

	return plaintext, nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
