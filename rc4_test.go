package RC4

import (
	"testing"
)

func TestRC4EncryptDecrypt(t *testing.T) {
	testCases := []struct {
		name string
		key  string
		data string
	}{
		{"Short message", "secret", "Hello, World!"},
		{"Empty message", "key", ""},
		{"Long message", "aVeryLongKey1234567890", "This is a longer message to test RC4 encryption with different key sizes"},
		{"Unicode", "密码", "你好，世界！"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encrypted := RC4([]byte(tc.key), []byte(tc.data))
			decrypted := RC4([]byte(tc.key), encrypted)

			if string(decrypted) != tc.data {
				t.Errorf("Decrypted data does not match original\nOriginal: %q\nDecrypted: %q", tc.data, decrypted)
			}
		})
	}
}
