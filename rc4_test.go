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
			encrypted, err := RC4([]byte(tc.key), []byte(tc.data))
			if err != nil {
				t.Fatalf("Encryption failed: %v", err)
			}

			decrypted, err := RC4([]byte(tc.key), encrypted)
			if err != nil {
				t.Fatalf("Decryption failed: %v", err)
			}

			if string(decrypted) != tc.data {
				t.Errorf("Decrypted data does not match original\nOriginal: %q\nDecrypted: %q", tc.data, decrypted)
			}
		})
	}
}

func TestInvalidKey(t *testing.T) {
	// 测试空密钥
	_, err := RC4([]byte(""), []byte("test"))
	if err == nil {
		t.Error("Expected error for empty key, got nil")
	}

	// 测试过长密钥(超过256字节)
	longKey := make([]byte, 257)
	_, err = RC4(longKey, []byte("test"))
	if err == nil {
		t.Error("Expected error for too long key, got nil")
	}
}
