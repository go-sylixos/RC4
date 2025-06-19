// Package RC4 provides RC4 stream cipher encryption and decryption functionality.
//
// WARNING: RC4 is considered cryptographically broken and should not be used
// for new systems. This package is provided for legacy compatibility only.
package RC4

import (
	"crypto/rc4"
	"errors"
)

var (
	// ErrKeySize is returned when the key size is invalid
	ErrKeySize = errors.New("RC4: invalid key size (must be 1-256 bytes)")
)

// Encrypt performs RC4 encryption/decryption on the given data.
//
// RC4 is a symmetric cipher, so the same function can be used for both
// encryption and decryption.
//
// Parameters:
//
//	key  - the encryption key (1-256 bytes)
//	data - the data to encrypt/decrypt
//
// Returns:
//
//	encrypted/decrypted data
//	error if key is invalid
func RC4(key, data []byte) ([]byte, error) {
	if len(key) < 1 || len(key) > 256 {
		return nil, ErrKeySize
	}

	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	dst := make([]byte, len(data))
	cipher.XORKeyStream(dst, data)
	return dst, nil
}
