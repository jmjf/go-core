package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	"io"
)

var (
	plain        = []byte("Load your secret key from a safe place and reuse it across multiple\nNewCipher calls. (Obviously don't use this example key for anything\nreal.) If you want to convert a passphrase to a key, use a suitable\n package like bcrypt or scrypt.")
	masterkey, _ = hex.DecodeString("6368616e676520746869732070617373")
)

// Example code based on standard library examples -- NOT SECURE

func encryptAES(plaintext []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
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
	return ciphertext
}

func encryptRC4(plaintext []byte, key []byte) []byte {
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, len(plaintext))
	cipher.XORKeyStream(ciphertext, plaintext)
	return ciphertext
}

func main() {
	fmt.Printf("AES: %x\n\n", encryptAES(plain, masterkey))
	fmt.Printf("RC4: %x\n\n", encryptRC4(plain, masterkey))
}
