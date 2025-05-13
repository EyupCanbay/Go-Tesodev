package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "qwertyuiopasdfghjklzxcvbnm"

// for true randomness
func init() {
	rand.Seed(time.Now().UnixNano())
}

// random intager generates a random between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandonStirng generates a random sittring of length
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomMoney generates a random amount of money
func RandomPrice() int64 {
	return RandomInt(0, 1000)
}
