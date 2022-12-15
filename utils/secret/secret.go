package secret

import (
	"crypto/sha512"
	"event-management/utils/env"

	p "github.com/wuriyanto48/go-pbkdf2"
)

func GenerateHash(secret string) string {
	pass := p.NewPassword(sha512.New, env.GetInt("HASH_SALT_SIZE", 1), env.GetInt("HASH_KEY_LENGTH", 12), env.GetInt("HASH_ITERATION", 100))

	hashed := pass.HashPassword(secret)
	hashedSecret := hashed.CipherText + hashed.Salt

	return hashedSecret
}

func VerifyHash(secret string, hashedSecret string) bool {
	pass := p.NewPassword(sha512.New, env.GetInt("HASH_SALT_SIZE", 1), env.GetInt("HASH_KEY_LENGTH", 12), env.GetInt("HASH_ITERATION", 100))

	actualHashedSecret := hashedSecret[:len(hashedSecret)-16]
	salt := hashedSecret[len(hashedSecret)-16:]

	return pass.VerifyPassword(secret, actualHashedSecret, salt)
}
