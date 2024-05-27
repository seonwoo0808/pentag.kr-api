package utils

import "math/rand"

func GenerateRandomString(length int) string {
	var result string
	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < length; i++ {
		result += string(charset[RandomInt(0, len(charset))])
	}
	return result
}

func RandomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}


