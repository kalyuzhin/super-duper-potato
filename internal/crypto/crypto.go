package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"math/big"

	"golang.org/x/crypto/argon2"

	"github.com/kalyuzhin/password-manager/internal/model"
)

// GenerateRandomSecurePassword – generates secure password
func GenerateRandomSecurePassword(length uint8) string {
	if length < 8 || length > 32 {
		return ""
	}

	passwordRunes := make([]rune, 0, length)

	for i := uint8(0); i < length; i++ {
		passwordRunes = append(passwordRunes, getRandomSymbol())
	}

	return string(passwordRunes)
}

func getRandomSymbol() rune {
	idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(model.All))))
	if err != nil {
		return 0
	}

	return rune(model.All[idx.Int64()])
}

// GetArgonKey – ...
func GetArgonKey(masterPassword string, salt []byte) []byte {
	key := argon2.Key([]byte(masterPassword), salt, 3, 32*1024, 4, 32)

	return key
}

// GenerateArgon2Key – generates new argon2 key
func GenerateArgon2Key(masterPassword string, saltLength int64) (key []byte, salt []byte) {
	salt = generateSalt(saltLength)
	key = argon2.IDKey([]byte(masterPassword), salt, 3, 32*1024, 4, 32)

	return key, salt
}

// Encrypt – ...
func Encrypt(password string, key []byte) (cipherText []byte, nonce []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil
	}

	nonce = generateSalt(int64(aead.NonceSize()))

	cipherText = aead.Seal(nil, nonce, []byte(password), nil)

	return cipherText, nonce
}

// Decrypt – ...
func Decrypt(cipherText, key, nonce []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return ""
	}

	decrypted, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return ""
	}

	return string(decrypted)
}

func generateSalt(saltLength int64) []byte {
	if saltLength == 0 {
		return nil
	}

	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return nil
	}

	return salt
}
