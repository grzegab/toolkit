package stringx

import (
	"crypto/rand"
	"math/big"
)

const (
	randomStringAll   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomStringSmall = "abcdefghijklmnopqrstuvwxyz"
)

// RandomString generates a random string of specified length using alphanumeric characters (a-z, A-Z, 0-9).
func RandomString(l int) string {
	return generateRandomString(l, randomStringAll)
}

// RandomSmallString generates a random string of specified length using lowercase letters (a-z).
func RandomSmallString(l int) string {
	return generateRandomString(l, randomStringSmall)
}

func generateRandomString(l int, charset string) string {
	if l <= 0 {
		return ""
	}

	b := make([]byte, l)

	n := big.NewInt(int64(len(charset)))
	for i := range b {
		num, err := rand.Int(rand.Reader, n)
		if err != nil {
			panic("crypto/rand failed: " + err.Error())
		}

		b[i] = charset[num.Int64()]
	}

	return string(b)
}
