package crypto

import "encoding/base64"

// ByteToBase64 – ...
func ByteToBase64(bytes []byte) string {
	return base64.RawStdEncoding.EncodeToString(bytes)
}

// ClearMemory – ...
func ClearMemory(data []byte) {
	for i := range data {
		data[i] = 0
	}
}
