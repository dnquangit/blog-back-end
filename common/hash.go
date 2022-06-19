package common

import "golang.org/x/crypto/bcrypt"

func HashString(input string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), 5)
	return string(bytes), err
}
