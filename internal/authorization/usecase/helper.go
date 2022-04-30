package usecase

import (
	"crypto/sha1"
	"fmt"
)

func (auth *AuthUseCase) hashPassword(password string) string {
	sha1Hasher := sha1.New()
	pwdBytes := []byte(password)
	// append hased password, with salt
	pwdBytes = append(pwdBytes, []byte(auth.HashSalt)...)
	sha1Hasher.Write(pwdBytes)                    // write bytes - to hasher
	return fmt.Sprintf("%x", sha1Hasher.Sum(nil)) // hashed password
	// base64EncodingPasswordHash := base64.URLEncoding.EncodeToString(hashPasswordBytes)
}
