package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	form "github.com/go-playground/form/v4"
	"github.com/rs/zerolog/log"
)

type User struct {
	Name   string `form:"name"`
	Age    uint   `form:"age"`
	Gender string `form:"gender"`
	Active bool   `form:"active"`
}

var (
	encoder *form.Encoder
)

const (
	cryptoKey = "example key 1234"
)

func init() {
	encoder = form.NewEncoder()
}

func main() {
	user := &User{
		Name:   "Pranc1ngPegasus",
		Age:    3,
		Gender: "Male",
		Active: true,
	}

	values, err := encoder.Encode(&user)
	if err != nil {
		log.Error().Stack().Err(err)
		return
	}

	for key := range values {
		value := values.Get(key)

		values.Set(key, encrypt(value))
	}

	log.Info().Msgf("%+v", values.Encode())
}

func encrypt(plain string) string {
	plaintext := []byte(plain)

	block, err := aes.NewCipher([]byte(cryptoKey))
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext)
}
