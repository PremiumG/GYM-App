package utils

import "golang.org/x/crypto/bcrypt"

//This is here as a previous test on storing hashed data and using it to login. test was ok and converted back to simple pass check. for simplicity

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
