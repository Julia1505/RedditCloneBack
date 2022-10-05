package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdef0123456789"

var randSeed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenarateId(length int) string {
	idString := make([]byte, length)
	for i := range idString {
		idString[i] = charset[randSeed.Intn(len(charset)-1)]
	}
	return string(idString)
}
