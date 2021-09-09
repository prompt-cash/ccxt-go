package randstr

import (
	"math/rand"
	"time"
)

type RandomStringGenerator struct {
	seededRand *rand.Rand
	charset    string
}

// Creates a new random string generator from the given charset.
// If charset is an empty string the alphabet & numbers will be used.
func NewGenerator(charset string) *RandomStringGenerator {
	if len(charset) == 0 {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}
	return &RandomStringGenerator{
		seededRand: rand.New(rand.NewSource(time.Now().UnixNano())),
		charset:    charset,
	}
}

// Returns a random string of the requested length.
func (gen *RandomStringGenerator) GetString(length int) string {
	b := make([]byte, length)
	bound := len(gen.charset)
	for i := range b {
		b[i] = gen.charset[gen.seededRand.Intn(bound)]
	}
	return string(b)
}
