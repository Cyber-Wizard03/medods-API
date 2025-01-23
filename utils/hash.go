package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

// CreatePassword will create password using bcrypt
func Hash(data string, hash string) string {
	pwdByte := []byte(data)
	salt := []byte(hash)

	// Create sha-512 hasher
	var sha512 = sha512.New()

	pwdByte = append(pwdByte, salt...)

	sha512.Write(pwdByte)

	// Get the SHA-512 hashed password
	var hashedPassword = sha512.Sum(nil)

	// Convert the hashed to hex string
	var hashedPasswordHex = hex.EncodeToString(hashedPassword)
	return hashedPasswordHex

}

// func CreatePassword(passwordString string) (string, error) {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordString), 8)
// 	if err != nil {
// 		return "", errors.New("Error occurred while creating a Hash")
// 	}

// 	return string(hashedPassword), nil
// }

// // ComparePasswords will create password using bcrypt
// func ComparePasswords(password string, hashedPassword string) error {
// 	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// 	if err != nil {
// 		return errors.New("The '" + password + "' and '" + hashedPassword + "' strings don't match")
// 	}
// 	return nil
// }
