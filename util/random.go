// to create a random test cases
package util

import (
	"math/rand"
	"strings"
	"time"
)

// init is called before main
func init() {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator
}

// randomInt returns a random integer between min and max, inclusive.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // rand.Int63n returns a non-negative pseudo-random 63-bit integer as an int64

}

// randomString returns a random string of length n.
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// randomOwner returns a random owner name.
func RandomOwner() string {
	return RandomString(6)
}

// randomMoney returns a random amount of money.
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// randomCurrency returns a random currency code.
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GBP", "INR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
