package library

import (
	"crypto/rand"
	"math/big"
)

// Code Source: https://github.com/atulsingh0/learning-golang/blob/main/freeCodeCamp/38-random-string-simplified.go
// License: Apache 2.0

const (
	letters         = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-?_"
	lettersAlphaNum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func RandomString(num int) string {

	// Creating an array having length <num>
	var randomStr []byte = make([]byte, num)

	for i := 0; i < num; i++ {
		// Generating an random index which varies from 0 to len(letters)
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		randomStr[i] = letters[idx.Int64()]
	}
	return string(randomStr)
}

func RandomStringAlphaNum(num int) string {

	// Creating an array having length <num>
	var randomStr []byte = make([]byte, num)

	for i := 0; i < num; i++ {
		// Generating an random index which varies from 0 to len(letters)
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(lettersAlphaNum))))
		randomStr[i] = lettersAlphaNum[idx.Int64()]
	}
	return string(randomStr)
}
