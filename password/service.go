package password

import (
	"math/rand"
	"time"
)

const (
	lowerCase   = "abcdefghijklmnopqrstuvwxyz"
	upperCase   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers     = "0123456789"
	specialChar = "!@#$%^&*()_-+={}[/?]"
)

func GeneratePassword(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	passwd := ""
	for i := 0; i < length; i++ {
		randomInt := random.Int31n(4)
		switch randomInt {
		case 0:
			randomCharIndex := random.Intn(len(lowerCase))
			passwd += string(lowerCase[randomCharIndex])
		case 1:
			randomCharIndex := random.Intn(len(upperCase))
			passwd += string(upperCase[randomCharIndex])
		case 2:
			randomCharIndex := random.Intn(len(numbers))
			passwd += string(numbers[randomCharIndex])
		case 3:
			randomCharIndex := random.Intn(len(specialChar))
			passwd += string(specialChar[randomCharIndex])
		}
	}
	return passwd
}
