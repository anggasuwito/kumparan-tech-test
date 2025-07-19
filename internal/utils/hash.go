package utils

import "golang.org/x/crypto/bcrypt"

func CompareHashCredential(credential string, hashCredential string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashCredential), []byte(credential))
	return err == nil
}

func HashCredential(credential string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(credential), 14)
	return string(bytes), err
}
