package crypto

import "golang.org/x/crypto/bcrypt"

func Encode(plaintext string) (string, error) {
	cipherBytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	return string(cipherBytes), err
}

func Matches(ciphertext, plaintext string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(ciphertext), []byte(plaintext))
	return err == nil
}
