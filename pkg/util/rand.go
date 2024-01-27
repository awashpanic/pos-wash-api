package util

import (
	"crypto/rand"
	"log"
	"math"
	"math/big"
	math_rand "math/rand"
	"strings"
)

const (
	CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

func GenerateRandomString(length int, upper bool) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = CHARS[math_rand.Intn(len(CHARS))]
	}

	if upper {
		return strings.ToUpper(string(b))
	}

	return string(b)
}

func GenerateRandomNumber(length int) (n int, err error) {
	min := int(math.Pow10(length - 1))
	max := int64(int(math.Pow10(length)) - 1)

	num, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Printf("[error-random-number]: %v\n", err)
		return 0, err
	}

	n = int(num.Int64())

	// Handling integers between 0, 10^(n-1) .. for n=4, handling cases between (0, 999)
	if n <= min {
		n += min
	}

	return n, nil
}
