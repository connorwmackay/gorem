package cmd

/*
 * Password Hashing Logic
 */

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

func GenSalt() []byte {
	salt := make([]byte, 16)
	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

func HashPassword(password string, salt []byte) string {
	hasher := sha512.New()

	passwordBytes := []byte(password)
	passwordBytes = append(passwordBytes, salt...)

	hasher.Write(passwordBytes)
	passwordHash := hasher.Sum(nil)
	passwordHashBase64 := base64.URLEncoding.EncodeToString(passwordHash)

	return passwordHashBase64
}

func CheckHashedPasswords(password string, passwordHash string, salt []byte) bool {
	passwordHashCheck := HashPassword(password, salt)

	if passwordHash == passwordHashCheck {
		return true
	} else {
		return false
	}
}
