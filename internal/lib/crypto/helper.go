package crypto

import "encoding/base64"

// ByteToBase64 – ...
func ByteToBase64(bytes []byte) string {
	return base64.RawStdEncoding.EncodeToString(bytes)
}
