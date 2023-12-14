package helpers

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	ePass, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		return "", err
	}
	return string(ePass), nil
}

func ValidatePassword(storedPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	return err == nil
}
