package pwd

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	buf, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func Verify(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
