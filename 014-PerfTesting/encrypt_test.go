package main

import (
	"encoding/hex"
	"testing"
)

var (
	testPlain  = []byte("Load your secret key from a safe place and reuse it across multiple\nNewCipher calls. (Obviously don't use this example key for anything\nreal.) If you want to convert a passphrase to a key, use a suitable\n package like bcrypt or scrypt.")
	testKey, _ = hex.DecodeString("6368616e676520746869732070617373")
)

func BenchmarkEncryptAES(b *testing.B) {
	// If I needed to do setup, I could do it here
	// then call b.StartTimer() to start the timer.
	for n := 0; n < b.N; n++ {
		encryptAES(testPlain, testKey)
	}
}

func BenchmarkEncryptRC4(b *testing.B) {

	for n := 0; n < b.N; n++ {
		encryptRC4(testPlain, testKey)
	}
}
