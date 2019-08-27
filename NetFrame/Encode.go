package NetFrame

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	b64 "encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

var MyKey []byte = []byte("1111222233334444")

//encode 输出消息体前面一段bytes
//字段1：type		int32
//字段2：command	int32
//字段3：messages	[]byte		序列化后的消息体
//长度字段：	字段1之前增加int32字段表示要后面要发送的字段长度

type Encode struct {
	len      int32
	thetype  int32
	command  int32
	head     []byte
	writePos int32
}

func NewEncode(inputLen, inputType, inputCommand int32) *Encode {
	e := &Encode{
		len:     inputLen,
		thetype: inputType,
		command: inputCommand,
	}
	return e
}
func (enc *Encode) WriteInt32(num int32) {

	buff := bytes.NewBuffer([]byte{})
	binary.Write(buff, binary.LittleEndian, num)
	copy(enc.head[enc.writePos:], buff.Bytes())
	enc.writePos += int32(binary.Size(num))
}

func (enc *Encode) Write() {
	enc.head = make([]byte, 12)
	enc.writePos = 0
	enc.WriteInt32(enc.len)
	enc.WriteInt32(enc.thetype)
	enc.WriteInt32(enc.command)
}
func (enc *Encode) GetBytes() (head []byte) {
	return enc.head
}

func WriteMessage(thetype int32, command int32, message []byte, messageLen int) (*bytes.Buffer, error) {
	encode := NewEncode(int32(8+messageLen), thetype, command)
	encode.Write()
	var buffer bytes.Buffer
	buffer.Write(encode.GetBytes())
	//增加加密部分
	/*
		en, err := AesEncrypt(message, MyKey)
		if err != nil {
			return nil, err
		} else {
			buffer.Write(en)
		}*/
	Encrypt(string(message), "1111222233334444")
	return &buffer, nil
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

const (
	ivDefValue = "0102030405060708"
)

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

// Encrypts text with the passphrase
func Encrypt(text string, passphrase string) string {
	salt := make([]byte, 8)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic(err.Error())
	}

	key, iv := __DeriveKeyAndIv(passphrase, string(salt))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	pad := __PKCS5Padding([]byte(text), block.BlockSize())
	ecb := cipher.NewCBCEncrypter(block, []byte(iv))
	encrypted := make([]byte, len(pad))
	ecb.CryptBlocks(encrypted, pad)

	return b64.StdEncoding.EncodeToString([]byte("Salted__" + string(salt) + string(encrypted)))
}

// Decrypts encrypted text with the passphrase
func Decrypt(encrypted string, passphrase string) string {
	ct, _ := b64.StdEncoding.DecodeString(encrypted)
	if len(ct) < 16 || string(ct[:8]) != "Salted__" {
		return ""
	}

	salt := ct[8:16]
	ct = ct[16:]
	key, iv := __DeriveKeyAndIv(passphrase, string(salt))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	cbc := cipher.NewCBCDecrypter(block, []byte(iv))
	dst := make([]byte, len(ct))
	cbc.CryptBlocks(dst, ct)

	return string(__PKCS5Trimming(dst))
}

func __PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func __PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func __DeriveKeyAndIv(passphrase string, salt string) (string, string) {
	salted := ""
	dI := ""

	for len(salted) < 48 {
		md := md5.New()
		md.Write([]byte(dI + passphrase + salt))
		dM := md.Sum(nil)
		dI = string(dM[:16])
		salted = salted + dI
	}

	key := salted[0:32]
	iv := salted[32:48]

	return key, iv
}
