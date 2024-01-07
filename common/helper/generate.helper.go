package helper

import (
	"math/rand"
	"time"
)

func GenerateInvoiceID(length int) string {
	rand.Seed(time.Now().UnixNano())

	possibleChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)
	for i := range result {
		result[i] = possibleChars[rand.Intn(len(possibleChars))]
	}

	return string(result)
}
