package utils

import (
	"fmt"
	"math/rand"
	"strconv"
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

func ThousandFormat(n int64, separator byte) string {
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits--
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = separator
		}
	}
}

func BilanganFormat(total int64, short bool) string {

	var newTotal = float64(total)

	if total <= 999999 {
		return fmt.Sprintf("%s", ThousandFormat(total, '.'))
	}

	if !short {
		if total > 999999 && total <= 999999999 {
			return fmt.Sprintf("%s Juta", strconv.FormatFloat(newTotal/1000000, 'f', -1, 64))
		}

		if total > 999999999 && total <= 999999999999 {
			return fmt.Sprintf("%sf Milyar", strconv.FormatFloat(newTotal/1000000000, 'f', -1, 64))
		}
	}

	if total > 999999 && total <= 999999999 {
		return fmt.Sprintf("%s Jt", strconv.FormatFloat(newTotal/1000000, 'f', -1, 64))
	}

	if total > 999999999 && total <= 999999999999 {
		return fmt.Sprintf("%sf M", strconv.FormatFloat(newTotal/1000000000, 'f', -1, 64))
	}

	return ""
}
