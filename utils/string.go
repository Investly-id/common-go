package utils

import (
	"math/rand"
	"time"
)

const CHARSET = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ"

const NUMSET = "0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandString(length int, isNumber bool) string {
	if length == 0 {
		length = 48
	}

	b := make([]byte, length)

	if isNumber {
		for i := range b {
			b[i] = NUMSET[seededRand.Intn(len(NUMSET))]
		}
	} else {
		for i := range b {
			b[i] = CHARSET[seededRand.Intn(len(CHARSET))]
		}
	}

	return string(b)
}
