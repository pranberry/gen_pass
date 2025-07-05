package password

import (
	"errors"
	"math/rand"
)


func GenPass(upper, lower, digits bool, pwSize int) (string, error) {

	if pwSize <= 0 {
		return "", errors.New("password length must be greater than 0")
	}
	if !upper && !lower && !digits {
		return "", errors.New("must select atleast one char-type")
	}

	// Generate the pool of characters for password
	charPool := make([]rune, 0)
	if upper {
		for i := 'A'; i <= 'Z'; i++ {
			charPool = append(charPool, i)
		}
	}
	if lower{
		for i := 'a'; i <= 'z'; i++ {
			charPool = append(charPool, i)
		}
	}
	if digits {
		for i := '0'; i <= '9'; i++ {
			charPool = append(charPool, i)
		}
	}

	password := make([]rune, pwSize)
	for i := range password{
		password[i] = charPool[rand.Intn(len(charPool))]
	}

	return string(password), nil
}