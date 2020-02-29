package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword is a helper function to hash passwords easier
func HashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePassword is a helper function to see if a given password matches the hash
func ComparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
