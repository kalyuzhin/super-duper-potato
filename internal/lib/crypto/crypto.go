package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"math/big"

	"golang.org/x/crypto/argon2"

	"github.com/kalyuzhin/password-manager/internal/model"
	"github.com/kalyuzhin/password-manager/pkg/errorspkg"
)

const saltSize = 16

// GenerateRandomSecurePassword – generates secure password
func GenerateRandomSecurePassword(length uint8) (string, error) {
	if length < 8 || length > 64 {
		return "", errorspkg.New("length must be between 8 and 64")
	}

	passwordRunes := make([]rune, 0, length)

	for i := uint8(0); i < length; i++ {
		symbol, err := getRandomSymbol()
		if err != nil {
			return "", err
		}
		passwordRunes = append(passwordRunes, symbol)
	}

	return string(passwordRunes), nil
}

func getRandomSymbol() (rune, error) {
	idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(model.All))))
	if err != nil {
		return 0, err
	}

	return rune(model.All[idx.Int64()]), nil
}

// GetArgonKey – ...
func GetArgonKey(masterPassword string, salt []byte) []byte {
	key := argon2.IDKey([]byte(masterPassword), salt, 3, 32*1024, 4, 32)

	return key
}

// GenerateArgon2Key – generates new argon2 key
func GenerateArgon2Key(masterPassword string) (key []byte, salt []byte, err error) {
	salt, err = generateSalt(saltSize)
	if err != nil {
		return nil, nil, err
	}
	key = argon2.IDKey([]byte(masterPassword), salt, 3, 32*1024, 4, 32)

	return key, salt, nil
}

// Encrypt – ...
func Encrypt(password string, key []byte) (cipherText []byte, nonce []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	nonce, err = generateSalt(int64(aead.NonceSize()))
	if err != nil {
		return nil, nil, err
	}

	cipherText = aead.Seal(nil, nonce, []byte(password), nil)

	return cipherText, nonce, nil
}

// Decrypt – ...
func Decrypt(cipherText, key, nonce []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	decrypted, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

func generateSalt(saltLength int64) ([]byte, error) {
	if saltLength == 0 {
		return nil, errorspkg.New("salt length cannot be 0")
	}

	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}
