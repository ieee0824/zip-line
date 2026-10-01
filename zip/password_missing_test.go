package zip

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestEncryptedOpenWithoutPassword(t *testing.T) {
	for _, method := range []EncryptionMethod{StandardEncryption, AES128Encryption, AES192Encryption, AES256Encryption} {
		path := filepath.Join(t.TempDir(), "input.txt")
		if err := os.WriteFile(path, []byte("hello"), 0600); err != nil {
			t.Fatal(err)
		}
		stat, err := os.Stat(path)
		if err != nil {
			t.Fatal(err)
		}
		var buf bytes.Buffer
		w := NewWriter(&buf)
		entry, err := w.Encrypt("input.txt", stat, "secret", method)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := entry.Write([]byte("hello")); err != nil {
			t.Fatal(err)
		}
		if err := w.Close(); err != nil {
			t.Fatal(err)
		}
		r, err := NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := r.File[0].Open(); !errors.Is(err, ErrPassword) {
			t.Fatalf("method %v: expected ErrPassword, got %v", method, err)
		}
	}
}
