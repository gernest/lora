package hashpass

import (
	"crypto/rand"
	"log"
	"strings"

	"code.google.com/p/go.crypto/bcrypt"
)

const (
	SaltLength  = 128
	EncryptCost = 31
)

type Password struct {
	Hash string
	Salt string
}

func HashPassword(salted_password string) string {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(salted_password), EncryptCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashed_password)
}

func Combine(salt string, raw_password string) string {
	pieces := []string{salt, raw_password}
	salted_password := strings.Join(pieces, "")
	return salted_password
}

func GenerateSalt() string {
	data := make([]byte, SaltLength)
	_, err := rand.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	// Convert to a string
	salt := string(data[:])
	return salt
}

func CreatePassword(raw_password string) *Password {
	password := new(Password)
	password.Salt = GenerateSalt()
	salted_password := Combine(password.Salt, raw_password)
	password.Hash = HashPassword(salted_password)
	return password
}

func (password *Password) MatchPassword(guess string) bool {
	salted_guess := Combine(password.Salt, guess)
	if bcrypt.CompareHashAndPassword([]byte(password.Hash), []byte(salted_guess)) != nil {
		return false
	}
	return true
}
