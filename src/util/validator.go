package util

import (
	"benings/model"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
)

func ValidateEmail(email string) bool {
	Re := regexp.MustCompile("((\\w+\\.)*\\w+)@(\\w+\\.)+(com|kr|net|us|info|biz|id)")
	return Re.MatchString(email)
}

func TokenUsingJWT(userName string) model.JwtResponse {
	var response model.JwtResponse
	var jwtKey = []byte(userName)
	currentTime := time.Now()
	expirationTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+1, 00, 00, 01, 000000000, time.UTC)
	tokens := &model.Token{
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokens)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Internal server error"
		response.Token = ""
	} else {
		response.OutError = 0
		response.OutMessage = "succsess"
		response.Token = tokenString
	}
	return response
}

func Encrypt(src string, key []byte, initialVector string) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(initialVector))
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

func Decrypt(crypt []byte, key []byte, initialVector string) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if len(crypt) == 0 {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCDecrypter(block, []byte(initialVector))
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return PKCS5Trimming(decrypted)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
