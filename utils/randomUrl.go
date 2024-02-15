package utils

import (
	"math/rand"
	"time"
)

func GetRandomString() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	characters := "abcdefghijklmnopqrstuvwxyz0123456789"

	var randomString string
	for i := 0; i < 9; i++ {
		index := rng.Intn(len(characters))
		randomString += string(characters[index])
	}
	return randomString
}
