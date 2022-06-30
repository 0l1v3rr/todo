package util

import "math/rand"

const chars string = "abcdefghijklmnopqrstuvwxyz0123456789"

func GenerateHash(length int) string {
	res := ""

	for i := 0; i <= length; i++ {
		res += string(chars[rand.Intn(len(chars)-1)])
	}

	return res
}
