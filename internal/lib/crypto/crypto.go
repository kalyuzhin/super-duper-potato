package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hkdf"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"math/big"
	"runtime"

	"golang.org/x/crypto/argon2"

	"github.com/kalyuzhin/password-manager/internal/model"
	"github.com/kalyuzhin/password-manager/pkg/errorspkg"
)

const saltSize = 16

const (
	encryption = "encryption"
	auth       = "auth"

	memory  = 1024 * 128
	kdfTime = 4
)

var (
	threads = uint8(runtime.NumCPU())
)

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
	key := argon2.IDKey([]byte(masterPassword), salt, kdfTime, memory, threads, 64)

	return key
}

// GenerateArgon2Key – generates new argon2 key
func GenerateArgon2Key(masterPassword string) (key []byte, salt []byte, err error) {
	salt, err = generateSalt(saltSize)
	if err != nil {
		return nil, nil, err
	}
	key = argon2.IDKey([]byte(masterPassword), salt, kdfTime, memory, threads, 64)

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

// DeriveKeys – ...
func DeriveKeys(masterKey []byte) (encKey, authKey []byte, err error) {
	hash := sha256.New

	encKey, err = hkdf.Expand(hash, masterKey, encryption, 32)
	if err != nil {
		return nil, nil, err
	}
	authKey, err = hkdf.Expand(hash, masterKey, auth, 32)
	if err != nil {
		return nil, nil, err
	}

	return encKey, authKey, nil
}

// GetHash – ...
func GetHash(key []byte) []byte {
	hash := sha256.Sum256(key)

	return hash[:]

}

// CompareHash – ...
func CompareHash(x, y []byte) bool {
	return subtle.ConstantTimeCompare(x, y) == 1
}
